import { Box, Tabs, VStack, Text, Float, Circle } from '@chakra-ui/react';
import { use, useContext, useEffect, useState } from 'react';
import { RewardTable } from '../RewardTable';
import { RewardFloat } from './RewardFloat';
import {
  RewardService,
  Reward_Type,
  GetUncompletedRewardCountResponse,
} from '@/__generated__/v1/odachin/reward_pb';
import { useClient } from '@/pages/api/ClientProvider';
import { CokiesContext } from '@/pages/api/CokiesContext';
import unauthorizedPage from '@/pages/unauthorized';

export interface RewardPageProps {}

export function RewardPage({}: RewardPageProps) {
  const client = useClient(RewardService);
  const cookies = useContext(CokiesContext);
  const [rewardType, setRewardType] = useState<Reward_Type>(Reward_Type.DAILY);
  const [count, setCount] = useState<GetUncompletedRewardCountResponse>();

  useEffect(() => {
    const fetchData = async () => {
      if (!cookies || !cookies.authorization) {
        unauthorizedPage();
        console.error('No authentication token found');
        return;
      }
      const res = await client.getUncompletedRewardCount(
        {},
        {
          headers: { authorization: cookies?.authorization },
        },
      );
      setCount(res);
    };
    try {
      fetchData();
    } catch (error) {
      console.error('Error fetching reward list:', error);
    }
  }, []);
  console.log('count', count);
  return (
    <>
      <VStack bgColor='white' minH='100vh' textAlign='center' w='100%'>
        <Box mt={12} w='80%'>
          <Tabs.Root
            colorPalette='orange'
            defaultValue={Reward_Type.DAILY.toString()}
            fitted
            onValueChange={(e) => setRewardType(Number(e.value) as Reward_Type)}
            variant='subtle'
          >
            <Tabs.List>
              <Tabs.Trigger borderTopRadius='24px' value={Reward_Type.DAILY.toString()}>
                <RewardFloat count={count?.dailyCount}></RewardFloat>
                <Text fontWeight='semibold' textStyle='2xl'>
                  今日のミッション
                </Text>
              </Tabs.Trigger>
              <Tabs.Trigger borderTopRadius={24} value={Reward_Type.WEEKLY.toString()}>
                <RewardFloat count={count?.weeklyCount}></RewardFloat>
                <Text fontWeight='semibold' textStyle='2xl'>
                  今週のミッション
                </Text>
              </Tabs.Trigger>
              <Tabs.Trigger borderTopRadius={24} value={Reward_Type.MONTHLY.toString()}>
                <RewardFloat count={count?.monthlyCount}></RewardFloat>
                <Text fontWeight='semibold' textStyle='2xl'>
                  今月のミッション
                </Text>
              </Tabs.Trigger>
            </Tabs.List>
          </Tabs.Root>
          <RewardTable rewardType={rewardType}></RewardTable>
        </Box>
      </VStack>
    </>
  );
}
