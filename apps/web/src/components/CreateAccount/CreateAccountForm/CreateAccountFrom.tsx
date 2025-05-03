import {
  Flex,
  Heading,
  Input,
  Button,
  Text,
  Box,
  VStack,
  Fieldset,
  Field,
  HStack,
  RadioGroup,
} from '@chakra-ui/react';
import { Controller } from 'react-hook-form';
import { useCreateAccountForm } from './lib';
export interface CreateAccountFormProps {}

//TODO ROLEの設定する
const roleOptions = [
  { label: '保護者', value: '0' },
  { label: '子供', value: '1' },
];

const fieldItems: { label: string; value: 'email' | 'userId' | 'userName' | 'password' }[] = [
  { label: 'メールアドレス', value: 'email' },
  { label: 'ユーザID', value: 'userId' },
  { label: 'ユーザ名', value: 'userName' },
  { label: 'パスワード', value: 'password' },
];
export function CreateAccountForm({}: CreateAccountFormProps) {
  const {
    register,
    onSubmit,
    control,
    formState: { errors },
  } = useCreateAccountForm();
  return (
    <Flex
      alignItems='center'
      bgSize='cover'
      direction='column'
      h='100vh'
      justifyContent='center'
      w='100vw'
    >
      <Box
        bg='white'
        bgImage="url('/mnt/data/image.png')"
        borderRadius='lg'
        boxShadow='lg'
        maxW='500px'
        p={8}
        w='100%'
      >
        <Flex alignItems='center' bg='white' direction='column' justify='center' p={12}>
          <Box maxW='400px' w='100%'>
            <form onSubmit={onSubmit}>
              <Fieldset.Root maxW='md' size='lg'>
                <Heading mb={6}>新規登録</Heading>
                <Fieldset.Content>
                  <Field.Root invalid={!!errors.role}>
                    <Field.Label fontWeight='semibold'>タイプ</Field.Label>
                    <Controller
                      control={control}
                      name='role'
                      render={({ field }) => (
                        <RadioGroup.Root
                          onValueChange={({ value }) => {
                            field.onChange(value);
                          }}
                          value={field.value}
                        >
                          <HStack gap='6'>
                            {roleOptions.map((item) => (
                              <RadioGroup.Item key={item.value} value={item.value}>
                                <RadioGroup.ItemHiddenInput />
                                <RadioGroup.ItemIndicator />
                                <RadioGroup.ItemText>{item.label}</RadioGroup.ItemText>
                              </RadioGroup.Item>
                            ))}
                          </HStack>
                        </RadioGroup.Root>
                      )}
                    />
                    <Field.ErrorText>{errors.role?.message}</Field.ErrorText>
                  </Field.Root>
                  {fieldItems.map((item) => (
                    <Field.Root invalid={!!errors[item.value]} key={item.value}>
                      <Field.Label fontWeight='semibold'>{item.label}</Field.Label>
                      <Input bg='gray.100' {...register(item.value)} />
                      <Field.ErrorText>{errors[item.value]?.message}</Field.ErrorText>
                    </Field.Root>
                  ))}

                  <VStack mb={2} mt={8}>
                    <Button
                      _hover={{ bg: 'gray.800' }}
                      bg='black'
                      color='white'
                      type='submit'
                      w='full'
                    >
                      新規登録
                    </Button>
                  </VStack>
                </Fieldset.Content>
              </Fieldset.Root>
            </form>
          </Box>
        </Flex>
      </Box>
    </Flex>
  );
}
