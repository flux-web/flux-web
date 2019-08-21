import { MutationTree } from 'vuex';
import { WorkloadsState } from '../types/Workloads/WorkloadsState';

export const mutations: MutationTree<WorkloadsState> = {
    CHANGE_SEARCH_TERM: (state: WorkloadsState, searchTerm: string) => state.searchTerm = searchTerm,
};
