import { MutationTree } from 'vuex';
import { WorkloadsState } from '../types/Workloads/WorkloadsState';
import { Workload } from '../types/Workloads/Workload';
import { WorkloadStatuses } from '../types/Workloads/WorkloadStatuses';

export const mutations: MutationTree<WorkloadsState> = {
    CHANGE_SEARCH_TERM: (state: WorkloadsState, searchTerm: string) => state.searchTerm = searchTerm,
    UPDATE_WORKLOADS: (state: WorkloadsState, workloads: Workload[]) => state.workloads = workloads,
    UPDATE_WORKLOAD_STATUS: (state: WorkloadsState, workload: Workload, status: WorkloadStatuses) => () => {
        const workloadInst = state.workloads.find(w => w.id == workload.id);
        if (!workloadInst) {
            throw `Unable to update workload, workload (${workload.id}) not found`;
        }
        workloadInst.status = status;
    },
};
