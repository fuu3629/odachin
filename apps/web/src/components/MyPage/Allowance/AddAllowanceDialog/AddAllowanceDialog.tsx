import { Dialog, Button, Portal, CloseButton, Link, Text } from '@chakra-ui/react';
import { IoIosAddCircleOutline } from 'react-icons/io';

export interface AddAllowanceDialogProps {}

//TODO これ実装する
export function AddAllowanceDialog({}: AddAllowanceDialogProps) {
  return (
    <>
      <Dialog.Root>
        <Dialog.Trigger asChild>
          <Link onClick={() => {}}>
            <IoIosAddCircleOutline size='1.5em' />
            <Text fontSize='lg' fontWeight='semibold'>
              ユーザーを追加する
            </Text>
          </Link>
        </Dialog.Trigger>
        <Portal>
          <Dialog.Backdrop />
          <Dialog.Positioner>
            <Dialog.Content>
              <Dialog.Header>
                <Dialog.Title>お小遣いを追加する</Dialog.Title>
              </Dialog.Header>
              <Dialog.Body>
                <p>
                  Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed do eiusmod tempor
                  incididunt ut labore et dolore magna aliqua.
                </p>
              </Dialog.Body>
              <Dialog.Footer>
                <Dialog.ActionTrigger asChild>
                  <Button variant='outline'>Cancel</Button>
                </Dialog.ActionTrigger>
                <Button>Save</Button>
              </Dialog.Footer>
              <Dialog.CloseTrigger asChild>
                <CloseButton size='sm' />
              </Dialog.CloseTrigger>
            </Dialog.Content>
          </Dialog.Positioner>
        </Portal>
      </Dialog.Root>
    </>
  );
}
