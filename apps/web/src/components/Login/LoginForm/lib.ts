import { zodResolver } from '@hookform/resolvers/zod';
import { setCookie } from 'nookies';
import { useForm } from 'react-hook-form';
import { z } from 'zod';
import { AuthService } from '@/__generated__/v1/odachin/auth_pb';
import { useClient } from '@/pages/api/ClientProvider';

export const loginFormSchema = z.object({
  userId: z.string().min(1, 'ユーザーIDは必須です'),
  password: z.string().min(1, 'パスワードは必須です'),
});

export type LoginFormSchemaType = z.infer<typeof loginFormSchema>;

export const useLoginForm = () => {
  const { register, handleSubmit, formState } = useForm<LoginFormSchemaType>({
    resolver: zodResolver(loginFormSchema),
  });
  const client = useClient(AuthService);
  const onSubmit = async (data: LoginFormSchemaType) => {
    const req = {
      userId: data.userId,
      password: data.password,
    };
    try {
      const res = await client.login(req);
      setCookie(null, 'authorization', res.token, {
        maxAge: 30 * 24 * 60 * 60,
        path: '/',
      });
      window.location.href = '/myPage';
    } catch (e) {
      alert('Login failed');
    }
  };
  return { register, onSubmit: handleSubmit(onSubmit), formState };
};
