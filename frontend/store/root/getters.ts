import { GetterTree } from 'vuex';
import { RootState } from '../types/RootState';

export const getters: GetterTree<RootState, RootState> = {
    message : (state: RootState): string  => state.socket.message,
    readOnly : (state: RootState): boolean  => state.readOnly,
};
