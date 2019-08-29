import { RootState } from '~/store/types/RootState';


export const root: RootState = {
    socket: {
        isConnected: false,
        message: '',
        reconnectError: false,
      },
};
