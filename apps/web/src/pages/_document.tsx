import { Html, Head, Main, NextScript } from 'next/document';

export default function Document() {
  return (
    <Html>
      <Head>
        <link href='/favicons/favicon.png' rel='icon' />
      </Head>
      <body style={{ backgroundColor: '' }}>
        <Main />
        <NextScript />
      </body>
    </Html>
  );
}
