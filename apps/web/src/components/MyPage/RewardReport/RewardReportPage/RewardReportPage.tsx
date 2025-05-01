import { VStack, Box } from '@chakra-ui/react';
import { RewardReportedTable } from '../RewardReportedTable';

export interface RewardReportPageProps {}

export function RewardReportPage({}: RewardReportPageProps) {
  return (
    <>
      <VStack h='90vh' w='100vw'>
        <Box mt='3%' w='80%'>
          <RewardReportedTable></RewardReportedTable>
        </Box>
      </VStack>
    </>
  );
}
