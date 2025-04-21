import {
  Dialog,
  Button,
  Portal,
  CloseButton,
  IconButton,
  VStack,
  Text,
  Input,
  HStack,
  Avatar,
  Box,
} from '@chakra-ui/react';
import { useRouter } from 'next/router';
import { useContext, useState } from 'react';
import { IoAdd } from 'react-icons/io5';
import { AuthService, GetUserInfoResponse } from '@/__generated__/v1/odachin/auth_pb';
import { FamilyService } from '@/__generated__/v1/odachin/faimily_pb';
import { useClient } from '@/pages/api/ClientProvider';
import { CokiesContext } from '@/pages/api/CokiesContext';

export interface AddFamilyDialogProps {}

export function AddFamilyDialog({}: AddFamilyDialogProps) {
  const client = useClient(FamilyService);
  const authClient = useClient(AuthService);
  const cookies = useContext(CokiesContext);
  const router = useRouter();
  const [userId, setUserId] = useState<string>('');
  const [confirmUserInfo, setConfirmUserInfo] = useState<GetUserInfoResponse>();
  const onCLick = async () => {
    if (!cookies || !cookies.authorization) {
      console.error('No authentication token found');
      return;
    }
    const req = {
      toUserId: userId,
    };
    try {
      await client.inviteUser(req, {
        headers: { authorization: cookies.authorization },
      });
    } catch (error) {
      console.error('Error inviting family member:', error);
    }
  };
  const onClickConfirm = async () => {
    if (!cookies || !cookies.authorization) {
      console.error('No authentication token found');
      return;
    }
    const req = {
      userId: userId,
    };
    try {
      const res = await authClient.getUserInfo(req, {
        headers: { authorization: cookies.authorization },
      });
      setConfirmUserInfo(res);
    } catch (error) {
      console.error('Error inviting family member:', error);
    }
  };
  return (
    <>
      <Dialog.Root>
        <Dialog.Trigger asChild>
          <VStack gapY={4} w='156px'>
            <IconButton
              aria-label='Add Family'
              borderColor='black'
              rounded='full'
              size='lg'
              variant='surface'
            >
              <IoAdd />
            </IconButton>
            <Text fontWeight='semibold' textStyle='md'>
              ユーザーを追加する
            </Text>
          </VStack>
        </Dialog.Trigger>
        <Portal>
          <Dialog.Backdrop />
          <Dialog.Positioner>
            <Dialog.Content>
              <Dialog.Header>
                <Dialog.Title>ユーザーを追加する</Dialog.Title>
              </Dialog.Header>
              <Dialog.Body>
                <Text fontSize='sm' fontWeight='bold' mb={1}>
                  お子様のユーザIDを入力してください
                </Text>
                <HStack>
                  <Input
                    bg='gray.100'
                    onChange={(e) => {
                      setUserId(e.target.value);
                    }}
                    placeholder='example userId'
                  />
                  <Button onClick={onClickConfirm}>ユーザーを確認</Button>
                </HStack>
                {confirmUserInfo && (
                  <>
                    <VStack mt={4}>
                      <Text fontSize='sm' fontWeight='bold'>
                        ユーザー情報
                      </Text>
                      <Avatar.Root>
                        <Avatar.Image alt='User Avatar' src={confirmUserInfo.avatarImageUrl} />
                        <Avatar.Fallback>?</Avatar.Fallback>
                      </Avatar.Root>
                      <Box>
                        <Text fontSize='sm'>ユーザーID: {confirmUserInfo.userId}</Text>
                        <Text fontSize='sm'>名前: {confirmUserInfo.name}</Text>
                      </Box>
                    </VStack>
                    <Text color='red.600' fontSize='sm' fontWeight='bold' mt={4}>
                      間違いなければ招待ボタンを押してください
                    </Text>
                  </>
                )}
              </Dialog.Body>
              <Dialog.Footer>
                <Dialog.ActionTrigger asChild>
                  <Button variant='outline'>キャンセル</Button>
                </Dialog.ActionTrigger>
                <Button onClick={onCLick}>招待</Button>
              </Dialog.Footer>
              <Dialog.CloseTrigger asChild>
                <CloseButton size='sm' />
              </Dialog.CloseTrigger>
            </Dialog.Content>
          </Dialog.Positioner>
        </Portal>
      </Dialog.Root>
    </>
  );
}
