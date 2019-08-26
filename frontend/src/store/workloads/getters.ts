import { GetterTree } from 'vuex';
import { RootState } from '../types/RootState';
import { WorkloadsState } from '../types/Workloads/WorkloadsState';
import { Workload } from '../types/Workloads/Workload';

export const getters: GetterTree<WorkloadsState, RootState> = {
    workloads : (state: WorkloadsState): Workload[]  => state.workloads,
    searchTerm : (state: WorkloadsState): string  => state.searchTerm,
    getWorkload : (state: WorkloadsState): any  => (workload: Workload) => state.workloads.find((w: Workload) => (workload.id == w.id && w.container == workload.container)),
};
