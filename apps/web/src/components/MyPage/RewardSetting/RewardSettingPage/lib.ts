import { zodResolver } from '@hookform/resolvers/zod';
import { Dispatch, SetStateAction, useContext } from 'react';
import { useForm } from 'react-hook-form';
import { z } from 'zod';
import { Reward_Type, RewardService } from '@/__generated__/v1/odachin/reward_pb';
import { useClient } from '@/pages/api/ClientProvider';
import { CokiesContext } from '@/pages/api/CokiesContext';

export const addRewardFormSchema = z.object({
  title: z.string().min(1, 'タイトルは必須です'),
  description: z.string().min(1, '説明は必須です'),
  amount: z.coerce
    .number({ required_error: '金額は必須です' })
    .min(1, '金額は1以上でなければなりません'),
});

export type AddRewardFormSchemaType = z.infer<typeof addRewardFormSchema>;

export const useAddRewardForm = (
  setRefreshKey: Dispatch<SetStateAction<number>>,
  setOpen: Dispatch<SetStateAction<boolean>>,
  to_user_id?: string,
  rewardType?: Reward_Type,
) => {
  const client = useClient(RewardService);
  const cookies = useContext(CokiesContext);

  const { register, handleSubmit, formState, ...rest } = useForm<AddRewardFormSchemaType>({
    resolver: zodResolver(addRewardFormSchema),
  });

  const onSubmit = async (data: AddRewardFormSchemaType) => {
    console.log('Submitting form:', data);
    if (!cookies || !cookies.authorization) {
      console.error('No authentication token found');
      return;
    }
    try {
      const req = {
        toUserId: to_user_id,
        title: data.title,
        description: data.description,
        amount: data.amount,
        rewardType: rewardType,
      };
      console.log('Request:', req);
      await client.registerReward(req, {
        headers: { authorization: cookies?.authorization },
      });
      setRefreshKey((prev) => prev + 1);
      setOpen(false);
    } catch (error) {
      console.error('Error adding reward:', error);
    }
  };

  return { register, onSubmit: handleSubmit(onSubmit), formState, ...rest };
};
