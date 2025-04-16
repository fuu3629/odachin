import { createClient, Interceptor } from '@connectrpc/connect';
import { createGrpcWebTransport, createConnectTransport } from '@connectrpc/connect-web';
import { OdachinService } from '@/__generated__/v1/odachin/odachin_pb';

export function clientProvider(token?: string) {
  // const authInterceptor: Interceptor = (next) => async (req) => {
  //   if (token != null) {
  //     // リクエストヘッダーにトークンをセットする
  //     req.header.set('authorization', `${token}`);
  //   }
  //   return await next(req);
  // };

  const transport = createGrpcWebTransport({
    baseUrl: process.env.NEXT_PUBLIC_BACKEND_URL as string,
    // interceptors: [authInterceptor],
  });

  return createClient(OdachinService, transport);
}
