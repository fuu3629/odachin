import {
  Dialog,
  Button,
  Portal,
  VStack,
  DataList,
  CloseButton,
  Text,
  Link,
  Field,
  Fieldset,
  Input,
  Stack,
} from '@chakra-ui/react';
import { Dispatch, SetStateAction, useContext, useState } from 'react';
import { IoIosAddCircleOutline } from 'react-icons/io';
import { SelectedUser } from './RewardSettingPage';
import { useAddRewardForm } from './lib';
import { Reward_Type } from '@/__generated__/v1/odachin/reward_pb';

export interface RewardAddDialogProps {
  rewardType: Reward_Type;
  toUser?: SelectedUser;
  setRefreshKey: Dispatch<SetStateAction<number>>;
}

const RewardTypeDict = {
  [Reward_Type.WEEKLY]: '毎週',
  [Reward_Type.MONTHLY]: '毎月',
  [Reward_Type.DAILY]: '毎日',
};

export function RewardAddDialog({ rewardType, toUser, setRefreshKey }: RewardAddDialogProps) {
  const [open, setOpen] = useState(false);
  const {
    register,
    onSubmit,
    formState: { errors },
  } = useAddRewardForm(setRefreshKey, setOpen, toUser?.id, rewardType);

  return (
    <>
      <Dialog.Root onOpenChange={(e) => setOpen(e.open)} open={open}>
        <Dialog.Trigger asChild>
          <Link>
            <IoIosAddCircleOutline size='1.5em' />
            <Text fontSize='lg' fontWeight='semibold'>
              <Text fontSize='lg' fontWeight='semibold'>
                ミッションの追加
              </Text>
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
                    <Dialog.Title>ミッションを作成する</Dialog.Title>
                  </Dialog.Header>
                  <Dialog.Body>
                    <Fieldset.Content>
                      <Field.Root>
                        <Field.Label>
                          <Text fontWeight='semibold'>名前</Text>
                        </Field.Label>
                        <Text>{toUser?.name}</Text>
                      </Field.Root>

                      <Field.Root>
                        <Field.Label fontWeight='semibold'>タイプ</Field.Label>
                        <Text>{RewardTypeDict[rewardType]}</Text>
                      </Field.Root>

                      <Field.Root invalid={!!errors.title}>
                        <Field.Label fontWeight='semibold'>タイトル</Field.Label>
                        <Input placeholder='タイトル' {...register('title')} />
                        <Field.ErrorText>{errors.title?.message}</Field.ErrorText>
                      </Field.Root>

                      <Field.Root invalid={!!errors.description}>
                        <Field.Label fontWeight='semibold'>説明</Field.Label>
                        <Input placeholder='説明' {...register('description')} />
                        <Field.ErrorText>{errors.description?.message}</Field.ErrorText>
                      </Field.Root>

                      <Field.Root invalid={!!errors.amount}>
                        <Field.Label fontWeight='semibold'>金額</Field.Label>
                        <Input placeholder='100' {...register('amount')} />
                        <Field.ErrorText>{errors.amount?.message}</Field.ErrorText>
                      </Field.Root>
                    </Fieldset.Content>
                  </Dialog.Body>
                  <Dialog.Footer>
                    <Dialog.ActionTrigger asChild>
                      <Button variant='outline'>Cancel</Button>
                    </Dialog.ActionTrigger>
                    <Button type='submit'>作成する</Button>
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
