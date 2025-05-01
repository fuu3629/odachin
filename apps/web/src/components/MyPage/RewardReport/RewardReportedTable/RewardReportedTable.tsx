import { Center, Table, Text } from '@chakra-ui/react';
import dayjs from 'dayjs';
import { useState, useEffect } from 'react';
import { CiWarning } from 'react-icons/ci';
import { CompletedTag } from '../../Reward/RewardTable';
import { RewardApproveDialog } from './RewardApproveDialog';
import { RewardInfo, RewardService } from '@/__generated__/v1/odachin/reward_pb';
import { useClient } from '@/pages/api/ClientProvider';

export interface RewardReportedTableProps {}

export function RewardReportedTable({}: RewardReportedTableProps) {
  const rewardClient = useClient(RewardService);
  const [refreshKey, setRefreshKey] = useState(0);
  const [items, setItems] = useState<RewardInfo[]>([]);

  useEffect(() => {
    const fetchData = async () => {
      try {
        const res = await rewardClient.getReportedRewardList({});
        const rewardList = res.rewardList;
        setItems(rewardList);
      } catch (error) {
        console.error('Error fetching reward list:', error);
      }
    };
    fetchData();
  }, [refreshKey]);
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
                <Center>
                  <CiWarning size='1.5em' />
                  <Text fontSize='lg' fontWeight='semibold'>
                    報告されているミッションはありません
                  </Text>
                </Center>
              </Table.Cell>
            </Table.Row>
          ) : (
            <>
              {items.map((item) => (
                <Table.Row borderColor='yellow.400' borderWidth={4} key={item.rewardPeriodId}>
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
                    <CompletedTag status={item.status}></CompletedTag>
                  </Table.Cell>
                  <Table.Cell>
                    <RewardApproveDialog
                      item={item}
                      setRefreshKey={setRefreshKey}
                    ></RewardApproveDialog>
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
