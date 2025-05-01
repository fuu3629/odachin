import { Dialog, Center, Button, Portal, VStack, CloseButton, Text } from '@chakra-ui/react';
import { Dispatch, SetStateAction, useState } from 'react';
import { RewardInfo, RewardService } from '@/__generated__/v1/odachin/reward_pb';
import { useClient } from '@/pages/api/ClientProvider';

export interface RewardApproveDialogProps {
  item: RewardInfo;
  setRefreshKey: Dispatch<SetStateAction<number>>;
}

export function RewardApproveDialog({ item, setRefreshKey }: RewardApproveDialogProps) {
  const [open, setOpen] = useState(false);
  const rewardClient = useClient(RewardService);
  const onClick = async () => {
    const req = {
      rewardPeriodId: item.rewardPeriodId,
    };
    try {
      const res = await rewardClient.approveReward(req);
      setRefreshKey((prev) => prev + 1);
      setOpen(false);
      console.log('報告完了', req);
    } catch (error) {
      console.error('Error reporting reward:', error);
    }
  };
  return (
    <>
      <Dialog.Root onOpenChange={(e) => setOpen(e.open)} open={open}>
        <Dialog.Trigger asChild>
          <Center>
            <Button colorPalette='orange' size='sm'>
              <Text fontWeight='semibold' textAlign='center'>
                承認する
              </Text>
            </Button>
          </Center>
        </Dialog.Trigger>
        <Portal>
          <Dialog.Backdrop />
          <Dialog.Positioner>
            <Dialog.Content>
              <Dialog.Header>
                <Dialog.Title>ミッションの報告</Dialog.Title>
              </Dialog.Header>
              <Dialog.Body>
                <VStack alignItems='start' gap={8}>
                  <Text fontSize='lg'>以下のミッションを承認しますか？</Text>
                  <Text fontSize='md' fontWeight='semibold'>
                    タイトル： {item.title}
                  </Text>
                  <Text fontSize='md'>説明： {item.description}</Text>
                </VStack>
              </Dialog.Body>
              <Dialog.Footer>
                <Dialog.ActionTrigger asChild>
                  <Button variant='outline'>キャンセル</Button>
                </Dialog.ActionTrigger>
                <Button onClick={onClick}>承認する</Button>
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
