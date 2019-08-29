import { ActionTree } from 'vuex';
import { RootState } from '../types/RootState';
import { WorkloadsState } from '../types/Workloads/WorkloadsState';
import { workloadsTransformer } from './transformers';
import { Workload } from '../types/Workloads/Workload';
import { WorkloadStatuses } from '../types/Workloads/WorkloadStatuses';
import axios from 'axios';

export const actions: ActionTree<WorkloadsState, RootState> = {
    changeSearchTerm: ({commit}, searchTerm: string) => commit('CHANGE_SEARCH_TERM', searchTerm),
    updateWorkloadStatus:  ({commit}, payload) => commit('UPDATE_WORKLOAD_STATUS', payload),
    updateWorloadRelease:  ({commit, getters}, {workload, tag}) => {
        const storeWorkload = getters.getWorkload(workload);
        const isStatusUpToDate = (workload: Workload) => tag == 'latest' || ( workload.available_tags.length && tag == workload.available_tags[0].tag)
        commit('UPDATE_WORKLOAD_CURRENT_TAG', {
            workload: storeWorkload,
            tag
        })
        commit('UPDATE_WORKLOAD_STATUS', {
            workload: storeWorkload, 
            status: isStatusUpToDate(storeWorkload) ? WorkloadStatuses.upToDate : WorkloadStatuses.behind
        })
    },
    updateSelectedTag:  ({commit}, payload) => commit('UPDATE_SELECTED_TAG', payload),
    fetchWorkloads: ({commit}, namespace: string): any => axios.get('/workloads/' + namespace).then(
        (response: any) => {
            const workloads = workloadsTransformer(response.data);
            commit('UPDATE_WORKLOADS', workloads);
        },
    ),
    releaseVersion: ({dispatch}, {workload, releaseData}): any => {
        axios.post('/release', releaseData).then(
            ()  => dispatch('updateWorkloadStatus', {workload, status: WorkloadStatuses.releasing}),
        );
    },
};
