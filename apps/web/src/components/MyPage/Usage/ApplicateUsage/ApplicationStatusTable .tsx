import { Box, HStack, Table, Text } from '@chakra-ui/react';
import { OptionType } from 'dayjs';
import { useState, useEffect } from 'react';
import { FaArrowAltCircleUp, FaCheckCircle, FaExclamationCircle } from 'react-icons/fa';
import { ApplicateUsageProps } from './ApplicateUsage';
import { useApplicateUsageForm } from './lib';
import { AuthService } from '@/__generated__/v1/odachin/auth_pb';
import {
  GetUsageApplicationResponse,
  UsageApplication,
  UsageService,
} from '@/__generated__/v1/odachin/usage_pb';
import { useClient } from '@/pages/api/ClientProvider';

export function ApplicationStatusTable() {
  const client = useClient(UsageService);
  const authClient = useClient(AuthService);
  const [usageApplication, setUsageApplication] = useState<UsageApplication[]>();
  const StatusCell = ({ status }: { status: string }) => {
    switch (status) {
      case 'APPLICATED':
        return (
          <HStack justify='end'>
            <FaArrowAltCircleUp color='orange' />
            <Text>申請中</Text>
          </HStack>
        );
      case 'APPROVED':
        return (
          <HStack justify='end'>
            <FaCheckCircle color='green' />
            <Text>申請受理</Text>
          </HStack>
        );
      case 'REJECTED':
        return (
          <HStack justify='end'>
            <FaExclamationCircle color='red' />
            <Text>申請却下</Text>
          </HStack>
        );
    }
  };
  function timestampToDateString(timestamp: { seconds: bigint; nanos?: number }): string {
    const date = new Date(
      Number(timestamp.seconds) * 1000 + Math.floor((timestamp.nanos || 0) / 1e6),
    );
    const year = date.getFullYear();
    const month = String(date.getMonth() + 1).padStart(2, '0');
    const day = String(date.getDate()).padStart(2, '0');
    return `${year}/${month}/${day}`;
  }

  useEffect(() => {
    const fetchData = async () => {
      try {
        const authResponse = await authClient.getOwnInfo({});
        const res = await client.getUsageApplication({ userId: [authResponse.userId] });

        setUsageApplication(res.usageApplications);
      } catch (error) {
        alert('申請情報の取得に失敗しました。');
      }
    };

    fetchData();
  }, []);

  return (
    <Table.ScrollArea borderWidth='1px' height='800px' rounded='md' w='900px'>
      <Table.Root size='lg' stickyHeader>
        <Table.Header>
          <Table.Row bg='bg.subtle'>
            <Table.ColumnHeader>日付</Table.ColumnHeader>
            <Table.ColumnHeader>タイトル</Table.ColumnHeader>
            <Table.ColumnHeader>説明</Table.ColumnHeader>
            <Table.ColumnHeader>ポイント</Table.ColumnHeader>
            <Table.ColumnHeader textAlign='end'>申請状況</Table.ColumnHeader>
          </Table.Row>
        </Table.Header>

        <Table.Body>
          {usageApplication?.map((usage) => (
            <Table.Row key={usage.usageId}>
              <Table.Cell>
                {usage.createdAt ? timestampToDateString(usage.createdAt) : '不明'}
              </Table.Cell>
              <Table.Cell maxW='200px' whiteSpace='normal' wordBreak='break-word'>
                {usage.title}
              </Table.Cell>
              <Table.Cell maxW='200px' whiteSpace='normal' wordBreak='break-word'>
                {usage.description}
              </Table.Cell>
              <Table.Cell maxW='200px' whiteSpace='normal' wordBreak='break-word'>
                {usage.amount}
              </Table.Cell>
              <Table.Cell>
                <StatusCell status={usage.status} />
              </Table.Cell>
            </Table.Row>
          ))}
        </Table.Body>
      </Table.Root>
    </Table.ScrollArea>
  );
}
