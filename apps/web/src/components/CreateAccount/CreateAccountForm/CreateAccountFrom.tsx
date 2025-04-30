import { Flex, Heading, Input, Button, Text, Box, VStack } from '@chakra-ui/react';
import { useCreateAccountForm } from './lib';
export interface CreateAccountFormProps {}

//TODO ROLEの設定する
export function CreateAccountForm({}: CreateAccountFormProps) {
  const {
    register,
    onSubmit,
    formState: { errors },
  } = useCreateAccountForm();
  return (
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
            <form onSubmit={onSubmit}>
              <Heading mb={6}>新規登録</Heading>
              <Text fontSize='sm' fontWeight='bold' mb={1}>
                メールアドレスを入力
              </Text>
              <Input bg='gray.100' {...register('email')} />
              {errors.email && (
                <Text color='red.500' fontSize='sm'>
                  {errors.email.message}
                </Text>
              )}
              <Text fontSize='sm' fontWeight='bold' mb={1}>
                ユーザIDを入力
              </Text>
              <Input bg='gray.100' {...register('userId')} />
              {errors.userId && (
                <Text color='red.500' fontSize='sm'>
                  {errors.userId.message}
                </Text>
              )}
              <Text fontSize='sm' fontWeight='bold' mb={1} mt={4}>
                ユーザ名（表示名）
              </Text>
              <Input bg='gray.100' {...register('userName')} w='full' />
              {errors.userName && (
                <Text color='red.500' fontSize='sm'>
                  {errors.userName.message}
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
                <Button _hover={{ bg: 'gray.800' }} bg='black' color='white' type='submit' w='full'>
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
