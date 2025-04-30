import { Box, Flex, Avatar, Text, Grid, IconButton, VStack, GridItem } from '@chakra-ui/react';
import { useRouter } from 'next/router';
import { useContext, useEffect, useState } from 'react';
import { FaUser, FaLaptop, FaCog, FaTrophy, FaPiggyBank } from 'react-icons/fa';
import { AuthService, GetOwnInfoResponse, Role } from '@/__generated__/v1/odachin/auth_pb';
import { useClient } from '@/pages/api/ClientProvider';
import { CokiesContext } from '@/pages/api/CokiesContext';

export interface MyPageProps {}

//TODO childの色々
//TODO 使い道のチャートとか作りたい
export function MyPage({}: MyPageProps) {
  const cookies = useContext(CokiesContext);
  const router = useRouter();
  const client = useClient(AuthService);
  const [userInfo, setuserInfo] = useState<GetOwnInfoResponse | null>(null);
  useEffect(() => {
    const fetchData = async () => {
      const req = {};
      try {
        const res = await client.getOwnInfo(req);
        setuserInfo(res);
      } catch (error) {
        router.push('/login');
        alert('Login failed');
      }
    };
    fetchData();
  }, []);

  const menuItems =
    userInfo?.role === Role.PARENT
      ? [
          {
            icon: FaPiggyBank,
            label: 'お小遣い',
            onCLick: () => {
              router.push('myPage/allowance');
            },
          },
          {
            icon: FaTrophy,
            label: 'ミッションを管理',
            onCLick: () => {
              router.push('myPage/rewardSetting');
            },
          },
          {
            icon: FaUser,
            label: '家族情報',
            onCLick: () => {
              router.push('myPage/family');
            },
          },
          { icon: FaLaptop, label: '取引履歴', onCLick: () => {} },
          {
            icon: FaCog,
            label: '設定',
            onCLick: () => {
              router.push('setting/account');
            },
          },
        ]
      : [
          {
            icon: FaTrophy,
            label: 'ミッション',
            onCLick: () => {
              router.push('myPage/reward');
            },
          },
          {
            icon: FaTrophy,
            label: '使う',
            onCLick: () => {},
          },
          { icon: FaLaptop, label: '取引履歴', onCLick: () => {} },
          {
            icon: FaCog,
            label: '設定',
            onCLick: () => {
              router.push('setting/account');
            },
          },
        ];

  return (
    <Box bg='white' minH='100vh' w='100%'>
      <Box bg='yellow.400' h='100px' />

      <Flex align='center' direction='column' mt='-50px'>
        <VStack align='flex-start' maxW='800px' w='full'>
          <Avatar.Root size='2xl' top='20px'>
            <Avatar.Fallback name='Segun Adebayo' />
            <Avatar.Image src={userInfo?.avaterImageUrl} />
          </Avatar.Root>
          <Text fontSize='2xl' fontWeight='bold' mt={2}>
            {userInfo?.name}
          </Text>
        </VStack>
        <Box maxW='800px' mt={10} px={4} w='full'>
          <Text fontSize='lg' fontWeight='bold' mb={4}>
            メニュー
          </Text>
          <Grid gap={6} px={4} templateColumns={{ base: 'repeat(2, 1fr)', md: 'repeat(4, 1fr)' }}>
            {menuItems.map((item, index) => (
              <GridItem key={index} onClick={item.onCLick} textAlign='center'>
                <Flex
                  _hover={{ bg: 'gray.100', cursor: 'pointer' }}
                  align='center'
                  bg='gray.50'
                  direction='column'
                  justify='center'
                  p={4}
                  rounded='xl'
                  shadow='md'
                >
                  <IconButton
                    aria-label={item.label}
                    colorScheme='blackAlpha'
                    mb={2}
                    rounded='full'
                    size='lg'
                  >
                    <item.icon size={24} />
                  </IconButton>
                  <Text>{item.label}</Text>
                </Flex>
              </GridItem>
            ))}
          </Grid>
        </Box>
      </Flex>
    </Box>
  );
}
