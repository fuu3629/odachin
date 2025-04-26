import { Dialog, Button, Portal, CloseButton, DataList, VStack, Text } from '@chakra-ui/react';
import { useContext } from 'react';
import { RewardItem } from './RewardSettingTable';
import { RewardService } from '@/__generated__/v1/odachin/reward_pb';
import { useClient } from '@/pages/api/ClientProvider';
import { CokiesContext } from '@/pages/api/CokiesContext';

export interface RewardDeleteDiadlogProps {
  rewardItem?: RewardItem;
}

export function RewardDeleteDiadlog({ rewardItem }: RewardDeleteDiadlogProps) {
  const client = useClient(RewardService);
  const cookies = useContext(CokiesContext);
  const onClick = () => {
    if (!rewardItem) return;
    if (!cookies || !cookies.authorization) {
      console.error('No authentication token found');
      return;
    }
    const req = { rewardId: rewardItem.id };
    client.deleteReward(req, { headers: { authorization: cookies?.authorization } }).then((res) => {
      console.log('Reward deleted:', res);
    });
  };
  return (
    <>
      <Dialog.Root>
        <Dialog.Trigger asChild>
          <Button colorPalette='red' variant='surface'>
            削除
          </Button>
        </Dialog.Trigger>
        <Portal>
          <Dialog.Backdrop />
          <Dialog.Positioner>
            <Dialog.Content>
              <Dialog.Header>
                <Dialog.Title>Dialog Title</Dialog.Title>
              </Dialog.Header>
              <Dialog.Body>
                <VStack alignItems='flex-start'>
                  <DataList.Root variant='bold'>
                    <DataList.Item>
                      <DataList.ItemLabel>
                        <Text fontWeight='semibold' textStyle='md'>
                          お子様の名前
                        </Text>
                      </DataList.ItemLabel>
                      <DataList.ItemValue>{rewardItem?.name}</DataList.ItemValue>
                    </DataList.Item>
                    <DataList.Item>
                      <DataList.ItemLabel>
                        <Text fontWeight='semibold' textStyle='md'>
                          タイトル
                        </Text>
                      </DataList.ItemLabel>
                      <DataList.ItemValue>{rewardItem?.title}</DataList.ItemValue>
                    </DataList.Item>
                    <DataList.Item>
                      <DataList.ItemLabel>
                        <Text fontWeight='semibold' textStyle='md'>
                          説明
                        </Text>
                      </DataList.ItemLabel>
                      <DataList.ItemValue>{rewardItem?.description}</DataList.ItemValue>
                    </DataList.Item>
                    <DataList.Item>
                      <DataList.ItemLabel>
                        <Text fontWeight='semibold' textStyle='md'>
                          お小遣い
                        </Text>
                      </DataList.ItemLabel>
                      <DataList.ItemValue>{rewardItem?.amount}</DataList.ItemValue>
                    </DataList.Item>
                  </DataList.Root>
                </VStack>
              </Dialog.Body>
              <Dialog.Footer>
                <Dialog.ActionTrigger asChild>
                  <Button variant='outline'>Cancel</Button>
                </Dialog.ActionTrigger>
                <Button colorPalette='red' onClick={onClick} variant='surface'>
                  削除
                </Button>
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
