import { Avatar, VStack, Text } from '@chakra-ui/react';
import { useRouter } from 'next/router';

export interface FamilyAvatarProps {
  avatarUrl?: string;
  userName: string;
  userId: string;
}

export function FamilyAvatar({ avatarUrl, userName, userId }: FamilyAvatarProps) {
  const router = useRouter();
  return (
    <>
      <VStack gapY={4} w='156px'>
        <Avatar.Root
          _hover={{ cursor: 'pointer' }}
          onClick={() => {
            router.push(`/myPage/family/${userId}`);
          }}
          size='lg'
        >
          <Avatar.Image src={avatarUrl}></Avatar.Image>
        </Avatar.Root>
        <Text fontWeight='semibold' textStyle='md'>
          {userName}
        </Text>
      </VStack>
    </>
  );
}
