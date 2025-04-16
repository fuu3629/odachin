import { Flex } from '@chakra-ui/react';
import { CreateAccountForm } from '@/components/CreateAccount/CreateAccountForm';

export default function CreateNewAccount() {
  return (
    <Flex alignItems='center' height='100vh' justifyContent='center' w='100%'>
      <CreateAccountForm />
    </Flex>
  );
}
