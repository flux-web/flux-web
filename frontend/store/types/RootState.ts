export interface RootState {
    socket: {
        isConnected: boolean;
        message: any;
        reconnectError: boolean;
      };
}
