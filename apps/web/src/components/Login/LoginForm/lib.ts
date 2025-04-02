import type { PartialMessage } from '@bufbuild/protobuf';
import { zodResolver } from '@hookform/resolvers/zod';
import { setCookie } from 'nookies';
import { useForm } from 'react-hook-form';
import { z } from 'zod';
import { OdachinService } from '@/__generated__/services/odachin_connectweb';
import { LoginRequest } from '@/__generated__/services/odachin_pb';
import { clientProvider } from '@/pages/api/ClientProvider';

export const loginFormSchema = z.object({
  name: z.string(),
  password: z.string(),
});

export type LoginFormSchemaType = z.infer<typeof loginFormSchema>;

export const useLoginForm = () => {
  const { register, handleSubmit, formState } = useForm<LoginFormSchemaType>({
    resolver: zodResolver(loginFormSchema),
  });
  const onSubmit = async (data: LoginFormSchemaType) => {
    const client = clientProvider();
    const req: PartialMessage<LoginRequest> = {
      userId: data.name,
      password: data.password,
    };
    try {
      console.log('login');
      const res = await client.login(req);
      setCookie(null, 'auth', res.token, {
        maxAge: 30 * 24 * 60 * 60,
        path: '/',
      });
      window.location.href = '/';
    } catch (e) {
      alert('Login failed');
    }
  };
  return { register, onSubmit: handleSubmit(onSubmit), formState };
};
