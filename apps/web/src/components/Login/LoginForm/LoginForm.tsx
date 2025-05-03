import { Flex, Input, Button, Center, Box, VStack, Field } from '@chakra-ui/react';
import { useRouter } from 'next/router';
import { useLoginForm } from './lib';
import { PasswordInput } from '@/components/ui/password-input';
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
        <Box bg='white' borderRadius='lg' boxShadow='lg' maxW='500px' p={8} w='100%'>
          <Flex alignItems='center' bg='white' direction='column' justify='center' p={12}>
            <Box maxW='400px' w='100%'>
              <form onSubmit={onSubmit} style={{ width: '100%' }}>
                <VStack gap={4}>
                  <Field.Root invalid={!!errors.userId}>
                    <Field.Label fontSize='sm' fontWeight='bold' mb={1}>
                      ユーザーIDを入力
                    </Field.Label>
                    <Input bg='gray.100' placeholder='example' {...register('userId')} />
                    <Field.ErrorText fontSize='sm'>{errors.userId?.message}</Field.ErrorText>
                  </Field.Root>

                  <Field.Root invalid={!!errors.password}>
                    <Field.Label fontSize='sm' fontWeight='bold' mb={1}>
                      パスワードを入力
                    </Field.Label>
                    <PasswordInput
                      bg='gray.100'
                      placeholder='*********'
                      {...register('password')}
                      w='full'
                    />
                    <Field.ErrorText fontSize='sm'>{errors.password?.message}</Field.ErrorText>
                  </Field.Root>
                </VStack>

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
