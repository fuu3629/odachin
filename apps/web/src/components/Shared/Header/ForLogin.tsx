import { Menu, Avatar, Icon, Text } from '@chakra-ui/react';
import { useRouter } from 'next/router';
import { IoMdPerson, IoMdSettings, IoMdLogOut } from 'react-icons/io';

export interface ForLoginProps {
  avaterImageUrl?: string;
}

export function ForLogin({ avaterImageUrl }: ForLoginProps) {
  const router = useRouter();
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
      <Menu.Root closeOnSelect={true} positioning={{ placement: 'bottom-end' }}>
        <Menu.Trigger>
          <Avatar.Root _hover={{ cursor: 'pointer' }} size='lg'>
            <Avatar.Image src={avaterImageUrl}></Avatar.Image>
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
    </>
  );
}
