import { Button, Center, Link, Table, Text } from '@chakra-ui/react';
import { useContext, useEffect, useState } from 'react';
import { CiWarning } from 'react-icons/ci';
import { CompletedTag } from './CompletedTag';
import { Reward_Type, RewardService } from '@/__generated__/v1/odachin/reward_pb';
import { useClient } from '@/pages/api/ClientProvider';
import { CokiesContext } from '@/pages/api/CokiesContext';
import unauthorizedPage from '@/pages/unauthorized';

export interface RewardTableProps {
  rewardType: Reward_Type;
}

export interface RewardPeriodItem {
  id: bigint;
  title: string;
  description: string;
  amount: number;
  isCompleted: boolean;
}

export function RewardTable({ rewardType }: RewardTableProps) {
  const client = useClient(RewardService);
  const cookies = useContext(CokiesContext);
  const [items, setItems] = useState<RewardPeriodItem[]>([]);
  console.log('rewardType;Page', rewardType, typeof rewardType);

  useEffect(() => {
    const fetchData = async () => {
      if (!cookies || !cookies.authorization) {
        unauthorizedPage();
        console.error('No authentication token found');
        return;
      }
      const res = await client.getRewardList(
        { rewardType: rewardType },
        {
          headers: { authorization: cookies?.authorization },
        },
      );
      console.log('res', res);
      const rewards = res.rewardList.map((reward) => {
        return {
          id: reward.rewardPeriodId,
          title: reward.title,
          description: reward.description,
          amount: reward.amount,
          isCompleted: reward.isCompleted,
        };
      });
      setItems(rewards);
    };
    try {
      fetchData();
    } catch (error) {
      console.error('Error fetching reward list:', error);
    }
  }, [rewardType]);

  return (
    <>
      <Table.Root borderColor='yellow.400' borderTopRadius={24} showColumnBorder variant='outline'>
        <Table.Header bgColor='yellow.400' borderColor='yellow.400' borderXWidth={4}>
          <Table.Row>
            <Table.ColumnHeader textAlign='center' w='25%'>
              <Text fontSize='lg' fontWeight='semibold'>
                タイトル
              </Text>
            </Table.ColumnHeader>
            <Table.ColumnHeader textAlign='center' w='40%'>
              <Text fontSize='lg' fontWeight='semibold'>
                説明
              </Text>
            </Table.ColumnHeader>
            <Table.ColumnHeader textAlign='center'>
              <Text fontSize='lg' fontWeight='semibold'>
                おこづかい
              </Text>
            </Table.ColumnHeader>
            <Table.ColumnHeader textAlign='center'>
              <Text fontSize='lg' fontWeight='semibold'>
                状態
              </Text>
            </Table.ColumnHeader>
            <Table.ColumnHeader textAlign='center'>
              <Text fontSize='lg' fontWeight='semibold'>
                申請
              </Text>
            </Table.ColumnHeader>
          </Table.Row>
        </Table.Header>
        <Table.Body bgColor='white'>
          {items.length === 0 ? (
            <Table.Row>
              <Table.Cell colSpan={4} textAlign='center'>
                <Link>
                  <CiWarning size='1.5em' />
                  <Text fontSize='lg' fontWeight='semibold'>
                    まだミッションはありません
                  </Text>
                </Link>
              </Table.Cell>
            </Table.Row>
          ) : (
            <>
              {items.map((item) => (
                <Table.Row borderColor='yellow.400' borderWidth={4} key={item.id}>
                  <Table.Cell>
                    <Text fontWeight='semibold' lineClamp='1'>
                      {item.title}
                    </Text>
                  </Table.Cell>
                  <Table.Cell>
                    <Text fontWeight='semibold' lineClamp='1'>
                      {item.description}
                    </Text>
                  </Table.Cell>
                  <Table.Cell>
                    <Text fontWeight='semibold' textAlign='center'>
                      {item.amount}
                    </Text>
                  </Table.Cell>
                  <Table.Cell>
                    <CompletedTag isCompleted={item.isCompleted}></CompletedTag>
                  </Table.Cell>
                  <Table.Cell>
                    <Center>
                      <Button
                        colorPalette={item.isCompleted ? 'gray' : 'orange'}
                        disabled={item.isCompleted}
                        size='sm'
                      >
                        <Text fontWeight='semibold' textAlign='center'>
                          {item.isCompleted ? '完了済み' : '申請する'}
                        </Text>
                      </Button>
                    </Center>
                  </Table.Cell>
                </Table.Row>
              ))}
            </>
          )}
        </Table.Body>
      </Table.Root>
    </>
  );
}
