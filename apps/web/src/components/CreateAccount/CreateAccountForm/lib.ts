import { create } from '@bufbuild/protobuf';
import { zodResolver } from '@hookform/resolvers/zod';
import { setCookie } from 'nookies';
import { useForm } from 'react-hook-form';
import { z } from 'zod';
import { CreateUserRequestSchema } from '@/__generated__/v1/odachin/odachin_pb';
import { clientProvider } from '@/pages/api/ClientProvider';

export const createAccountFormSchema = z.object({
  userId: z.string().min(1, 'User ID must be at least 1 characters'),
  userName: z.string().min(1, 'User Name must be at least 1 characters'),
  email: z.string().email('Invalid email format'),
  password: z.string().min(8, 'Password must be at least 8 characters'),
});

export type CreateAccountFormSchemaType = z.infer<typeof createAccountFormSchema>;

export const useCreateAccountForm = () => {
  const { register, handleSubmit, formState, ...rest } = useForm<CreateAccountFormSchemaType>({
    resolver: zodResolver(createAccountFormSchema),
  });
  const onSubmit = async (data: CreateAccountFormSchemaType) => {
    const client = clientProvider();
    const req = {
      userId: data.userId,
      name: data.userName,
      email: data.email,
      password: data.password,
    };
    const res = await client.createUser(req);
    setCookie(null, 'authorization', res.token, {
      maxAge: 60 * 60 * 24 * 7, // 1 week
      path: '/',
    });
    window.location.href = '/myPage';
  };
  return { register, onSubmit: handleSubmit(onSubmit), formState, ...rest };
};
