import { MutationTree } from 'vuex';
import Vue from 'vue';
import { RootState } from '~/store/types/RootState';
import { StoreNamespaces } from '../types/StoreNamespaces';
import store from '../index';

export const mutations: MutationTree<RootState> = {
    SOCKET_ONOPEN: (state: RootState, event: any) => {
        Vue.prototype.$socket = event.currentTarget;
        state.socket.isConnected = true;
    },
    SOCKET_ONCLOSE: (state: RootState) => state.socket.isConnected = false,
    SOCKET_ONERROR: (state: RootState, event: any) => console.error(state, event),
    SOCKET_ONMESSAGE: (state: RootState, message: any) => {
        console.info('Got a new message from ws:', message);
        store().dispatch(StoreNamespaces.workloads + '/updateWorloadRelease', {
            workload: {
                id: message.Workload,
                container: message.Container,
            },
            tag: message.Tag.split(':').pop(),
        })
    },
    SOCKET_RECONNECT: (state: RootState, count: number) => console.info(state, count),
    SOCKET_RECONNECT_ERROR: (state: RootState) => state.socket.reconnectError = true,
};
