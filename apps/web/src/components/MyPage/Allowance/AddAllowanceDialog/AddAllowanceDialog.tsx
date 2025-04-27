import {
  Dialog,
  Button,
  Portal,
  CloseButton,
  Link,
  Text,
  Grid,
  GridItem,
  Fieldset,
  createListCollection,
  HStack,
  RadioGroup,
  Field,
  Input,
  Select,
} from '@chakra-ui/react';
import { Dispatch, SetStateAction, useContext, useEffect, useState } from 'react';
import { Controller } from 'react-hook-form';
import { IoIosAddCircleOutline } from 'react-icons/io';
import { FamilyAvatar } from '../../Family/FamilyPage/FamilyAvatar';
import { ringCss, SelectedUser } from '../../RewardSetting/RewardSettingPage/RewardSettingPage';
import { useAddAllowanceForm } from './lib';
import { Role } from '@/__generated__/v1/odachin/auth_pb';
import { FamilyService, FamilyUser } from '@/__generated__/v1/odachin/faimily_pb';
import { useClient } from '@/pages/api/ClientProvider';
import { CokiesContext } from '@/pages/api/CokiesContext';

export interface AddAllowanceDialogProps {
  setRefreshKey: Dispatch<SetStateAction<number>>;
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

export function AddAllowanceDialog({ setRefreshKey }: AddAllowanceDialogProps) {
  const familyClient = useClient(FamilyService);
  const cookies = useContext(CokiesContext);
  const [selectedUser, setSelectedUser] = useState<SelectedUser | undefined>(undefined);
  const [children, setChildren] = useState<FamilyUser[]>([]);
  const [open, setOpen] = useState(false);

  const {
    register,
    control,
    onSubmit,
    formState: { errors },
    watch,
  } = useAddAllowanceForm(setOpen, setRefreshKey, selectedUser);
  useEffect(() => {
    const fetchData = async () => {
      if (!cookies || !cookies.authorization) {
        console.error('No authentication token found');
        return;
      }
      try {
        const res = await familyClient.getFamilyInfo({});
        const childresRes = res.familyMembers.filter((member) => member.role === Role.CHILD);
        //TODO 0の時どうするか
        if (childresRes.length > 0) {
          setChildren(childresRes);
          setSelectedUser({ id: childresRes[0].userId, name: childresRes[0].name });
        }
      } catch (error) {
        console.error('Error fetching reward list:', error);
      }
    };
    fetchData();
  }, []);
  return (
    <>
      <Dialog.Root onOpenChange={(e) => setOpen(e.open)} open={open}>
        <Dialog.Trigger asChild>
          <Link onClick={() => {}}>
            <IoIosAddCircleOutline size='1.5em' />
            <Text fontSize='lg' fontWeight='semibold'>
              お小遣いを追加する
            </Text>
          </Link>
        </Dialog.Trigger>
        <Portal>
          <Dialog.Backdrop />
          <Dialog.Positioner>
            <Dialog.Content>
              <form onSubmit={onSubmit}>
                <Fieldset.Root maxW='md' size='lg'>
                  <Dialog.Header>
                    <Dialog.Title>お小遣いを追加する</Dialog.Title>
                  </Dialog.Header>
                  <Dialog.Body>
                    <Fieldset.Content>
                      <Field.Root invalid={!!errors.intervalType}>
                        <Field.Label fontWeight='semibold'>お子様</Field.Label>
                        <Grid
                          gap={6}
                          px={4}
                          templateColumns={{ base: 'repeat(2, 1fr)', md: 'repeat(3, 1fr)' }}
                        >
                          {children.map((child) => (
                            <GridItem
                              key={child.userId}
                              onClick={() =>
                                setSelectedUser({ id: child.userId, name: child.name })
                              }
                            >
                              <FamilyAvatar
                                avatarUrl={child.avatarImageUrl}
                                colorPalette='orange'
                                css={selectedUser?.id === child.userId ? ringCss : {}}
                                key={child.userId}
                                onClick={() => {
                                  setSelectedUser({ id: child.userId, name: child.name });
                                }}
                                userId={child.userId}
                                userName={child.name}
                              />
                            </GridItem>
                          ))}
                        </Grid>
                      </Field.Root>

                      <Field.Root invalid={!!errors.intervalType}>
                        <Field.Label fontWeight='semibold'>タイプ</Field.Label>
                        <Controller
                          control={control}
                          name='intervalType'
                          render={({ field }) => (
                            <RadioGroup.Root
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
                        <Field.ErrorText>{errors.intervalType?.message}</Field.ErrorText>
                      </Field.Root>

                      <Field.Root invalid={!!errors.amount}>
                        <Field.Label fontWeight='semibold'>金額</Field.Label>
                        <Input placeholder='100' {...register('amount')} />
                        <Field.ErrorText>{errors.amount?.message}</Field.ErrorText>
                      </Field.Root>

                      {watch('intervalType') === 'WEEKLY' && (
                        <>
                          <Field.Root invalid={!!errors.dayOfWeek}>
                            <Controller
                              control={control}
                              name='dayOfWeek'
                              render={({ field }) => (
                                <Select.Root
                                  collection={frameworks}
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
                            <Field.ErrorText>{errors.dayOfWeek?.message}</Field.ErrorText>
                          </Field.Root>
                        </>
                      )}
                      {watch('intervalType') === 'MONTHLY' && (
                        <>
                          <Field.Root invalid={!!errors.date}>
                            <Controller
                              control={control}
                              name='date'
                              render={({ field }) => (
                                <Select.Root
                                  collection={dateFrameworks}
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
                            <Field.ErrorText>{errors.date?.message}</Field.ErrorText>
                          </Field.Root>
                        </>
                      )}
                    </Fieldset.Content>
                  </Dialog.Body>
                  <Dialog.Footer>
                    <Dialog.ActionTrigger asChild>
                      <Button variant='outline'>キャンセル</Button>
                    </Dialog.ActionTrigger>
                    <Button type='submit'>登録</Button>
                  </Dialog.Footer>
                  <Dialog.CloseTrigger asChild>
                    <CloseButton size='sm' />
                  </Dialog.CloseTrigger>
                </Fieldset.Root>
              </form>
            </Dialog.Content>
          </Dialog.Positioner>
        </Portal>
      </Dialog.Root>
    </>
  );
}
