import { DescService } from '@bufbuild/protobuf';
import { Client, createClient, Interceptor } from '@connectrpc/connect';
import { createGrpcWebTransport } from '@connectrpc/connect-web';
import { useMemo } from 'react';

const transport = createGrpcWebTransport({
  baseUrl: process.env.NEXT_PUBLIC_BACKEND_URL as string,
});

/**
 * Get a promise client for the given service.
 */
export function useClient<T extends DescService>(service: T): Client<T> {
  // We memoize the client, so that we only create one instance per service.
  return useMemo(() => createClient(service, transport), [service]);
}
