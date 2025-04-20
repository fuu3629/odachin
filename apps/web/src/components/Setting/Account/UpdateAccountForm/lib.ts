import { zodResolver } from '@hookform/resolvers/zod';
import { useRouter } from 'next/router';
import { useContext } from 'react';
import { useForm } from 'react-hook-form';
import { z } from 'zod';
import { AuthService } from '@/__generated__/v1/odachin/auth_pb';
import { useClient } from '@/pages/api/ClientProvider';
import { CokiesContext } from '@/pages/api/CokiesContext';

export const updateAccountFormSchema = z.object({
  userName: z.string().min(1, 'User Name must be at least 1 characters'),
  email: z.string().email('Invalid email format'),
  avatar: z.custom<FileList>().transform((file) => (file ? file[0] : undefined)),
});

export type UpdateAccountFormSchemaType = z.infer<typeof updateAccountFormSchema>;

async function fileToUint8Array(file?: File): Promise<Uint8Array | undefined> {
  if (!file) return undefined;

  const arrayBuffer = await file.arrayBuffer(); // File を ArrayBuffer に変換
  return new Uint8Array(arrayBuffer); // ArrayBuffer を Uint8Array に変換
}

export const useUpdateAccountForm = () => {
  const cookies = useContext(CokiesContext);
  const router = useRouter();
  const { register, handleSubmit, formState, ...rest } = useForm<UpdateAccountFormSchemaType>({
    resolver: zodResolver(updateAccountFormSchema),
  });
  const client = useClient(AuthService);
  const onSubmit = async (data: UpdateAccountFormSchemaType) => {
    if (!cookies || !cookies.authorization) {
      console.error('No authentication token found');
      return;
    }
    const file = await fileToUint8Array(data.avatar);
    const req = {
      name: data.userName,
      email: data.email,
      profileImage: file,
    };
    try {
      const res = await client.updateUser(req, {
        headers: { authorization: cookies?.authorization },
      });
      router.push('/myPage');
    } catch (error) {
      alert('アップデートに失敗しました。');
    }
  };
  return { register, onSubmit: handleSubmit(onSubmit), formState, ...rest };
};
