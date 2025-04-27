import { Center, Link, Table, Tabs, Text } from '@chakra-ui/react';
import { useContext, useEffect, useState } from 'react';
import { CiWarning } from 'react-icons/ci';
import { RewardFloat } from '../../Reward/RewardPage';
import { RewardAddDialog } from './RewardAddDialog';
import { RewardDeleteDiadlog } from './RewardDeleteDiadlog';
import { SelectedUser } from './RewardSettingPage';
import { Reward_Type, RewardService } from '@/__generated__/v1/odachin/reward_pb';
import { useClient } from '@/pages/api/ClientProvider';
import { CokiesContext } from '@/pages/api/CokiesContext';

export interface RewardSettingTableProps {
  selectedUser?: SelectedUser;
}

export interface RewardItem {
  id: bigint;
  title: string;
  description: string;
  name?: string;
  amount: number;
}

export function RewardSettingTable({ selectedUser }: RewardSettingTableProps) {
  const client = useClient(RewardService);
  const cookies = useContext(CokiesContext);
  const [rewardType, setRewardType] = useState<Reward_Type>(Reward_Type.DAILY);
  const [items, setItems] = useState<RewardItem[]>([]);
  const [refreshKey, setRefreshKey] = useState(0);

  useEffect(() => {
    const fetchData = async () => {
      if (!cookies || !cookies.authorization) {
        console.error('No authentication token found');
        return;
      }
      try {
        const req = { childId: selectedUser?.id, rewardType: rewardType };
        const res2 = await client.getChildRewardList(req, {
          headers: { authorization: cookies?.authorization },
        });
        const rewards = res2.rewardList.map((reward) => {
          return {
            id: reward.rewardPeriodId,
            title: reward.title,
            description: reward.description,
            amount: reward.amount,
            name: selectedUser?.name,
          };
        });
        setItems(rewards);
      } catch (error) {
        console.error('Error fetching reward list:', error);
      }
    };
    fetchData();
  }, [selectedUser, rewardType, refreshKey]);
  return (
    <>
      <Tabs.Root
        colorPalette='orange'
        defaultValue={Reward_Type.DAILY.toString()}
        fitted
        onValueChange={(e) => setRewardType(Number(e.value) as Reward_Type)}
        variant='subtle'
      >
        <Tabs.List>
          <Tabs.Trigger borderTopRadius='24px' value={Reward_Type.DAILY.toString()}>
            <RewardFloat count={0}></RewardFloat>
            <Text fontWeight='semibold' textStyle='2xl'>
              毎日のミッション
            </Text>
          </Tabs.Trigger>
          <Tabs.Trigger borderTopRadius={24} value={Reward_Type.WEEKLY.toString()}>
            <RewardFloat count={0}></RewardFloat>
            <Text fontWeight='semibold' textStyle='2xl'>
              毎週のミッション
            </Text>
          </Tabs.Trigger>
          <Tabs.Trigger borderTopRadius={24} value={Reward_Type.MONTHLY.toString()}>
            <RewardFloat count={0}></RewardFloat>
            <Text fontWeight='semibold' textStyle='2xl'>
              毎月のミッション
            </Text>
          </Tabs.Trigger>
        </Tabs.List>
      </Tabs.Root>
      <Table.Root borderColor='yellow.400' showColumnBorder variant='outline'>
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
                削除
              </Text>
            </Table.ColumnHeader>
          </Table.Row>
        </Table.Header>
        <Table.Body bgColor='white'>
          {items.length === 0 ? (
            <Table.Row>
              <Table.Cell borderColor='yellow.400' borderWidth={4} colSpan={4} textAlign='center'>
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
                    <Center>
                      <RewardDeleteDiadlog rewardItem={item}></RewardDeleteDiadlog>
                    </Center>
                  </Table.Cell>
                </Table.Row>
              ))}
            </>
          )}
          <Table.Row borderColor='yellow.400' borderWidth={4}>
            <Table.Cell borderColor='yellow.400' borderWidth={4} colSpan={4} textAlign='center'>
              <RewardAddDialog
                rewardType={rewardType}
                setRefreshKey={setRefreshKey}
                toUser={selectedUser}
              ></RewardAddDialog>
            </Table.Cell>
          </Table.Row>
        </Table.Body>
      </Table.Root>
    </>
  );
}
