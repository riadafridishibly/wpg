import { Button, Group, Modal, Text } from '@mantine/core';
import { CheckPasswordBreach } from '@wailsjs/go/main/PasswordService';
import { group } from 'console';
import React, { useEffect } from 'react';
import { atom, useRecoilState } from 'recoil';
import { modalButtonColor, modalButtonMargin, modalButtonSize } from '../constants/modal.constants';

let color = 'green';
let innerText = '';

export const showPasswordBreachModal = atom<{ opened: boolean; password: string; count: number }>({
    key: 'showPasswordBreachModal',
    default: {
        opened: false,
        password: '',
        count: -1,
    },
    effects: [
        ({ setSelf, onSet }) => {
            onSet(async (newValue) => {
                let count = 0;
                try {
                    count = await CheckPasswordBreach(newValue.password);
                } catch (ex) {
                    console.log(ex);
                } finally {
                    setSelf({ ...newValue, count: count });
                    console.log(count, color);
                }
                color = count > 0 ? 'red' : 'green';
                innerText =
                    count !== -1
                        ? count > 0
                            ? `Password breach found in ${count} instances!`
                            : `No password breach found`
                        : ``;
            });
        },
    ],
});

const PwnCheckModal = () => {
    const [passwordBreachModalState, setPasswordBreachModalState] =
        useRecoilState(showPasswordBreachModal);

    return (
        <>
            <Modal
                centered={true}
                opened={passwordBreachModalState.opened}
                onClose={() =>
                    setPasswordBreachModalState({ opened: false, password: '', count: -1 })
                }
            >
                <Text color={color} align="center">
                    {passwordBreachModalState.count > 0
                        ? `Password breach found in ${passwordBreachModalState.count} instances!`
                        : passwordBreachModalState.count !== -1
                        ? `No password breach found`
                        : ``}
                </Text>
                <Group position="center">
                    <Button
                        size={modalButtonSize}
                        m={modalButtonMargin}
                        bg={modalButtonColor}
                        onClick={() =>
                            setPasswordBreachModalState({ opened: false, password: '', count: -1 })
                        }
                    >
                        OK
                    </Button>
                </Group>
            </Modal>
        </>
    );
};

export default PwnCheckModal;
