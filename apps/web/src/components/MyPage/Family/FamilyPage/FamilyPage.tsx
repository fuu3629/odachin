import { Text, HStack, VStack, Box } from '@chakra-ui/react';
import { useRouter } from 'next/router';
import { useContext, useEffect, useState } from 'react';
import { AddFamilyDialog } from './AddFamilyDialog';
import { FamilyAvatar } from './FamilyAvatar';
import {
  FamilyService,
  GetFamilyInfoResponse,
  GetInvitationListResponse,
} from '@/__generated__/v1/odachin/faimily_pb';
import { useClient } from '@/pages/api/ClientProvider';
import { CokiesContext } from '@/pages/api/CokiesContext';

export interface FamilyPageProps {}

export function FamilyPage({}: FamilyPageProps) {
  const cookies = useContext(CokiesContext);
  const router = useRouter();
  const client = useClient(FamilyService);
  const [familyInfo, setuserInfo] = useState<GetFamilyInfoResponse>();
  const [invitingUser, setInvitingUser] = useState<GetInvitationListResponse>();
  useEffect(() => {
    if (!cookies || !cookies.authorization) {
      console.error('No authentication token found');
      return;
    }
    const fetchData = async () => {
      const req = {};
      try {
        const res = await client.getFamilyInfo(req, {
          headers: { authorization: cookies.authorization },
        });
        setuserInfo(res);
        const res2 = await client.getInvitationList(req, {
          headers: { authorization: cookies.authorization },
        });
        setInvitingUser(res2);
      } catch (error) {
        router.push('/myPage/createGroup');
        console.error('Error fetching user info:', error);
      }
    };
    fetchData();
  }, []);
  return (
    <>
      <VStack h='100%' w='100%'>
        <VStack py='5%' w='60%'>
          <Box mb='12px' textAlign='left' w='100%'>
            <Text fontWeight='semibold' textStyle='2xl'>
              {familyInfo?.familyName}
            </Text>
          </Box>
          <HStack mb={24} w='100%'>
            {familyInfo?.familyMembers.map((member) => (
              <FamilyAvatar
                avatarUrl={member.avatarImageUrl}
                key={member.userId}
                onClick={() => {
                  router.push(`/myPage/family/${member.userId}`);
                }}
                userId={member.userId}
                userName={member.name}
              />
            ))}
            <AddFamilyDialog></AddFamilyDialog>
          </HStack>
          <Box mb='12px' textAlign='left' w='100%'>
            <Text fontWeight='semibold' textStyle='2xl'>
              招待中のユーザー
            </Text>
          </Box>
          <HStack mb={24} w='100%'>
            {invitingUser?.invitationMembers.map((member) => (
              <FamilyAvatar
                avatarUrl={member.avatarImageUrl}
                key={member.userId}
                userId={member.userId}
                userName={member.name}
              />
            ))}
          </HStack>
        </VStack>
      </VStack>
    </>
  );
}
