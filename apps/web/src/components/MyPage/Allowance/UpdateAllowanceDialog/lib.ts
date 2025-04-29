import { zodResolver } from '@hookform/resolvers/zod';
import { Dispatch, SetStateAction, useContext } from 'react';
import { useForm } from 'react-hook-form';
import { z } from 'zod';
import { AllowanceItem } from '../AllowancePage';
import {
  Alloance_Type,
  AllowanceService,
  DayOfWeek,
} from '@/__generated__/v1/odachin/allowance_pb';
import { useClient } from '@/pages/api/ClientProvider';
import { CokiesContext } from '@/pages/api/CokiesContext';

export const updateAccountFormSchema = z
  .object({
    amount: z.string().refine(
      (amount) => {
        const num = Number(amount);
        return !isNaN(num) && num > 0;
      },
      {
        message: '金額は必須です',
      },
    ),
    allowanceType: z.enum(['DAILY', 'WEEKLY', 'MONTHLY'], {
      errorMap: () => ({ message: 'タイプは必須です' }),
    }),
    dayOfWeek: z
      .enum(['MONDAY', 'TUESDAY', 'WEDNESDAY', 'THURSDAY', 'FRIDAY', 'SATURDAY', 'SUNDAY'], {
        errorMap: () => ({ message: '曜日は必須です' }),
      })
      .optional(),
    date: z
      .number()
      .min(1, '日付は1以上でなければなりません')
      .max(31, '日付は31以下でなければなりません')
      .optional(),
  })
  .superRefine((data, ctx) => {
    if (data.allowanceType === 'WEEKLY' && !data.dayOfWeek) {
      ctx.addIssue({
        path: ['dayOfWeek'],
        message: '曜日は必須です',
        code: z.ZodIssueCode.custom,
      });
    }

    if (data.allowanceType === 'MONTHLY' && data.date == null) {
      ctx.addIssue({
        path: ['date'],
        message: '日付は必須です',
        code: z.ZodIssueCode.custom,
      });
    }
  });

export type UpdateAccountFormSchemaType = z.infer<typeof updateAccountFormSchema>;

export const useUpdateAccountForm = (
  defaultValues: UpdateAccountFormSchemaType,
  allowanceItem: AllowanceItem,
  setMessage: Dispatch<SetStateAction<boolean | undefined>>,
  setRefreshKey: Dispatch<SetStateAction<number>>,
  setOpen: Dispatch<SetStateAction<boolean>>,
) => {
  const client = useClient(AllowanceService);
  const cokkies = useContext(CokiesContext);

  const { register, handleSubmit, formState, ...rest } = useForm<UpdateAccountFormSchemaType>({
    resolver: zodResolver(updateAccountFormSchema),
    defaultValues: defaultValues,
  });
  const onSubmit = async (data: UpdateAccountFormSchemaType) => {
    if (!cokkies || !cokkies.authorization) {
      console.error('No authentication token found');
      return;
    }
    let date = undefined;
    let dayOfWeek = undefined;
    if (data.allowanceType === 'WEEKLY') {
      dayOfWeek = DayOfWeek[data.dayOfWeek! as keyof typeof DayOfWeek];
    }
    if (data.allowanceType === 'MONTHLY') {
      date = data.date;
    }
    const req = {
      allowanceId: allowanceItem.id,
      amount: Number(data.amount),
      intervalType: Alloance_Type[data.allowanceType as keyof typeof Alloance_Type],
      date: date,
      dayOfWeek: dayOfWeek,
    };
    try {
      const res = await client.updateAllowance(req, {
        headers: { authorization: cokkies.authorization },
      });
      setMessage(true);
      setRefreshKey((prev) => prev + 1);
      setOpen(false);
    } catch (error) {
      console.error('Error updating allowance:', error);
      setMessage(false);
    }
  };
  return { register, onSubmit: handleSubmit(onSubmit), formState, ...rest };
};
