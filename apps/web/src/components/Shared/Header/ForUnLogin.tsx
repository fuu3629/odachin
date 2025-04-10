import { Button, HStack } from '@chakra-ui/react';
import { useRouter } from 'next/router';

export interface ForUnLoginProps {}

export function ForUnLogin({}: ForUnLoginProps) {
  const router = useRouter();
  const handleClickLogin = () => {
    router.push('/login');
  };
  const handleCreateNewAccount = () => {
    router.push('/createNewAccount');
  };
  return (
    <>
      <HStack>
        <Button onClick={handleClickLogin}>ログイン</Button>
        <Button onClick={handleCreateNewAccount}>新規登録</Button>
      </HStack>
    </>
  );
}
