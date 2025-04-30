import { Dialog, Button, Portal, CloseButton, Center, Text, VStack } from '@chakra-ui/react';
import { Dispatch, SetStateAction, useState } from 'react';
import { RewardPeriodItem } from './RewardTable';
import { RewardService } from '@/__generated__/v1/odachin/reward_pb';
import { useClient } from '@/pages/api/ClientProvider';

export interface RewardReportDialogProps {
  item: RewardPeriodItem;
  setRefreshKey: Dispatch<SetStateAction<number>>;
}

export function RewardReportDialog({ item, setRefreshKey }: RewardReportDialogProps) {
  const [open, setOpen] = useState(false);
  const client = useClient(RewardService);
  const onCLick = async () => {
    const req = {
      rewardPeriodId: item.id,
    };
    try {
      const res = await client.reportReward(req);
      setRefreshKey((prev) => prev + 1);
      setOpen(false);
      console.log('報告完了', res);
    } catch (error) {
      console.error('Error reporting reward:', error);
    }
  };
  return (
    <>
      <Dialog.Root onOpenChange={(e) => setOpen(e.open)} open={open}>
        <Dialog.Trigger asChild>
          <Center>
            <Button
              colorPalette={
                item.status === 'COMPLETED' || item.status === 'REPORTED' ? 'gray' : 'orange'
              }
              disabled={item.status === 'COMPLETED' || item.status === 'REPORTED'}
              size='sm'
            >
              <Text fontWeight='semibold' textAlign='center'>
                {item.status === 'COMPLETED' || item.status === 'REPORTED'
                  ? item.status === 'COMPLETED'
                    ? '完了'
                    : '報告済み'
                  : '申請する'}
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
                  <Text fontSize='lg'>以下のミッションの完了を報告しますか？</Text>
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
                <Button onClick={onCLick}>報告する</Button>
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
