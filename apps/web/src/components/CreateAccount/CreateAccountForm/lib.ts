import type { PartialMessage } from '@bufbuild/protobuf';
import { zodResolver } from '@hookform/resolvers/zod';
import { setCookie } from 'nookies';
import { Dispatch, SetStateAction } from 'react';
import { useForm } from 'react-hook-form';
import { z } from 'zod';
import { CreateUserRequest } from '@/__generated__/services/odachin_pb';
import { clientProvider } from '@/pages/api/ClientProvider';

export const createAccountFormSchema = z.object({
  userId: z.string().min(1, 'User ID must be at least 1 characters'),
  userName: z.string().min(1, 'User Name must be at least 1 characters'),
  email: z.string().email('Invalid email format'),
  password: z.string().min(8, 'Password must be at least 8 characters'),
});

export type CreateAccountFormSchemaType = z.infer<typeof createAccountFormSchema>;

export const useCreateAccountForm = (setToken: Dispatch<SetStateAction<string>>) => {
  const { register, handleSubmit, formState } = useForm<CreateAccountFormSchemaType>({
    resolver: zodResolver(createAccountFormSchema),
  });
  const onSubmit = async (data: CreateAccountFormSchemaType) => {
    const client = clientProvider();
    const req: PartialMessage<CreateUserRequest> = {
      userId: data.userId,
      name: data.userName,
      email: data.email,
      password: data.password,
    };
    const res = await client.createUser(req);
    setToken(res.status.toString());
    setCookie(null, 'auth', res.status.toString(), {
      maxAge: 60 * 60 * 24 * 7, // 1 week
      path: '/',
    });
    //TODO
    window.location.href = '/';
  };
  return { register, onSubmit: handleSubmit(onSubmit), formState };
};
