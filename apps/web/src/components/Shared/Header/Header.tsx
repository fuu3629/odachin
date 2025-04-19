import { HStack, Spacer, Text } from '@chakra-ui/react';
import { Pacifico } from 'next/font/google';
import { useRouter } from 'next/router';
import { useContext, useState, useEffect } from 'react';
import { ForLogin } from './ForLogin';
import { ForUnLogin } from './ForUnLogin';
import { GetOwnInfoResponse } from '@/__generated__/v1/odachin/odachin_pb';
import { clientProvider } from '@/pages/api/ClientProvider';
import { CokiesContext } from '@/pages/api/CokiesContext';

export interface HeaderProps {}

const pacifico = Pacifico({
  subsets: ['latin'],
  weight: '400',
});

export function Header({}: HeaderProps) {
  const cookies = useContext(CokiesContext);
  const router = useRouter();
  const [userInfo, setuserInfo] = useState<GetOwnInfoResponse | null>(null);
  useEffect(() => {
    if (!cookies || !cookies.authorization) {
      console.error('No authentication token found');
      return;
    }
    const fetchData = async () => {
      const client = clientProvider();
      const req = {};
      try {
        const res = await client.getOwnInfo(req, {
          headers: { authorization: cookies.authorization },
        });
        console.log('res', res);
        setuserInfo(res);
      } catch (error) {
        router.push('/login');
        console.error('Error fetching user info:', error);
      }
    };
    fetchData();
  }, []);
  const { pathname } = router;
  const paths = ['/login', '/createNewAccount', '/'];

  return (
    <>
      <HStack
        bg='white'
        bgColor='orange.400'
        borderBottomColor='gray.200'
        borderBottomWidth={1}
        h='64px'
        pl={12}
        position='sticky'
        pr={6}
        top={0}
        w='100%'
      >
        <Text className={pacifico.className} fontSize='4xl'>
          Odachin
        </Text>
        <Text color='white' fontSize='2xl' fontWeight='semibold' ml={4}>
          家庭に安全にお小遣いを導入できるアプリ
        </Text>
        <Spacer></Spacer>
        {paths.includes(pathname) ? (
          <ForUnLogin></ForUnLogin>
        ) : (
          <HStack>
            <Text fontSize='xl' fontWeight='semibold' mr={4}>
              {userInfo?.name}さん
            </Text>
            <ForLogin avaterImageUrl={userInfo?.avaterImageUrl}></ForLogin>
          </HStack>
        )}
      </HStack>
    </>
  );
}
