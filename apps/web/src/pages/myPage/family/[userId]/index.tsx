import { GetStaticPaths, GetStaticProps } from 'next';
import { FamilyUserPage } from '@/components/MyPage/Family/FamilyUserPage/FamilyUserPage';

interface Props {
  userId: string;
}
//TODO　これつくる
export default function FamilyUser({ userId }: Props) {
  return <FamilyUserPage userId={userId}></FamilyUserPage>;
}

export const getStaticPaths: GetStaticPaths = async () => {
  return {
    paths: [], // 静的にビルドしたいパス（空でも fallback を使えばOK）
    fallback: 'blocking', // リクエスト時に生成される
  };
};

export const getStaticProps: GetStaticProps = async (context) => {
  const { userId } = context.params!;
  return {
    props: {
      userId,
    },
  };
};
