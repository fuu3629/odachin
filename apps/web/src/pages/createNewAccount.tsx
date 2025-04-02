import { Flex } from '@chakra-ui/react';
import { useState } from 'react';
import { CreateAccountForm } from '@/components/CreateAccount/CreateAccountForm';

export default function CreateNewAccount() {
  const [token, setToken] = useState<string>('');

  return (
    <Flex alignItems='center' height='100vh' justifyContent='center' w='100%'>
      <CreateAccountForm setToken={setToken} />
    </Flex>
  );
}
