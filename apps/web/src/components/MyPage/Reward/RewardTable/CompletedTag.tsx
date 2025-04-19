import { Center, Tag } from '@chakra-ui/react';

export interface CompletedTagProps {
  isCompleted: boolean;
}

export function CompletedTag({ isCompleted }: CompletedTagProps) {
  return (
    <>
      <Center h='100%' w='100%'>
        {isCompleted ? (
          <Tag.Root colorPalette='green' size='sm' variant='surface'>
            <Tag.Label>完了</Tag.Label>
          </Tag.Root>
        ) : (
          <Tag.Root colorPalette='red' size='sm' variant='surface'>
            <Tag.Label>未完</Tag.Label>
          </Tag.Root>
        )}
      </Center>
    </>
  );
}
