import { Box, Flex, Avatar, Text, Grid, IconButton, VStack, GridItem } from '@chakra-ui/react';
import { useRouter } from 'next/router';
import { useContext, useEffect, useState } from 'react';
import { FaUser, FaLaptop, FaCog, FaTrophy, FaPiggyBank, FaCheck } from 'react-icons/fa';
import { ApplicateUsage } from './ApplicateUsage';
import { ApproveUsage } from './ApproveUsage';
import { AuthService, GetOwnInfoResponse, Role } from '@/__generated__/v1/odachin/auth_pb';
import { RewardService } from '@/__generated__/v1/odachin/reward_pb';
import { useClient } from '@/pages/api/ClientProvider';
import { CokiesContext } from '@/pages/api/CokiesContext';

export interface UsagePageProps {}

//TODO childの色々使い道のチャートとか作りたい
export function UsagePage({}: UsagePageProps) {
  const cookies = useContext(CokiesContext);
  const router = useRouter();
  const client = useClient(AuthService);
  const rewardClient = useClient(RewardService);
  const [userInfo, setuserInfo] = useState<GetOwnInfoResponse | null>(null);
  const [reportedRewardCount, setReportedRewardCount] = useState(0);
  useEffect(() => {
    const fetchData = async () => {
      const req = {};
      try {
        const res = await client.getOwnInfo(req);
        setuserInfo(res);
      } catch (error) {
        router.push('/login');
        alert('Login failed');
      }
    };
    fetchData();
  }, []);

  return userInfo?.role === Role.PARENT ? (
    <ApproveUsage></ApproveUsage>
  ) : (
    <ApplicateUsage></ApplicateUsage>
  );
}
