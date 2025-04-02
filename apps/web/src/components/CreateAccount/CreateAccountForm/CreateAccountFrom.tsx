import { Flex, Heading, Input, Button, HStack, Text, Spacer } from '@chakra-ui/react';
import { Dispatch, SetStateAction } from 'react';
import { useCreateAccountForm } from './lib';
export interface CreateAccountFormProps {
  setToken: Dispatch<SetStateAction<string>>;
}

export function CreateAccountForm({ setToken }: CreateAccountFormProps) {
  const { register, onSubmit, formState } = useCreateAccountForm(setToken);
  return (
    <Flex background='gray.200' direction='column' padding={12} rounded={6}>
      <form onSubmit={onSubmit}>
        <Heading mb={6}>新規登録</Heading>
        <HStack mb={6}>
          <Text h='100%' w='250px'>
            User ID
          </Text>
          <Input bg='white' placeholder='sample' {...register('userId')} />
        </HStack>
        <HStack mb={6}>
          <Text h='100%' w='250px'>
            User Name(表示名)
          </Text>
          <Input bg='white' placeholder='sample' {...register('userName')} />
        </HStack>
        <HStack mb={6}>
          <Text h='100%' w='250px'>
            Password
          </Text>
          <Input bg='white' placeholder='********' {...register('password')} />
        </HStack>
        <HStack>
          <Spacer></Spacer>
          <Button colorScheme='teal' mb={6} type='submit'>
            crete New Account
          </Button>
        </HStack>
      </form>
    </Flex>
  );
}
