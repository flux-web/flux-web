import { RootState } from '~/store/types/RootState';


export const root: RootState = {
    readOnly: process.env.READ_ONLY == 'true',
    socket: {
        isConnected: false,
        message: '',
        reconnectError: false,
      },
};
