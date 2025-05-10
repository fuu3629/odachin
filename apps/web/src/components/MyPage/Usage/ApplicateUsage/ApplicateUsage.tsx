import { register } from 'module';
import { Box, VStack, Field, Input, Button, Center } from '@chakra-ui/react';
import { number } from 'zod';
import { useApplicateUsageForm } from './lib';
import { useLoginForm } from '@/components/Login/LoginForm/lib';
import { PasswordInput } from '@/components/ui/password-input';

export interface ApplicateUsageProps {}

export function ApplicateUsage({}: ApplicateUsageProps) {
  const {
    register,
    onSubmit,
    formState: { errors },
  } = useApplicateUsageForm();
  return (
    <>
      <Box maxW='400px' w='100%'>
        <form onSubmit={onSubmit} style={{ width: '100%' }}>
          <VStack gap={4}>
            <Field.Root invalid={!!errors.title}>
              <Field.Label fontSize='sm' fontWeight='bold' mb={1}>
                タイトルを入力
              </Field.Label>
              <Input bg='gray.100' placeholder='example' {...register('title')} />
              <Field.ErrorText fontSize='sm'>{errors.title?.message}</Field.ErrorText>
            </Field.Root>

            <Field.Root invalid={!!errors.amount}>
              <Field.Label fontSize='sm' fontWeight='bold' mb={1}>
                使用ポイント
              </Field.Label>
              <Input
                bg='gray.100'
                placeholder='example'
                type='number'
                {...register('amount', { valueAsNumber: true })}
              />
              <Field.ErrorText fontSize='sm'>{errors.amount?.message}</Field.ErrorText>
            </Field.Root>

            <Field.Root invalid={!!errors.description}>
              <Field.Label fontSize='sm' fontWeight='bold' mb={1}>
                説明を入力
              </Field.Label>
              <Input bg='gray.100' placeholder='example' {...register('description')} />
              <Field.ErrorText fontSize='sm'>{errors.description?.message}</Field.ErrorText>
            </Field.Root>

            <Field.Root invalid={!!errors.category}>
              <Field.Label fontSize='sm' fontWeight='bold' mb={1}>
                カテゴリを入力
              </Field.Label>
              <Input bg='gray.100' placeholder='example' {...register('category')} />
              <Field.ErrorText fontSize='sm'>{errors.category?.message}</Field.ErrorText>
            </Field.Root>

            <Button _hover={{ bg: 'gray.800' }} bg='black' color='white' type='submit' w='full'>
              申請する
            </Button>
          </VStack>
        </form>
      </Box>
    </>
  );
}
