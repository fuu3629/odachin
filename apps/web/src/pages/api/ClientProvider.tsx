import { createClient } from '@connectrpc/connect';
import { createGrpcWebTransport } from '@connectrpc/connect-web';
import { OdachinService } from '@/__generated__/v1/odachin/odachin_pb';

export function clientProvider() {
  const transport = createGrpcWebTransport({
    baseUrl: process.env.NEXT_PUBLIC_BACKEND_URL as string,
  });
  return createClient(OdachinService, transport);
}
