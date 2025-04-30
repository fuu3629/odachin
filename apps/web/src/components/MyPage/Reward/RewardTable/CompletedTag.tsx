import { Center, Tag } from '@chakra-ui/react';

export interface CompletedTagProps {
  status: string;
}

export function CompletedTag({ status }: CompletedTagProps) {
  function getRoot(status: string) {
    switch (status) {
      case 'COMPLETED':
        return (
          <Tag.Root colorPalette='green' size='sm' variant='surface'>
            <Tag.Label>完了</Tag.Label>
          </Tag.Root>
        );
      case 'IN_PROGRESS':
        return (
          <Tag.Root colorPalette='red' size='sm' variant='surface'>
            <Tag.Label>未完</Tag.Label>
          </Tag.Root>
        );
      case 'REPORTED':
        return (
          <Tag.Root colorPalette='blue' size='sm' variant='surface'>
            <Tag.Label>申請中</Tag.Label>
          </Tag.Root>
        );
      case 'REJECTED':
        return (
          <Tag.Root colorPalette='red' size='sm' variant='surface'>
            <Tag.Label>却下</Tag.Label>
          </Tag.Root>
        );
    }
  }
  return (
    <>
      <Center h='100%' w='100%'>
        {getRoot(status)}
      </Center>
    </>
  );
}
