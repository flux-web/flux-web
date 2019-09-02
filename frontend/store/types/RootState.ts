export interface RootState {
    readOnly: boolean;
    socket: {
        isConnected: boolean;
        message: any;
        reconnectError: boolean;
      };
}
