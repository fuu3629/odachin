import { Chart, useChart } from '@chakra-ui/charts';
import {
  Box,
  Flex,
  Avatar,
  Text,
  Grid,
  IconButton,
  VStack,
  GridItem,
  HStack,
  Spacer,
} from '@chakra-ui/react';
import { useRouter } from 'next/router';
import { useEffect, useState } from 'react';
import { FaUser, FaLaptop, FaCog, FaTrophy, FaPiggyBank, FaCheck } from 'react-icons/fa';
import { IoChatbubble } from 'react-icons/io5';
import { PieChart, Pie, Cell } from 'recharts';
import { MyPageFloat } from './MyPageFloat';
import { AuthService, GetOwnInfoResponse, Role } from '@/__generated__/v1/odachin/auth_pb';
import { RewardService } from '@/__generated__/v1/odachin/reward_pb';
import { UsageService } from '@/__generated__/v1/odachin/usage_pb';
import { useClient } from '@/pages/api/ClientProvider';

export interface MyPageProps {}

export interface ChartData {
  name: string;
  value: number;
  color: string;
}

const COLORS = ['orange.500', 'green.500', 'blue.500', 'red.500', 'purple.500', 'pink.500'];

//TODO childの色々使い道のチャートとか作りたい
export function MyPage({}: MyPageProps) {
  const router = useRouter();
  const client = useClient(AuthService);
  const rewardClient = useClient(RewardService);
  const usageClient = useClient(UsageService);
  const [userInfo, setuserInfo] = useState<GetOwnInfoResponse | null>(null);
  const [reportedRewardCount, setReportedRewardCount] = useState(0);
  const [summary, setSummary] = useState<ChartData[]>([]);

  useEffect(() => {
    const fetchData = async () => {
      const req = {};
      try {
        const res = await client.getOwnInfo(req);
        setuserInfo(res);
        const res2 = await rewardClient.getReportedRewardList({});
        setReportedRewardCount(res2.rewardList.length);
        const res3 = await usageClient.getUsageSummary({});
        const data: ChartData[] = res3.usageSummaries.map((summary, index) => ({
          name: summary.category,
          value: summary.amount,
          color: COLORS[index % COLORS.length],
        }));
        setSummary(data);
      } catch (error) {
        router.push('/login');
        alert('Login failed');
      }
    };
    fetchData();
  }, []);

  const chart = useChart({
    data: summary,
  });

  const menuItems =
    userInfo?.role === Role.PARENT
      ? [
          {
            icon: FaPiggyBank,
            label: 'お小遣いを管理',
            onCLick: () => {
              router.push('myPage/allowance');
            },
          },
          {
            icon: FaCheck,
            label: 'ミッションの報告',
            onCLick: () => {
              router.push('myPage/rewardReport');
            },
            count: reportedRewardCount,
          },
          {
            icon: FaTrophy,
            label: 'ミッションを管理',
            onCLick: () => {
              router.push('myPage/rewardSetting');
            },
          },
          {
            icon: FaUser,
            label: '家族情報',
            onCLick: () => {
              router.push('myPage/family');
            },
          },
          {
            icon: IoChatbubble,
            label: '使用リクエスト',
            onCLick: () => {
              router.push('myPage/usage');
            },
          },
          {
            icon: FaCog,
            label: '設定',
            onCLick: () => {
              router.push('setting/account');
            },
          },
        ]
      : [
          {
            icon: FaTrophy,
            label: 'ミッション',
            onCLick: () => {
              router.push('myPage/reward');
            },
          },
          {
            icon: FaTrophy,
            label: '使う',
            onCLick: () => {
              router.push('myPage/usage');
            },
          },
          {
            icon: FaLaptop,
            label: '取引履歴',
            onCLick: () => {
              router.push('myPage/transaction');
            },
          },
          {
            icon: FaCog,
            label: '設定',
            onCLick: () => {
              router.push('setting/account');
            },
          },
        ];

  return (
    <Box bg='white' minH='100vh' w='100%'>
      <Box bg='yellow.400' h='100px' />

      <Flex align='center' direction='column' mt='-50px'>
        <VStack align='flex-start' maxW='800px' w='full'>
          <Avatar.Root size='2xl' top='20px'>
            <Avatar.Fallback name='Segun Adebayo' />
            <Avatar.Image src={userInfo?.avaterImageUrl} />
          </Avatar.Root>
          <Text fontSize='2xl' fontWeight='bold' mt={2}>
            {userInfo?.name}
          </Text>
        </VStack>
        <HStack w='700px'>
          <Spacer></Spacer>
          <Chart.Root boxSize='200px' chart={chart} mx='auto'>
            <PieChart>
              <Pie
                data={chart.data}
                dataKey={chart.key('value')}
                innerRadius={0}
                isAnimationActive={false}
                label={({ name, index }) => {
                  const { value } = chart.data[index ?? -1];
                  const percent = value / chart.getTotal('value');
                  return `${name}: ${(percent * 100).toFixed(1)}%`;
                }}
                labelLine={false}
                outerRadius={100}
              >
                {chart.data.map((item) => {
                  return (
                    <Cell fill={chart.color(item.color)} fontWeight='semibold' key={item.name} />
                  );
                })}
              </Pie>
            </PieChart>
          </Chart.Root>
        </HStack>
        <Box maxW='800px' mt={10} px={4} w='full'>
          <Text fontSize='lg' fontWeight='bold' mb={4}>
            メニュー
          </Text>
          <Grid gap={6} px={4} templateColumns={{ base: 'repeat(2, 1fr)', md: 'repeat(4, 1fr)' }}>
            {menuItems.map((item, index) => (
              <GridItem key={index} onClick={item.onCLick} textAlign='center'>
                <Flex
                  _hover={{ bg: 'gray.100', cursor: 'pointer' }}
                  align='center'
                  bg='gray.50'
                  direction='column'
                  justify='center'
                  p={4}
                  rounded='xl'
                  shadow='md'
                >
                  <IconButton
                    aria-label={item.label}
                    colorScheme='blackAlpha'
                    mb={2}
                    rounded='full'
                    size='lg'
                  >
                    <MyPageFloat count={item.count}></MyPageFloat>
                    <item.icon size={24} />
                  </IconButton>
                  <Text>{item.label}</Text>
                </Flex>
              </GridItem>
            ))}
          </Grid>
        </Box>
      </Flex>
    </Box>
  );
}
