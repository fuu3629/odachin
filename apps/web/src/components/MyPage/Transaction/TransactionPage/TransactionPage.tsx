import { Box, HStack } from '@chakra-ui/react';
import { EventSourceInput } from '@fullcalendar/core/index.js';
import dayGridPlugin from '@fullcalendar/daygrid';
import FullCalendar from '@fullcalendar/react'; // どのpluginよりも先にimportする必要があります
import { useEffect, useState } from 'react';
import { TransactionService } from '@/__generated__/v1/odachin/transaction_pb';
import { useClient } from '@/pages/api/ClientProvider';

export interface TransactionPageProps {}

export function TransactionPage({}: TransactionPageProps) {
  const transactionClient = useClient(TransactionService);
  const [year, setYear] = useState<number | undefined>(undefined);
  const [month, setMonth] = useState<number | undefined>(undefined);
  const [events, setEvents] = useState<EventSourceInput>([]);
  useEffect(() => {
    const fetchData = async () => {
      if (year === undefined || month === undefined) {
        return;
      }
      const req = {
        startYear: year,
        startMonth: month,
        endYear: year,
        endMonth: month,
      };
      try {
        const res = await transactionClient.getTransactionList(req);
        const events = res.transactionList.map((transaction) => ({
          id: transaction.transactionId.toString(),
          title: 'お小遣い',
          date: new Date(Number(transaction.createdAt?.seconds!) * 1000),
          start: new Date(Number(transaction.createdAt?.seconds!) * 1000),
          end: new Date(Number(transaction.createdAt?.seconds!) * 1000),
          backgroundColor: '#dbbc29',
          borderColor: '#dbbc29',
        }));
        setEvents(events);
      } catch (error) {
        console.error('Error fetching transaction list:', error);
      }
    };
    fetchData();
  }, [year, month]);
  return (
    <>
      <HStack height='calc(100vh - 64px)' w='100%'>
        <Box bg='white' h='100%' w='100%'>
          <FullCalendar
            buttonText={{
              today: '今日に戻る',
              month: '月',
              week: '週',
              day: '日',
            }}
            datesSet={(dateInfo) => {
              setYear(dateInfo.view.currentStart.getFullYear());
              setMonth(dateInfo.view.currentStart.getMonth() + 1);
            }}
            displayEventTime={false}
            eventDisplay='block'
            events={events}
            headerToolbar={{
              center: 'title',
              left: undefined,
            }}
            height='100%'
            initialView='dayGridMonth'
            locale='ja'
            plugins={[dayGridPlugin]}
          />
        </Box>
      </HStack>
    </>
  );
}
