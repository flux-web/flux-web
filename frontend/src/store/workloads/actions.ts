import { ActionTree } from 'vuex';
import { RootState } from '../types/RootState';
import { WorkloadsState } from '../types/Workloads/WorkloadsState';

export const actions: ActionTree<WorkloadsState, RootState> = {
    changeSearchTerm: ({commit}, searchTerm: string) => commit('CHANGE_SEARCH_TERM', searchTerm),
};
