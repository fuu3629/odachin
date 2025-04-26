import { Table, Center, Button, Text, Link, Avatar } from '@chakra-ui/react';
import { CiWarning } from 'react-icons/ci';
import { AddAllowanceDialog } from '../AddAllowanceDialog';
import { UpdateAllowanceDialog } from '../UpdateAllowanceDialog';
import {
  Alloance_Type,
  AllowanceService,
  DayOfWeek,
} from '@/__generated__/v1/odachin/allowance_pb';
import { useClient } from '@/pages/api/ClientProvider';

export interface AllowanceItem {
  id: bigint;
  toUserId: string;
  avatarImageUrl: string | undefined;
  toUserName: string;
  amount: number;
  allowanceType: Alloance_Type;
  viewDate: any;
  dayOfWeek?: DayOfWeek;
  date?: number;
}

export interface AllowanceTableProps {
  allowanceList?: AllowanceItem[];
}

const AllowanceTypeDict = {
  [Alloance_Type.WEEKLY]: '毎週',
  [Alloance_Type.MONTHLY]: '毎月',
  [Alloance_Type.DAILY]: '毎日',
};

export function AllowanceTable({ allowanceList }: AllowanceTableProps) {
  return (
    <>
      <Table.Root borderColor='yellow.400' borderTopRadius={24} showColumnBorder variant='outline'>
        <Table.Header bgColor='yellow.400' borderColor='yellow.400' borderXWidth={4}>
          <Table.Row>
            <Table.ColumnHeader textAlign='center' w='25%'>
              <Text fontSize='lg' fontWeight='semibold'>
                ユーザー名
              </Text>
            </Table.ColumnHeader>
            <Table.ColumnHeader textAlign='center'>
              <Text fontSize='lg' fontWeight='semibold'>
                おこづかいの金額
              </Text>
            </Table.ColumnHeader>
            <Table.ColumnHeader textAlign='center'>
              <Text fontSize='lg' fontWeight='semibold'>
                タイプ
              </Text>
            </Table.ColumnHeader>
            <Table.ColumnHeader textAlign='center'>
              <Text fontSize='lg' fontWeight='semibold'>
                あげる日
              </Text>
            </Table.ColumnHeader>
            <Table.ColumnHeader textAlign='center'>
              <Text fontSize='lg' fontWeight='semibold'>
                編集
              </Text>
            </Table.ColumnHeader>
          </Table.Row>
        </Table.Header>
        <Table.Body bgColor='white'>
          {allowanceList?.length === 0 ? (
            <Table.Row borderColor='yellow.400' borderWidth={4}>
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
              {allowanceList?.map((item) => (
                <Table.Row borderColor='yellow.400' borderWidth={4} key={item.id}>
                  <Table.Cell textAlign='center'>
                    <Avatar.Root>
                      <Avatar.Image src={item.avatarImageUrl}></Avatar.Image>
                    </Avatar.Root>
                    <Text fontWeight='semibold' lineClamp='1' mt={4}>
                      {item.toUserName}
                    </Text>
                  </Table.Cell>
                  <Table.Cell>
                    <Text fontWeight='semibold' lineClamp='1' textAlign='center'>
                      {item.amount}
                    </Text>
                  </Table.Cell>
                  <Table.Cell>
                    <Text fontWeight='semibold' textAlign='center'>
                      {AllowanceTypeDict[item.allowanceType]}
                    </Text>
                  </Table.Cell>
                  <Table.Cell>
                    <Text fontWeight='semibold' textAlign='center'>
                      {item.viewDate}
                    </Text>
                  </Table.Cell>
                  <Table.Cell>
                    <Center>
                      <UpdateAllowanceDialog allowanceItem={item}></UpdateAllowanceDialog>
                    </Center>
                  </Table.Cell>
                </Table.Row>
              ))}
            </>
          )}
          <Table.Row>
            <Table.Cell borderColor='yellow.400' borderWidth={4} colSpan={5} textAlign='center'>
              <Center>
                <AddAllowanceDialog></AddAllowanceDialog>
              </Center>
            </Table.Cell>
          </Table.Row>
        </Table.Body>
      </Table.Root>
    </>
  );
}
