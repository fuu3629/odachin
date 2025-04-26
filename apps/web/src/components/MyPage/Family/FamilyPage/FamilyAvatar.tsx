import { Avatar, VStack, Text, SystemStyleObject } from '@chakra-ui/react';
import { useRouter } from 'next/router';

export interface FamilyAvatarProps {
  avatarUrl?: string;
  userName: string;
  userId: string;
  onClick?: () => void;
  css?: SystemStyleObject | Omit<(SystemStyleObject | undefined)[], keyof any[]> | undefined;
  colorPalette?: string;
}

export function FamilyAvatar({
  avatarUrl,
  userName,
  userId,
  onClick,
  css,
  colorPalette,
}: FamilyAvatarProps) {
  const router = useRouter();
  return (
    <>
      <VStack gapY={4} w='156px'>
        <Avatar.Root
          _hover={{ cursor: 'pointer' }}
          css={css}
          onClick={() => {
            onClick?.();
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
