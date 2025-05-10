import { register } from 'module';
import {
  Box,
  VStack,
  Field,
  Input,
  Button,
  Center,
  Text,
  HStack,
  Card,
  CardBody,
  List,
  ListItem,
  Stack,
} from '@chakra-ui/react';
import router from 'next/router';
import { useContext, useEffect, useState } from 'react';
import { number } from 'zod';
import { AddFamilyDialog } from '../../Family/FamilyPage/AddFamilyDialog';
import { FamilyAvatar } from '../../Family/FamilyPage/FamilyAvatar';
import { useApplicateUsageForm } from '../ApplicateUsage/lib';
import { FamilyService, GetFamilyInfoResponse } from '@/__generated__/v1/odachin/faimily_pb';
import { GetUsageApplicationResponse, UsageService } from '@/__generated__/v1/odachin/usage_pb';
import { useLoginForm } from '@/components/Login/LoginForm/lib';
import { useColorModeValue } from '@/components/ui/color-mode';
import { PasswordInput } from '@/components/ui/password-input';
import { useClient } from '@/pages/api/ClientProvider';
import { CokiesContext } from '@/pages/api/CokiesContext';
import usage from '@/pages/myPage/usage';

export interface ApproveUsageProps {}

export function ApproveUsage({}: ApproveUsageProps) {
  const client = useClient(FamilyService);
  const usageClient = useClient(UsageService);
  const cookies = useContext(CokiesContext);
  const [familyInfo, setFamilyInfo] = useState<GetFamilyInfoResponse>();
  const [usageApplication, setUsageApplication] = useState<GetUsageApplicationResponse>();

  useEffect(() => {
    const fetchData = async () => {
      if (!cookies || !cookies.authorization) {
        console.error('No authentication token found');
        return;
      }
      try {
        const res = await client.getFamilyInfo(
          {},
          {
            headers: { authorization: cookies.authorization },
          },
        );

        const req = {
          userId:
            res?.familyMembers
              .filter((member) => member.role !== 0)
              .map((member) => member.userId) || [],
        };
        if (req) {
          setUsageApplication(await usageClient.getUsageApplication(req));
        }

        console.log(req);
      } catch (error) {
        alert('使用リクエストの取得に失敗しました。');
      }
    };

    fetchData();
  }, []);

  const apploveUsageApplication = async (usageId: bigint) => {
    await usageClient.approveUsage({ usageId: usageId });
  };
  return (
    <>
      <HStack gap={8} mb={24} w='100%'>
        {usageApplication?.usageApplications.map((usage) => (
          <Center key={usage.usageId} py={6}>
            <Box bg='white' boxShadow='2xl' overflow='hidden' rounded='md' w='330px'>
              <VStack align='center' direction='row' justify='center'>
                <HStack>
                  <Text fontSize='3xl'>{usage.amount}</Text>
                  <Text>pt</Text>
                </HStack>
                <Text fontSize='4xl' fontWeight={800}>
                  {usage.title}
                </Text>
              </VStack>

              <Box bg='gray.100' px={6} py={10}>
                <Text>{usage.description}</Text>

                <Button
                  _focus={{
                    bg: 'orange.500',
                  }}
                  _hover={{
                    bg: 'orange.500',
                  }}
                  bg='orange.400'
                  boxShadow='0 5px 20px 0px'
                  color='white'
                  fontWeight={600}
                  mt={10}
                  onClick={() => {
                    apploveUsageApplication(usage.usageId);
                    window.location.reload();
                  }}
                  rounded='xl'
                  w='full'
                >
                  承認する
                </Button>
              </Box>
            </Box>
          </Center>
        ))}
      </HStack>
    </>
  );
}
