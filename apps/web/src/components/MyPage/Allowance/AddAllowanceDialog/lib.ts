import { zodResolver } from '@hookform/resolvers/zod';
import { Dispatch, SetStateAction, useContext } from 'react';
import { set, useForm } from 'react-hook-form';
import { z } from 'zod';
import { SelectedUser } from '../../RewardSetting/RewardSettingPage/RewardSettingPage';
import {
  Alloance_Type,
  AllowanceService,
  DayOfWeek,
} from '@/__generated__/v1/odachin/allowance_pb';
import { useClient } from '@/pages/api/ClientProvider';
import { CokiesContext } from '@/pages/api/CokiesContext';

export const addAllowanceFormSchema = z
  .object({
    intervalType: z.enum(['DAILY', 'WEEKLY', 'MONTHLY'], {
      errorMap: () => ({ message: 'タイプは必須です' }),
    }),
    amount: z.coerce
      .number({ required_error: '金額は必須です' })
      .min(1, '金額は1以上でなければなりません'),
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
    if (data.intervalType === 'WEEKLY' && !data.dayOfWeek) {
      ctx.addIssue({
        path: ['dayOfWeek'],
        message: '曜日は必須です',
        code: z.ZodIssueCode.custom,
      });
    }

    if (data.intervalType === 'MONTHLY' && data.date == null) {
      ctx.addIssue({
        path: ['date'],
        message: '日付は必須です',
        code: z.ZodIssueCode.custom,
      });
    }
  });

export type AddAllowanceFormSchemaType = z.infer<typeof addAllowanceFormSchema>;

export const useAddAllowanceForm = (
  setOpen: Dispatch<SetStateAction<boolean>>,
  setRefreshKey: Dispatch<SetStateAction<number>>,
  selectedUser?: SelectedUser,
) => {
  const client = useClient(AllowanceService);

  const { register, handleSubmit, formState, ...rest } = useForm<AddAllowanceFormSchemaType>({
    resolver: zodResolver(addAllowanceFormSchema),
  });

  const onSubmit = async (data: AddAllowanceFormSchemaType) => {
    try {
      let date = undefined;
      let dayOfWeek = undefined;
      if (data.intervalType === 'WEEKLY') {
        dayOfWeek = DayOfWeek[data.dayOfWeek! as keyof typeof DayOfWeek];
      }
      if (data.intervalType === 'MONTHLY') {
        date = data.date;
      }
      const req = {
        toUserId: selectedUser?.id,
        amount: data.amount,
        intervalType: Alloance_Type[data.intervalType as keyof typeof Alloance_Type],
        dayOfWeek: dayOfWeek,
        date: date,
      };
      const res = await client.registerAllowance(req);
      setRefreshKey((prev) => prev + 1);
      setOpen(false);
    } catch (error) {
      console.error('Error adding allowance:', error);
    }
  };
  return { register, onSubmit: handleSubmit(onSubmit), formState, ...rest };
};
