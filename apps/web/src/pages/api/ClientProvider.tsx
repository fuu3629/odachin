import { createPromiseClient, Interceptor } from '@bufbuild/connect';
import { createGrpcWebTransport } from '@bufbuild/connect-web';
import { useContext } from 'react';
import { CokiesContext } from './CokiesContext';
import { OdachinService } from '@/__generated__/services/v1/odachin/odachin_connectweb';

export function clientProvider(token?: string) {
  const authInterceptor: Interceptor = (next) => async (req) => {
    if (token != null) {
      // リクエストヘッダーにトークンをセットする
      req.header.set('authorization', `${token}`);
    }
    return await next(req);
  };

  const transport = createGrpcWebTransport({
    baseUrl: process.env.NEXT_PUBLIC_BACKEND_URL as string,
    interceptors: [authInterceptor],
  });

  return createPromiseClient(OdachinService, transport);
}
