import { title } from 'process';
import { zodResolver } from '@hookform/resolvers/zod';
import { useRouter } from 'next/router';
import { setCookie } from 'nookies';
import { useForm } from 'react-hook-form';
import { z } from 'zod';
import { AuthService } from '@/__generated__/v1/odachin/auth_pb';
import { FamilyService } from '@/__generated__/v1/odachin/faimily_pb';
import { UsageService } from '@/__generated__/v1/odachin/usage_pb';
import { useClient } from '@/pages/api/ClientProvider';

export const applicateUsageSchema = z.object({
  title: z.string().min(1, 'タイトルは必須です'),
  amount: z.number().positive('正の値を入力してください'),
  description: z.string().min(1, '説明は必須です'),
  category: z.string().min(1, 'カテゴリは必須です'),
});

export type applicateUsageSchemaType = z.infer<typeof applicateUsageSchema>;

export const useApplicateUsageForm = () => {
  const { register, handleSubmit, formState } = useForm<applicateUsageSchemaType>({
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
      category: data.category,
    };
    console.log(req);
    try {
      await client.applicateUsage(req);
      await new Promise((resolve) => setTimeout(resolve, 1000));
    } catch (e) {
      alert('ApplicateUsage failed');
    }
  };
  return { register, onSubmit: handleSubmit(onSubmit), formState };
};
