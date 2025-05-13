import { zodResolver } from '@hookform/resolvers/zod';
import { useRouter } from 'next/router';
import { useForm } from 'react-hook-form';
import { z } from 'zod';
import { UsageService } from '@/__generated__/v1/odachin/usage_pb';
import { toaster } from '@/components/ui/toaster';
import { useClient } from '@/pages/api/ClientProvider';

export const applicateUsageSchema = z.object({
  title: z.string().min(1, 'タイトルは必須です'),
  amount: z.number().positive('正の値を入力してください'),
  description: z.string().min(1, '説明は必須です'),
  category: z.object({ label: z.string(), value: z.string() }).refine(
    (data) => {
      return data.value !== '';
    },
    {
      message: 'カテゴリは必須です',
    },
  ),
});

export type applicateUsageSchemaType = z.infer<typeof applicateUsageSchema>;

export const useApplicateUsageForm = () => {
  const { register, handleSubmit, formState, control } = useForm<applicateUsageSchemaType>({
    resolver: zodResolver(applicateUsageSchema),
  });
  const client = useClient(UsageService);
  const router = useRouter();
  const onSubmit = async (data: applicateUsageSchemaType) => {
    const req = {
      type: 'USAGE',
      title: data.title,
      amount: data.amount,
      description: data.description,
      category: data.category.value,
    };
    console.log(req);
    try {
      await client.applicateUsage(req);
      toaster.create({
        title: 'リクエストが完了しました',
        type: 'success',
      });
    } catch (e) {
      toaster.create({
        title: 'エラーが発生しました',
        type: 'error',
      });
    }
  };
  return { register, onSubmit: handleSubmit(onSubmit), formState, control };
};
