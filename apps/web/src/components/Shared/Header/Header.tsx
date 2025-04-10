import { HStack, Spacer, Text } from '@chakra-ui/react';
import { Pacifico } from 'next/font/google';
import { useRouter } from 'next/router';
import { ForLogin } from './ForLogin';
import { ForUnLogin } from './ForUnLogin';

export interface HeaderProps {}

const pacifico = Pacifico({
  subsets: ['latin'],
  weight: '400',
});

export function Header({}: HeaderProps) {
  const router = useRouter();
  const { pathname } = router;
  const paths = ['/login', '/createNewAccount', '/'];

  return (
    <>
      <HStack bg='white' h='64px' pl={12} position='sticky' pr={6} top={0} w='100%'>
        <Text className={pacifico.className} fontSize='4xl'>
          Odachin
        </Text>
        <Text fontSize='2xl' fontWeight='semibold' ml={4}>
          家庭に安全にお小遣いを導入できるアプリ
        </Text>
        <Spacer></Spacer>
        {paths.includes(pathname) ? <ForLogin></ForLogin> : <ForUnLogin></ForUnLogin>}
      </HStack>
    </>
  );
}
