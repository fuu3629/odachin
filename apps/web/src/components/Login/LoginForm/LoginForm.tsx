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
        alignItems='center'
        bgSize='cover'
        direction='column'
        h='100vh'
        justifyContent='center'
        w='100vw'
      >
        <Box
          bg='white'
          bgImage="url('/mnt/data/image.png')"
          borderRadius='lg'
          boxShadow='lg'
          maxW='500px'
          p={8}
          w='100%'
        >
          <Flex alignItems='center' bg='white' direction='column' justify='center' p={12}>
            <Box maxW='400px' w='100%'>
              <form onSubmit={onSubmit} style={{ width: '100%' }}>
                <Text fontSize='sm' fontWeight='bold' mb={1}>
                  メールアドレスを入力
                </Text>
                <Input bg='gray.100' {...register('email')} />
                {errors.email && (
                  <Text color='red.500' fontSize='sm'>
                    {errors.email.message}
                  </Text>
                )}
                <Text fontSize='sm' fontWeight='bold' mb={1} mt={4}>
                  パスワードを入力
                </Text>
                <Input bg='gray.100' {...register('password')} w='full' />
                {errors.password && (
                  <Text color='red.500' fontSize='sm'>
                    {errors.password.message}
                  </Text>
                )}
                <VStack mb={2} mt={8}>
                  <Button
                    _hover={{ bg: 'gray.800' }}
                    bg='black'
                    color='white'
                    type='submit'
                    w='full'
                  >
                    ログイン
                  </Button>
                </VStack>
              </form>

              <Center color='gray.500' fontSize='sm' mb={4} mt={6}>
                または
              </Center>

              <Button
                borderColor='gray.300'
                mb={2}
                onClick={() => {
                  handleCreateNewAccount();
                }}
                variant='outline'
                w='full'
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
