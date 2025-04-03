import { createPromiseClient } from '@bufbuild/connect';
import { createGrpcWebTransport } from '@bufbuild/connect-web';
import { OdachinService } from '@/__generated__/services/odachin_connectweb';

const transport = createGrpcWebTransport({
  baseUrl: process.env.NEXT_PUBLIC_BACKEND_URL as string,
});

export function clientProvider() {
  return createPromiseClient(OdachinService, transport);
}
