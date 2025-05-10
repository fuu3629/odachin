import { zodResolver } from '@hookform/resolvers/zod';
import { useRouter } from 'next/router';
import { setCookie } from 'nookies';
import { useForm } from 'react-hook-form';
import { z } from 'zod';
import { AuthService } from '@/__generated__/v1/odachin/auth_pb';
import { FamilyService } from '@/__generated__/v1/odachin/faimily_pb';
import { useClient } from '@/pages/api/ClientProvider';

export const createGroupSchema = z.object({
  familyName: z.string().min(1, 'ファミリー名は必須です'),
});

export type createGroupSchemaType = z.infer<typeof createGroupSchema>;

export const useCreateGroupForm = () => {
  const { register, handleSubmit, formState } = useForm<createGroupSchemaType>({
    resolver: zodResolver(createGroupSchema),
  });
  const client = useClient(FamilyService);
  const router = useRouter();
  const onSubmit = async (data: createGroupSchemaType) => {
    const req = {
      familyName: data.familyName,
    };
    try {
      await client.createGroup(req);
      await new Promise((resolve) => setTimeout(resolve, 1000));
      router.push('/myPage/family');
    } catch (e) {
      alert('Create Group failed');
    }
  };
  return { register, onSubmit: handleSubmit(onSubmit), formState };
};
