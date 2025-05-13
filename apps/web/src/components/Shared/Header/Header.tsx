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

  const onCLick = () => {
    if (paths.includes(pathname)) {
      router.push('/');
    } else {
      router.push('/myPage');
    }
  };

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
        zIndex={1000001}
      >
        <Text
          _hover={{ cursor: 'pointer' }}
          className={pacifico.className}
          fontSize='4xl'
          onClick={onCLick}
        >
          Odachin
        </Text>
        <Text color='white' fontSize='2xl' fontWeight='semibold' ml={4}>
          家庭に安全にお小遣いを導入できるアプリ
        </Text>
        <Spacer></Spacer>
        {paths.includes(pathname) ? <ForUnLogin></ForUnLogin> : <ForLogin></ForLogin>}
      </HStack>
    </>
  );
}
