import {
  Dialog,
  Button,
  Portal,
  CloseButton,
  Text,
  Input,
  HStack,
  RadioGroup,
  Select,
  createListCollection,
  Icon,
} from '@chakra-ui/react';
import { useState } from 'react';
import { Controller } from 'react-hook-form';
import { CiCircleCheck } from 'react-icons/ci';
import { AllowanceItem } from '../AllowancePage';
import { UpdateAccountFormSchemaType, useUpdateAccountForm } from './lib';
import { Alloance_Type, DayOfWeek } from '@/__generated__/v1/odachin/allowance_pb';

export interface UpdateAllowanceDialogProps {
  allowanceItem: AllowanceItem;
}

const typeOptions = [
  { label: '毎日', value: 'DAILY' },
  { label: '毎週', value: 'WEEKLY' },
  { label: '毎月', value: 'MONTHLY' },
];

const weekOptions = [
  { label: '月曜日', value: 'MONDAY' },
  { label: '火曜日', value: 'TUESDAY' },
  { label: '水曜日', value: 'WEDNESDAY' },
  { label: '木曜日', value: 'THURSDAY' },
  { label: '金曜日', value: 'FRIDAY' },
  { label: '土曜日', value: 'SATURDAY' },
  { label: '日曜日', value: 'SUNDAY' },
];
const frameworks = createListCollection({ items: weekOptions });

const dateOptions = [];
for (let i = 1; i <= 31; i++) {
  dateOptions.push({ label: `${i}日`, value: `${i}` });
}
const dateFrameworks = createListCollection({ items: dateOptions });

