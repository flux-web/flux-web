import { ActionTree } from 'vuex';
import { RootState } from '../types/RootState';
import { StoreNamespaces } from '../types/StoreNamespaces';

export const actions: ActionTree<RootState, RootState> = {
    updateRelease: ({dispatch}, message: any) => {
        dispatch(StoreNamespaces.workloads + '/updateWorloadRelease', {
            workload: {
                id: message.Workload,
                container: message.Container,
            },
            tag: message.Tag.split(':').pop(),
        })
    },
};
