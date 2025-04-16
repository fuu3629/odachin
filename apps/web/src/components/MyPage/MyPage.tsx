import { Box, Flex, Avatar, Text, Grid, IconButton, VStack, GridItem } from '@chakra-ui/react';
import { useContext, useEffect, useState } from 'react';
import { FaUser, FaLaptop, FaCog } from 'react-icons/fa';
import { GetOwnInfoResponse } from '@/__generated__/v1/odachin/odachin_pb';
import { clientProvider } from '@/pages/api/ClientProvider';
import { CokiesContext } from '@/pages/api/CokiesContext';
export interface MyPageProps {}

//TODO ROLEで分岐する
export function MyPage({}: MyPageProps) {
  const cookies = useContext(CokiesContext);
  const [userInfo, setuserInfo] = useState<GetOwnInfoResponse | null>(null);
  useEffect(() => {
    if (!cookies || !cookies.auth) {
      //TODO 401エラー
      console.error('No authentication token found');
      return;
    }
    const fetchData = async () => {
      const client = clientProvider();
      //TODO connectの型付け調べる
      const req = {};
      try {
        const res = await client.getOwnInfo(req, {
          headers: { authorization: cookies.authorization },
        });
        setuserInfo(res);
      } catch (error) {
        // TODO login画面返す
        console.error('Error fetching user info:', error);
      }
    };
    fetchData();
  }, []);

  const menuItems = [
    { icon: FaUser, label: '家族情報' },
    { icon: FaLaptop, label: '取引履歴' },
    { icon: FaCog, label: '設定' },
  ];

  return (
    <Box bg='white' minH='100vh' w='100%'>
      <Box bg='yellow.400' h='100px' />

      <Flex align='center' direction='column' mt='-50px'>
        <VStack align='flex-start' maxW='800px' w='full'>
          <Avatar.Root size='2xl' top='20px'>
            <Avatar.Fallback name='Segun Adebayo' />
            <Avatar.Image src='https://bit.ly/sage-adebayo' />
          </Avatar.Root>
          <Text fontSize='2xl' fontWeight='bold' mt={2}>
            {userInfo?.name}
          </Text>
        </VStack>
        <Box maxW='800px' mt={10} px={4} w='full'>
          <Text fontSize='lg' fontWeight='bold' mb={4}>
            マイメニュー
          </Text>
          <Grid gap={6} px={4} templateColumns={{ base: 'repeat(2, 1fr)', md: 'repeat(4, 1fr)' }}>
            {menuItems.map((item, index) => (
              <GridItem key={index} textAlign='center'>
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
          ;
        </Box>
      </Flex>
    </Box>
  );
}
