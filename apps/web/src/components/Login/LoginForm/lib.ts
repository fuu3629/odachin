import type { PartialMessage } from '@bufbuild/protobuf';
import { zodResolver } from '@hookform/resolvers/zod';
import { setCookie } from 'nookies';
import { useForm } from 'react-hook-form';
import { z } from 'zod';
import { LoginRequest } from '@/__generated__/services/odachin_pb';
import { clientProvider } from '@/pages/api/ClientProvider';

export const loginFormSchema = z.object({
  userId: z.string().min(1, 'ユーザーIDは必須です'),
  password: z.string().min(1, 'パスワードは必須です'),
});

export type LoginFormSchemaType = z.infer<typeof loginFormSchema>;

export const useLoginForm = () => {
  const { register, handleSubmit, formState } = useForm<LoginFormSchemaType>({
    resolver: zodResolver(loginFormSchema),
  });
  const onSubmit = async (data: LoginFormSchemaType) => {
    const client = clientProvider();
    const req: PartialMessage<LoginRequest> = {
      userId: data.userId,
      password: data.password,
    };
    try {
      const res = await client.login(req);
      setCookie(null, 'auth', res.token, {
        maxAge: 30 * 24 * 60 * 60,
        path: '/',
      });
      console.log(res);
      window.location.href = '/myPage';
    } catch (e) {
      console.error(e);
      alert('Login failed');
    }
  };
  return { register, onSubmit: handleSubmit(onSubmit), formState };
};