export function UpdateAllowanceDialog({ allowanceItem }: UpdateAllowanceDialogProps) {
  const [message, setMessage] = useState<boolean | undefined>(undefined);
  const defaultValues: UpdateAccountFormSchemaType = {
    amount: String(allowanceItem.amount),
    allowanceType: Alloance_Type[allowanceItem.allowanceType] as keyof typeof Alloance_Type,
    dayOfWeek: allowanceItem.dayOfWeek
      ? (DayOfWeek[allowanceItem.dayOfWeek] as keyof typeof DayOfWeek)
      : undefined,
    date: allowanceItem.date,
  };
  const {
    register,
    control,
    onSubmit,
    formState: { errors },
    watch,
  } = useUpdateAccountForm(defaultValues, allowanceItem, setMessage);

  return (
    <>
      <Dialog.Root>
        <Dialog.Trigger asChild>
          <Button size='sm'>
            <Text fontWeight='semibold' textAlign='center'>
              編集する
            </Text>
          </Button>
        </Dialog.Trigger>
        <Portal>
          <Dialog.Backdrop />
          <Dialog.Positioner>
            <Dialog.Content>
              <form onSubmit={onSubmit}>
                <Dialog.Header>
                  <Dialog.Title>お小遣いを更新する</Dialog.Title>
                </Dialog.Header>
                <Dialog.Body>
                  <Text fontSize='sm' fontWeight='bold' mb={1}>
                    お小遣い
                  </Text>
                  <Input bg='gray.100' {...register('amount')} type='number' />
                  {errors.amount && (
                    <Text color='red.500' fontSize='sm'>
                      {errors.amount.message}
                    </Text>
                  )}
                  <Text fontSize='sm' fontWeight='bold' mb={1} mt={4}>
                    タイプ
                  </Text>
                  <Controller
                    control={control}
                    name='allowanceType'
                    render={({ field }) => (
                      <RadioGroup.Root
                        defaultValue={defaultValues.allowanceType}
                        onValueChange={({ value }) => {
                          field.onChange(value);
                        }}
                        value={field.value}
                      >
                        <HStack gap='6'>
                          {typeOptions.map((item) => (
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
                  {errors.allowanceType && (
                    <Text color='red.500' fontSize='sm'>
                      {errors.allowanceType.message}
                    </Text>
                  )}
                  {watch('allowanceType') === 'WEEKLY' && (
                    <>
                      <Controller
                        control={control}
                        name='dayOfWeek'
                        render={({ field }) => (
                          <Select.Root
                            collection={frameworks}
                            defaultValue={field.value ? [field.value] : []}
                            onValueChange={({ value }) => field.onChange(value[0])}
                            size='sm'
                            width='320px'
                          >
                            <Select.HiddenSelect />
                            <Select.Label>
                              <Text fontSize='sm' fontWeight='bold' mb={1} mt={4}>
                                曜日
                              </Text>
                            </Select.Label>
                            <Select.Control>
                              <Select.Trigger _hover={{ cursor: 'pointer' }}>
                                <Select.ValueText placeholder='曜日を選択してください' />
                              </Select.Trigger>
                              <Select.IndicatorGroup>
                                <Select.Indicator />
                              </Select.IndicatorGroup>
                            </Select.Control>
                            <Portal>
                              <Select.Positioner>
                                <Select.Content zIndex={1000000}>
                                  {frameworks.items.map((framework) => (
                                    <Select.Item
                                      _hover={{ cursor: 'pointer' }}
                                      item={framework}
                                      key={framework.value}
                                    >
                                      {framework.label}
                                      <Select.ItemIndicator />
                                    </Select.Item>
                                  ))}
                                </Select.Content>
                              </Select.Positioner>
                            </Portal>
                          </Select.Root>
                        )}
                      />

                      {errors.dayOfWeek && (
                        <Text color='red.500' fontSize='sm'>
                          {errors.dayOfWeek.message}
                        </Text>
                      )}
                    </>
                  )}
                  {watch('allowanceType') === 'MONTHLY' && (
                    <>
                      <Controller
                        control={control}
                        name='date'
                        render={({ field }) => (
                          <Select.Root
                            collection={dateFrameworks}
                            defaultValue={field.value ? [String(field.value)] : []}
                            onValueChange={({ value }) => field.onChange(Number(value[0]))}
                            size='sm'
                            width='320px'
                          >
                            <Select.HiddenSelect />
                            <Select.Label>
                              <Text fontSize='sm' fontWeight='bold' mb={1} mt={4}>
                                日付
                              </Text>
                            </Select.Label>
                            <Select.Control>
                              <Select.Trigger _hover={{ cursor: 'pointer' }}>
                                <Select.ValueText placeholder='日付を選択してください' />
                              </Select.Trigger>
                              <Select.IndicatorGroup>
                                <Select.Indicator />
                              </Select.IndicatorGroup>
                            </Select.Control>
                            <Portal>
                              <Select.Positioner>
                                <Select.Content zIndex={1000000}>
                                  {dateFrameworks.items.map((framework) => (
                                    <Select.Item
                                      _hover={{ cursor: 'pointer' }}
                                      item={framework}
                                      key={framework.value}
                                    >
                                      {framework.label}
                                      <Select.ItemIndicator />
                                    </Select.Item>
                                  ))}
                                </Select.Content>
                              </Select.Positioner>
                            </Portal>
                          </Select.Root>
                        )}
                      />

                      {errors.dayOfWeek && (
                        <Text color='red.500' fontSize='sm'>
                          {errors.dayOfWeek.message}
                        </Text>
                      )}
                    </>
                  )}
                  {message === true && (
                    <HStack>
                      <Icon color='green.500'>
                        <CiCircleCheck />
                      </Icon>
                      <Text color='green.500' fontSize='sm'>
                        お小遣いを更新しました
                      </Text>
                    </HStack>
                  )}
                  {message === false && (
                    <HStack>
                      <Icon color='red.500'>
                        <CiCircleCheck />
                      </Icon>
                      <Text color='red.500' fontSize='sm'>
                        お小遣いの更新に失敗しました
                      </Text>
                    </HStack>
                  )}
                </Dialog.Body>
                <Dialog.Footer>
                  <Dialog.ActionTrigger asChild>
                    <Button variant='outline'>キャンセル</Button>
                  </Dialog.ActionTrigger>
                  <Button type='submit'>更新する</Button>
                </Dialog.Footer>
                <Dialog.CloseTrigger asChild>
                  <CloseButton size='sm' />
                </Dialog.CloseTrigger>
              </form>
            </Dialog.Content>
          </Dialog.Positioner>
        </Portal>
      </Dialog.Root>
    </>
  );
}
