import { register } from 'module';
import { Box, VStack, Field, Input, Button, Center } from '@chakra-ui/react';
import { useCreateGroupForm } from './lib';
import { useLoginForm } from '@/components/Login/LoginForm/lib';
import { PasswordInput } from '@/components/ui/password-input';

export interface NotInFamilyPageProps {}

//TODO 家族グループのメンバーでない場合のページ
// 親なら家族グループを作成するフォームを作成する
// 親が子供のアカウント作成するのもあり？

export function NotInFamilyPage({}: NotInFamilyPageProps) {
  const {
    register,
    onSubmit,
    formState: { errors },
  } = useCreateGroupForm();
  return (
    <>
      <div>NotInFamilyPage</div>
      <Box maxW='400px' w='100%'>
        <form onSubmit={onSubmit} style={{ width: '100%' }}>
          <VStack gap={4}>
            <Field.Root invalid={!!errors.familyName}>
              <Field.Label fontSize='sm' fontWeight='bold' mb={1}>
                ファミリー名を入力
              </Field.Label>
              <Input bg='gray.100' placeholder='example' {...register('familyName')} />
              <Field.ErrorText fontSize='sm'>{errors.familyName?.message}</Field.ErrorText>
            </Field.Root>

            <Button _hover={{ bg: 'gray.800' }} bg='black' color='white' type='submit' w='full'>
              作成する
            </Button>
          </VStack>
        </form>
      </Box>
    </>
  );
}
