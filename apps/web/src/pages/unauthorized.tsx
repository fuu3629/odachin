import { Link } from '@chakra-ui/react';

export default function UnauthorizedPage() {
  return (
    <main>
      <p>Authorization Error</p>
      <Link href='/login'>ログインページに進む</Link>
    </main>
  );
}
