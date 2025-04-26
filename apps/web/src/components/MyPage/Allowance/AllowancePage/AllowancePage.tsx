import { Box, VStack } from '@chakra-ui/react';
import { useRouter } from 'next/router';
import { useContext, useEffect, useState } from 'react';
import { AllowanceItem, AllowanceTable } from './AllowanceTable';
import { Alloance_Type, AllowanceService } from '@/__generated__/v1/odachin/allowance_pb';
import { useClient } from '@/pages/api/ClientProvider';
import { CokiesContext } from '@/pages/api/CokiesContext';

export interface AllowancePageProps {}

const dayOfWeekDict = {
  0: '月曜日',
  1: '火曜日',
  2: '水曜日',
  3: '木曜日',
  4: '金曜日',
  5: '土曜日',
  6: '日曜日',
};

export function AllowancePage({}: AllowancePageProps) {
  const cookies = useContext(CokiesContext);
  const router = useRouter();
  const client = useClient(AllowanceService);
  const [allowanceInfo, setAllowanceInfo] = useState<AllowanceItem[]>();
  useEffect(() => {
    if (!cookies || !cookies.authorization) {
      console.error('No authentication token found');
      return;
    }
    const fetchData = async () => {
      const req = {};
      try {
        const res = await client.getAllowanceByFromUserId(req, {
          headers: { authorization: cookies.authorization },
        });
        const allowanceList: AllowanceItem[] = res.allowances.map((allowance) => {
          let date: any = allowance.date;
          if (allowance.intervalType === Alloance_Type.MONTHLY) {
            date = allowance.date;
          }
          if (allowance.intervalType === Alloance_Type.WEEKLY) {
            date = dayOfWeekDict[allowance.dayOfWeek!];
          }
          if (allowance.intervalType === Alloance_Type.DAILY) {
            date = '******';
          }
          return {
            id: allowance.allowanceId,
            toUserId: allowance.toUserId,
            toUserName: allowance.toUserName,
            amount: allowance.amount,
            avatarImageUrl: allowance.avatarImageUrl,
            allowanceType: allowance.intervalType,
            viewDate: date,
            date: allowance.date,
            dayOfWeek: allowance.dayOfWeek,
          };
        });
        setAllowanceInfo(allowanceList);
      } catch (error) {
        router.push('/login');
        console.error('Error fetching user info:', error);
      }
    };
    fetchData();
  }, []);
  return (
    <>
      <VStack h='90vh' w='100vw'>
        <Box mt='3%' w='60%'>
          <AllowanceTable allowanceList={allowanceInfo}></AllowanceTable>
        </Box>
      </VStack>
    </>
  );
}
