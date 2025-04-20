import {
  Avatar,
  Box,
  Button,
  Circle,
  FileUpload,
  Flex,
  Float,
  Heading,
  Input,
  SimpleGrid,
  Spacer,
  Text,
  VStack,
} from '@chakra-ui/react';
import { useContext, useEffect } from 'react';
import { MdEdit } from 'react-icons/md';
import { useUpdateAccountForm } from './lib';
import { AuthService } from '@/__generated__/v1/odachin/auth_pb';
import { useClient } from '@/pages/api/ClientProvider';
import { CokiesContext } from '@/pages/api/CokiesContext';

export interface UpdateAccountFormProps {}

export function UpdateAccountForm({}: UpdateAccountFormProps) {
  const {
    register,
    onSubmit,
    formState: { errors },
    watch,
    reset,
  } = useUpdateAccountForm();
  const cookies = useContext(CokiesContext);
  const client = useClient(AuthService);
  const file = watch('avatar');
  useEffect(() => {
    const fetchData = async () => {
      if (!cookies || !cookies.authorization) {
        console.error('No authentication token found');
        return;
      }
      try {
        const res = await client.getOwnInfo(
          {},
          {
            headers: { authorization: cookies.authorization },
          },
        );
        reset({
          userName: res.name,
          email: res.email,
          avatar: undefined,
        });
      } catch (error) {
        // router.push('/login');
        console.error('Error fetching user info:', error);
      }
    };
    fetchData();
  }, []);
  return (
    <Box bg='white' borderRadius='xl' boxShadow='lg' m='auto' p={8} w='40%'>
      <form onSubmit={onSubmit}>
        <VStack>
          <FileUpload.Root allowDrop={true} {...register('avatar')} alignItems='center'>
            <FileUpload.HiddenInput />
            <FileUpload.Trigger asChild bg='white'>
              <Flex align='center' direction='column' mb={8}>
                <Avatar.Root onClick={() => {}} size='2xl' top='20px'>
                  <Avatar.Fallback name='Segun Adebayo' />
                  <Avatar.Image src='https://bit.ly/sage-adebayo' />
                  <Float offsetX='1' offsetY='1' placement='bottom-end'>
                    <Circle bg='orange.500' outline='0.2em solid' outlineColor='bg' size='23px'>
                      <MdEdit color='white' />
                    </Circle>
                  </Float>
                </Avatar.Root>
              </Flex>
            </FileUpload.Trigger>
            <FileUpload.ItemGroup>
              <FileUpload.Items />
            </FileUpload.ItemGroup>
          </FileUpload.Root>

          <Heading mb={6}>ユーザー情報の更新</Heading>
          <SimpleGrid columns={2} gap={10}>
            <Box>
              <Text fontSize='sm' fontWeight='bold'>
                メールアドレス
              </Text>
              <Input bg='gray.100' {...register('email')} />
              {errors.email && (
                <Text color='red.500' fontSize='sm'>
                  {errors.email.message}
                </Text>
              )}
            </Box>
            <Box>
              <Text fontSize='sm' fontWeight='bold'>
                ユーザ名（表示名）
              </Text>
              <Input bg='gray.100' {...register('userName')} w='full' />
              {errors.userName && (
                <Text color='red.500' fontSize='sm'>
                  {errors.userName.message}
                </Text>
              )}
            </Box>
          </SimpleGrid>
          <Spacer />
          <Button type='submit'>Update</Button>
        </VStack>
      </form>
    </Box>
  );
}
