import { Button, FileUpload, Heading, Input, Text } from '@chakra-ui/react';
import { useRouter } from 'next/router';
import { useContext, useEffect, useState } from 'react';
import { HiUpload } from 'react-icons/hi';
import { useUpdateAccountForm } from './lib';
import { GetOwnInfoResponse } from '@/__generated__/v1/odachin/odachin_pb';
import { clientProvider } from '@/pages/api/ClientProvider';
import { CokiesContext } from '@/pages/api/CokiesContext';

export interface UpdateAccountFormProps {}

export function UpdateAccountForm({}: UpdateAccountFormProps) {
  const router = useRouter();
  const [defaultData, setDefaultData] = useState<GetOwnInfoResponse>();
  const {
    register,
    onSubmit,
    formState: { errors },
    watch,
  } = useUpdateAccountForm(defaultData);
  const cookies = useContext(CokiesContext);
  useEffect(() => {
    const fetchData = async () => {
      if (!cookies || !cookies.authorization) {
        console.error('No authentication token found');
        return;
      }
      const client = clientProvider();
      try {
        const res = await client.getOwnInfo(
          {},
          {
            headers: { authorization: cookies.authorization },
          },
        );
        setDefaultData(res);
      } catch (error) {
        // router.push('/login');
        console.error('Error fetching user info:', error);
      }
    };
    fetchData();
  }, []);
  console.log(watch('email'));
  return (
    <>
      <form onSubmit={onSubmit}>
        <Heading mb={6}>ユーザー情報の更新</Heading>
        <Text fontSize='sm' fontWeight='bold' mb={1}>
          メールアドレス
        </Text>
        <Input bg='gray.100' defaultValue={defaultData?.email} {...register('email')} />
        {errors.email && (
          <Text color='red.500' fontSize='sm'>
            {errors.email.message}
          </Text>
        )}
        <Text fontSize='sm' fontWeight='bold' mb={1} mt={4}>
          ユーザ名（表示名）
        </Text>
        <Input bg='gray.100' {...register('userName')} defaultValue={defaultData?.name} w='full' />
        {errors.userName && (
          <Text color='red.500' fontSize='sm'>
            {errors.userName.message}
          </Text>
        )}
        <Text h='100%' w='150px'>
          Avatar
        </Text>
        <FileUpload.Root allowDrop={true} {...register('avatar')}>
          <FileUpload.HiddenInput />
          <FileUpload.Trigger asChild bg='white'>
            <Button px='20px' size='sm' variant='outline'>
              <HiUpload /> Upload file
            </Button>
          </FileUpload.Trigger>
          <FileUpload.ItemGroup>
            <FileUpload.Items />
          </FileUpload.ItemGroup>
        </FileUpload.Root>
        {errors.userName && (
          <Text color='red.500' fontSize='sm'>
            {errors.userName.message}
          </Text>
        )}
      </form>
    </>
  );
}
