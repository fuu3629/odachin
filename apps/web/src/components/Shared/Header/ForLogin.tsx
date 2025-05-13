import { Menu, Avatar, Icon, Text, HStack } from '@chakra-ui/react';
import { useRouter } from 'next/router';
import { useContext, useEffect, useState } from 'react';
import { IoMdPerson, IoMdSettings, IoMdLogOut } from 'react-icons/io';
import { AuthService, GetOwnInfoResponse } from '@/__generated__/v1/odachin/auth_pb';
import { useClient } from '@/pages/api/ClientProvider';
import { CokiesContext } from '@/pages/api/CokiesContext';

export interface ForLoginProps {}

export function ForLogin({}: ForLoginProps) {
  const cookies = useContext(CokiesContext);
  const client = useClient(AuthService);
  const router = useRouter();
  const [userInfo, setuserInfo] = useState<GetOwnInfoResponse | null>(null);
  const { pathname } = router;
  useEffect(() => {
    console.log(cookies?.authorization);
    if (!cookies || !cookies.authorization) {
      router.push('/login');
      console.error('No authentication token found');
      return;
    }
    const fetchData = async () => {
      const req = {};
      try {
        const res = await client.getOwnInfo(req);
        setuserInfo(res);
      } catch (error) {
        router.push('/login');
        console.error('Error fetching user info:', error);
      }
    };
    fetchData();
  }, [pathname]);
  const onCLickMyPage = () => {
    router.push('/myPage');
  };
  const onCLickSetting = () => {
    router.push('/setting');
  };
  const onCLickLogout = () => {
    router.push('/login');
  };
  return (
    <>
      <HStack>
        <Text fontSize='xl' fontWeight='semibold' mr={4}>
          {userInfo?.name}さん
        </Text>
        <Menu.Root closeOnSelect={true} positioning={{ placement: 'bottom-end' }}>
          <Menu.Trigger>
            <Avatar.Root _hover={{ cursor: 'pointer' }} size='lg'>
              <Avatar.Image src={userInfo?.avaterImageUrl}></Avatar.Image>
            </Avatar.Root>
          </Menu.Trigger>
          <Menu.Positioner>
            <Menu.Content p={2} w='250px'>
              <Menu.Item
                _hover={{ cursor: 'pointer' }}
                onClick={onCLickMyPage}
                px={3}
                py={2}
                value='myPage'
              >
                <Icon>
                  <IoMdPerson />
                </Icon>
                <Text textStyle='md'>マイページ</Text>
              </Menu.Item>
              <Menu.Item
                _hover={{ cursor: 'pointer' }}
                onClick={onCLickSetting}
                px={3}
                py={2}
                value='setting'
              >
                <Icon>
                  <IoMdSettings />
                </Icon>
                <Text textStyle='md'>設定</Text>
              </Menu.Item>
              <Menu.Item
                _hover={{ cursor: 'pointer' }}
                onClick={onCLickLogout}
                px={3}
                py={2}
                value='logout'
              >
                <Icon>
                  <IoMdLogOut />
                </Icon>
                <Text textStyle='md'>ログアウト</Text>
              </Menu.Item>
            </Menu.Content>
          </Menu.Positioner>
        </Menu.Root>
      </HStack>
    </>
  );
}
