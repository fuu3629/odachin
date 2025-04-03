import {
  Flex,
  Heading,
  Input,
  Button,
  HStack,
  Text,
  Spacer,
  Box,
  Center,
  VStack,
} from '@chakra-ui/react';
import { Dispatch, SetStateAction } from 'react';
import { useCreateAccountForm } from './lib';
export interface CreateAccountFormProps {
  setToken: Dispatch<SetStateAction<string>>;
}

export function CreateAccountForm({ setToken }: CreateAccountFormProps) {
  const {
    register,
    onSubmit,
    formState: { errors },
  } = useCreateAccountForm(setToken);
  return (
    <Flex
      h='100vh'
      w='100vw'
      direction='column'
      alignItems='center'
      justifyContent='center'
      bgSize='cover'
    >
      <Box
        w='100%'
        maxW='500px'
        bg='white'
        p={8}
        boxShadow='lg'
        borderRadius='lg'
        bgImage="url('/mnt/data/image.png')"
      >
        <Flex bg='white' direction='column' justify='center' alignItems='center' p={12}>
          <Box w='100%' maxW='400px'>
            <form onSubmit={onSubmit}>
              <Heading mb={6}>新規登録</Heading>
              <Text fontSize='sm' mb={1} fontWeight='bold'>
                メールアドレスを入力
              </Text>
              <Input bg='gray.100' {...register('userId')} />
              {errors.email && (
                <Text color='red.500' fontSize='sm'>
                  {errors.email.message}
                </Text>
              )}
              <Text fontSize='sm' mb={1} mt={4} fontWeight='bold'>
                ユーザ名（表示名）
              </Text>
              <Input bg='gray.100' {...register('email')} w='full' />
              {errors.userName && (
                <Text color='red.500' fontSize='sm'>
                  {errors.userName.message}
                </Text>
              )}
              <Text fontSize='sm' mb={1} mt={4} fontWeight='bold'>
                パスワードを入力
              </Text>
              <Input bg='gray.100' {...register('password')} w='full' />
              {errors.password && (
                <Text color='red.500' fontSize='sm'>
                  {errors.password.message}
                </Text>
              )}

              <VStack mt={8} mb={2}>
                <Button type='submit' bg='black' color='white' w='full' _hover={{ bg: 'gray.800' }}>
                  新規登録
                </Button>
              </VStack>
            </form>
          </Box>
        </Flex>
      </Box>
    </Flex>
  );
}
