import { Flex } from '@chakra-ui/react';
import { LoginForm } from '@/components/Login/LoginForm';

export default function Login() {
  return (
    <Flex alignItems='center' height='100vh' justifyContent='center' w='100%'>
      <LoginForm></LoginForm>
    </Flex>
  );
}
