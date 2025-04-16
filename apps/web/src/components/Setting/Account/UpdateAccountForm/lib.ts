import { zodResolver } from '@hookform/resolvers/zod';
import { useForm } from 'react-hook-form';
import { z } from 'zod';
import { GetOwnInfoResponse } from '@/__generated__/v1/odachin/odachin_pb';
import { clientProvider } from '@/pages/api/ClientProvider';

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

export const useUpdateAccountForm = (defaultValues?: GetOwnInfoResponse) => {
  const { register, handleSubmit, formState, ...rest } = useForm<UpdateAccountFormSchemaType>({
    resolver: zodResolver(updateAccountFormSchema),
    defaultValues: {
      userName: defaultValues?.name,
      email: defaultValues?.email,
      avatar: undefined,
    },
  });
  const onSubmit = async (data: UpdateAccountFormSchemaType) => {
    const client = clientProvider();
    const req = {
      name: data.userName,
      email: data.email,
    };
    const res = await client.createUser(req);
  };
  return { register, onSubmit: handleSubmit(onSubmit), formState, ...rest };
};
