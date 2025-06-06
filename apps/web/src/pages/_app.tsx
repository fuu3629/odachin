// import '@/styles/globals.css';
import { Box, Flex, Toast } from '@chakra-ui/react';
import { NextPageContext } from 'next';
import type { AppProps } from 'next/app';
import { useRouter } from 'next/router';
import { parseCookies } from 'nookies';
import { useEffect } from 'react';
import { CokiesContext } from './api/CokiesContext';
import { Header } from '@/components/Shared/Header';
import { Provider } from '@/components/ui/provider';
import { Toaster } from '@/components/ui/toaster';

export default function App({ Component, pageProps }: AppProps, ctx: NextPageContext) {
  const router = useRouter();
  const cookies = parseCookies(ctx);

  // 第二引数に空配列を指定してマウント・アンマウント毎（CSRでの各画面遷移時）に呼ばれるようにする
  useEffect(() => {
    // CSR用認証チェック

    // const notAuthPath = ["/login", "/createNewAccount", "/_error", "/", ""];

    router.beforePopState(({ url, as, options }) => {
      // ログイン画面とエラー画面遷移時のみ認証チェックを行わない
      if (url !== '/login' && url !== '/createNewAccount' && url !== '/_error' && url !== '/') {
        if (typeof cookies.authorization === 'undefined') {
          // CSR用リダイレクト処理
          window.location.href = '/login';
          return false;
        }
      }
      return true;
    });
  }, []);
  return (
    <Provider>
      <CokiesContext.Provider value={cookies}>
        <Toaster />
        <Box>
          <Header></Header>
          <Flex>
            <Component {...pageProps} />
          </Flex>
        </Box>
      </CokiesContext.Provider>
    </Provider>
  );
}

App.getInitialProps = async (appContext: any) => {
  // SSR用認証チェック

  const cookies = parseCookies(appContext.ctx);
  // ログイン画面とエラー画面遷移時のみ認証チェックを行わない
  if (
    appContext.ctx.pathname !== '/login' &&
    appContext.ctx.pathname !== '/createNewAccount' &&
    appContext.ctx.pathname !== '/_error' &&
    appContext.ctx.pathname !== '/'
  ) {
    if (typeof cookies.authorization === 'undefined') {
      // SSR or CSRを判定
      const isServer = typeof window === 'undefined';
      if (isServer) {
        console.log('in ServerSide');
        appContext.ctx.res.statusCode = 302;
        appContext.ctx.res.setHeader('Location', '/login');
        return {};
      } else {
        console.log('in ClientSide');
      }
    }
  }
  return {
    pageProps: {
      ...(appContext.Component.getInitialProps
        ? await appContext.Component.getInitialProps(appContext.ctx)
        : {}),
      pathname: appContext.ctx.pathname,
    },
  };
};
