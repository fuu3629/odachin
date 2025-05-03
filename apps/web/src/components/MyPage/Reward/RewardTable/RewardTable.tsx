import { Link, Table, Text } from '@chakra-ui/react';
import dayjs from 'dayjs';
import relativeTime from 'dayjs/plugin/relativeTime';
import { Dispatch, SetStateAction, useContext, useEffect, useState } from 'react';
import { CiWarning } from 'react-icons/ci';
import { CompletedTag } from './CompletedTag';
import { RewardReportDialog } from './RewardReportDialog';
import { Reward_Type, RewardService } from '@/__generated__/v1/odachin/reward_pb';
import { useClient } from '@/pages/api/ClientProvider';
import { CokiesContext } from '@/pages/api/CokiesContext';
import unauthorizedPage from '@/pages/unauthorized';
dayjs.extend(relativeTime);
import 'dayjs/locale/ja';
dayjs.locale('ja');

export interface RewardTableProps {
  rewardType: Reward_Type;
  refreshKey: number;
  setRefreshKey: Dispatch<SetStateAction<number>>;
}

export interface RewardPeriodItem {
  id: bigint;
  title: string;
  description: string;
  amount: number;
  status: string;
  endDate: dayjs.Dayjs;
}

export function RewardTable({ rewardType, refreshKey, setRefreshKey }: RewardTableProps) {
  const client = useClient(RewardService);
  const cookies = useContext(CokiesContext);
  const [items, setItems] = useState<RewardPeriodItem[]>([]);

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
      const rewards = res.rewardList.map((reward) => {
        return {
          id: reward.rewardPeriodId,
          title: reward.title,
          description: reward.description,
          amount: reward.amount,
          status: reward.status,
          endDate: dayjs.unix(Number(reward.endDate?.seconds)),
        };
      });
      setItems(rewards);
    };
    try {
      fetchData();
    } catch (error) {
      console.error('Error fetching reward list:', error);
    }
  }, [rewardType, refreshKey]);

  return (
    <>
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
            <Table.ColumnHeader textAlign='center' w='10%'>
              <Text fontSize='lg' fontWeight='semibold'>
                おこづかい
              </Text>
            </Table.ColumnHeader>
            <Table.ColumnHeader textAlign='center' w='10%'>
              <Text fontSize='lg' fontWeight='semibold'>
                残り時間
              </Text>
            </Table.ColumnHeader>
            <Table.ColumnHeader textAlign='center' w='10%'>
              <Text fontSize='lg' fontWeight='semibold'>
                状態
              </Text>
            </Table.ColumnHeader>
            <Table.ColumnHeader textAlign='center' w='15%'>
              <Text fontSize='lg' fontWeight='semibold'>
                申請
              </Text>
            </Table.ColumnHeader>
          </Table.Row>
        </Table.Header>
        <Table.Body bgColor='white'>
          {items.length === 0 ? (
            <Table.Row>
              <Table.Cell colSpan={5} textAlign='center'>
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
                    <Text fontWeight='semibold' textAlign='center'>
                      {item.endDate.fromNow()}
                    </Text>
                  </Table.Cell>
                  <Table.Cell>
                    <CompletedTag status={item.status}></CompletedTag>
                  </Table.Cell>
                  <Table.Cell>
                    <RewardReportDialog
                      item={item}
                      setRefreshKey={setRefreshKey}
                    ></RewardReportDialog>
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
