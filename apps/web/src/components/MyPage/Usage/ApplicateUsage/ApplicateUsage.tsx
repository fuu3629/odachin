import { Box, VStack, Field, Input, Button, Text, HStack, Flex } from '@chakra-ui/react';
import { useState, useEffect } from 'react';
import { Controller } from 'react-hook-form';
import CreatableSelect from 'react-select/creatable';
import { ApplicationStatusTable } from './ApplicationStatusTable ';
import { useApplicateUsageForm } from './lib';
import { UsageService } from '@/__generated__/v1/odachin/usage_pb';
import App from '@/pages/_app';
import { useClient } from '@/pages/api/ClientProvider';

export interface ApplicateUsageProps {}

const createOption = (label: string) => ({
  label,
  value: label,
});

type OptionType = {
  label: string;
  value: string;
};

export function ApplicateUsage({}: ApplicateUsageProps) {
  const {
    register,
    onSubmit,
    formState: { errors },
    control,
  } = useApplicateUsageForm();
  const client = useClient(UsageService);
  const [categoryInfo, setCategoryInfo] = useState<string[]>([]);

  useEffect(() => {
    const fetchData = async () => {
      try {
        const res = await client.getUsageCategories({});
        setCategoryInfo(res.categories);
      } catch (error) {
        alert('カテゴリの取得に失敗しました。');
      }
    };

    fetchData();
  }, []);
  const categoryOptions: OptionType[] = categoryInfo?.map((category) => {
    return createOption(category);
  });

  return (
    <Flex gap={8} m={10} w='100%'>
      <Box bg='white' borderRadius='xl' boxShadow='lg' m='auto' p={8} w='50%'>
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

            <Field.Root invalid={!!errors.category} w='full'>
              <Field.Label fontSize='sm' fontWeight='bold' mb={1}>
                カテゴリを入力
              </Field.Label>
              <Controller
                control={control}
                name='category'
                render={({ field }) => (
                  <CreatableSelect<OptionType, false>
                    {...field}
                    onChange={(value) => field.onChange(value)}
                    onCreateOption={(inputValue) => {
                      const newOption = { label: inputValue, value: inputValue };
                      field.onChange(newOption);
                    }}
                    options={[{ label: 'カテゴリー', options: categoryOptions }]}
                    placeholder='例: 本'
                    styles={{
                      container: (provided) => ({
                        ...provided,
                        width: '100%',
                      }),
                      control: (provided) => ({
                        ...provided,
                        minHeight: '40px',
                      }),
                      menu: (provided) => ({
                        ...provided,
                        zIndex: 9999,
                      }),
                    }}
                    value={field.value}
                  />
                )}
              />
              <Field.ErrorText fontSize='sm'>{errors.category?.message}</Field.ErrorText>
            </Field.Root>

            <Button _hover={{ bg: 'gray.800' }} bg='black' color='white' type='submit' w='full'>
              申請する
            </Button>
          </VStack>
        </form>
      </Box>
      <ApplicationStatusTable />
    </Flex>
  );
}
