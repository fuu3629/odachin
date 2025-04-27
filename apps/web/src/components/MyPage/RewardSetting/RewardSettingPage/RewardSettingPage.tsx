import { VStack, Box, Button, HStack, Text, defineStyle } from '@chakra-ui/react';
import { useContext, useEffect, useState } from 'react';
import { FamilyAvatar } from '../../Family/FamilyPage/FamilyAvatar';
import { RewardSettingTable } from './RewardSettingTable';
import { Role } from '@/__generated__/v1/odachin/auth_pb';
import { FamilyService, FamilyUser } from '@/__generated__/v1/odachin/faimily_pb';
import { useClient } from '@/pages/api/ClientProvider';
import { CokiesContext } from '@/pages/api/CokiesContext';

export interface RewardSettingPageProps {}

export const ringCss = defineStyle({
  outlineWidth: '2px',
  outlineColor: 'orange.500',
  outlineOffset: '4px',
  outlineStyle: 'solid',
});

export interface SelectedUser {
  id: string;
  name: string;
}

export function RewardSettingPage({}: RewardSettingPageProps) {
  const familyClient = useClient(FamilyService);
  const cookies = useContext(CokiesContext);
  const [selectedUser, setSelectedUser] = useState<SelectedUser | undefined>(undefined);
  const [children, setChildren] = useState<FamilyUser[]>([]);

  useEffect(() => {
    const fetchData = async () => {
      if (!cookies || !cookies.authorization) {
        console.error('No authentication token found');
        return;
      }
      try {
        const res = await familyClient.getFamilyInfo(
          {},
          {
            headers: { authorization: cookies?.authorization },
          },
        );
        const childresRes = res.familyMembers.filter((member) => member.role === Role.CHILD);
        if (childresRes.length > 0) {
          setChildren(childresRes);
          setSelectedUser({ id: childresRes[0].userId, name: childresRes[0].name });
        }
      } catch (error) {
        console.error('Error fetching reward list:', error);
      }
    };
    fetchData();
  }, []);
  return (
    <>
      <VStack bgColor='white' minH='100vh' w='100%'>
        <Box mt='24px' w='80%'>
          <Text fontWeight='semibold' mb='24px' textStyle='2xl'>
            お子様
          </Text>
          <HStack>
            {children.map((child) => (
              <FamilyAvatar
                avatarUrl={child.avatarImageUrl}
                colorPalette='orange'
                css={selectedUser?.id === child.userId ? ringCss : {}}
                key={child.userId}
                onClick={() => {
                  setSelectedUser({ id: child.userId, name: child.name });
                }}
                userId={child.userId}
                userName={child.name}
              />
            ))}
          </HStack>
        </Box>
        <Box mt={12} w='80%'>
          <RewardSettingTable selectedUser={selectedUser}></RewardSettingTable>
        </Box>
      </VStack>
    </>
  );
}
