import { MutationTree } from 'vuex';
import Vue from 'vue';
import { RootState } from '@/store/types/RootState';
import { StoreNamespaces } from '../types/StoreNamespaces';
import { WorkloadStatuses } from '../types/Workloads/WorkloadStatuses';

export const mutations: MutationTree<RootState> = {
    SOCKET_ONOPEN: (state: RootState, event: any) => {
        Vue.prototype.$socket = event.currentTarget;
        state.socket.isConnected = true;
    },
    SOCKET_ONCLOSE: (state: RootState, event: any) => state.socket.isConnected = false,
    SOCKET_ONERROR: (state: RootState, event: any) => console.error(state, event),
    SOCKET_ONMESSAGE: (state: RootState, message: any) => {
        const workload = {
            id: message.data.Workload,
            container: message.data.Container,
        };
        // tslint:disable-next-line: max-line-length
        this.commit(StoreNamespaces.workloads + '/UPDATE_WORKLOAD_STATUS', {workload, status: WorkloadStatuses.upToDate});
        state.socket.message = message.data;
    },
    SOCKET_RECONNECT: (state: RootState, count: number) => console.info(state, count),
    SOCKET_RECONNECT_ERROR: (state: RootState) => state.socket.reconnectError = true,
};
