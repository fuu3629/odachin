import {
  Flex,
  Heading,
  Input,
  Button,
  HStack,
  Text,
  Select,
  Spacer,
  Link,
  Center,
} from '@chakra-ui/react';
import { useRouter } from 'next/router';
import { useLoginForm } from './lib';
export interface LoginFormProps {}

export function LoginForm({}: LoginFormProps) {
  const { register, onSubmit, formState } = useLoginForm();
  const router = useRouter();
  const handleCreateNewAccount = async () => {
    await router.push('/createNewAccount');
  };
  return (
    <>
      <Flex background='gray.200' direction='column' pb={6} pt={12} px={12} rounded={6}>
        <Heading mb={6}>Log in</Heading>
        <form onSubmit={onSubmit}>
          <HStack mb={6}>
            <Text h='100%' w='150px'>
              UserName
            </Text>
            <Input bg='white' placeholder='sample@sample.com' {...register('name')} />
          </HStack>
          <HStack mb={12}>
            <Text h='100%' w='150px'>
              Password
            </Text>
            <Input bg='white' placeholder='********' {...register('password')} />
          </HStack>
          <HStack>
            <Spacer></Spacer>
            <Button colorScheme='teal' mb={16} type='submit' w={32}>
              Login
            </Button>
          </HStack>
          <Center>
            <Link
              _hover={{ cursor: 'pointer' }}
              color='blue.400'
              onClick={() => {
                handleCreateNewAccount();
              }}
            >
              crete New Account
            </Link>
          </Center>
        </form>
      </Flex>
    </>
  );
}
