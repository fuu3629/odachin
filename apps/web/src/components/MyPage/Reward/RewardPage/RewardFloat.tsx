import { Float, Circle } from '@chakra-ui/react';

export interface RewardFloatProps {
  count?: number;
}

export function RewardFloat({ count }: RewardFloatProps) {
  return (
    <>
      {count === 0 || count == undefined ? null : (
        <Float offset='6px' placement='top-end'>
          <Circle bg='orange' color='white' size='5'>
            {count}
          </Circle>
        </Float>
      )}
    </>
  );
}
