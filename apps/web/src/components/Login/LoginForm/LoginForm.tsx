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
  Box,
  VStack,
} from '@chakra-ui/react';
import { useRouter } from 'next/router';
import { useLoginForm } from './lib';
export interface LoginFormProps {}

export function LoginForm({}: LoginFormProps) {
  const {
    register,
    onSubmit,
    formState: { errors },
  } = useLoginForm();
  const router = useRouter();
  const handleCreateNewAccount = async () => {
    await router.push('/createNewAccount');
  };
  return (
    <>
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
              <form onSubmit={onSubmit} style={{ width: '100%' }}>
                <Text fontSize='sm' mb={1} fontWeight='bold'>
                  メールアドレスを入力
                </Text>
                <Input bg='gray.100' {...register('email')} />
                {errors.email && (
                  <Text color='red.500' fontSize='sm'>
                    {errors.email.message}
                  </Text>
                )}
                <Text fontSize='sm' mb={1} fontWeight='bold' mt={4}>
                  パスワードを入力
                </Text>
                <Input bg='gray.100' {...register('password')} w='full' />
                {errors.password && (
                  <Text color='red.500' fontSize='sm'>
                    {errors.password.message}
                  </Text>
                )}
                <VStack mt={8} mb={2}>
                  <Button
                    type='submit'
                    bg='black'
                    color='white'
                    w='full'
                    _hover={{ bg: 'gray.800' }}
                  >
                    ログイン
                  </Button>
                </VStack>
              </form>

              <Center mt={6} mb={4} fontSize='sm' color='gray.500'>
                または
              </Center>

              <Button
                variant='outline'
                w='full'
                mb={2}
                borderColor='gray.300'
                onClick={() => {
                  handleCreateNewAccount();
                }}
              >
                新規登録
              </Button>
            </Box>
          </Flex>
        </Box>
      </Flex>
    </>
  );
}
