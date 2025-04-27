import { DescService } from '@bufbuild/protobuf';
import { Client, createClient, Interceptor } from '@connectrpc/connect';
import { createGrpcWebTransport } from '@connectrpc/connect-web';
import { useContext, useMemo } from 'react';
import { CokiesContext } from './CokiesContext';

export function useClient<T extends DescService>(service: T): Client<T> {
  const cookies = useContext(CokiesContext);

  const transport = useMemo(() => {
    const authInterceptor: Interceptor = (next) => async (req) => {
      if (cookies?.authorization) {
        req.header.set('authorization', cookies.authorization);
      }
      return await next(req);
    };

    return createGrpcWebTransport({
      baseUrl: process.env.NEXT_PUBLIC_BACKEND_URL!,
      interceptors: [authInterceptor],
    });
  }, [cookies?.authorization]);

  return useMemo(() => createClient(service, transport), [service, transport]);
}
