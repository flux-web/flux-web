import { ActionTree } from 'vuex';
import { RootState } from '../types/RootState';
import { WorkloadsState } from '../types/Workloads/WorkloadsState';
import { workloadsTransformer } from './transformers';
import axios from 'axios';
import { Workload } from '../types/Workloads/Workload';
import { WorkloadStatuses } from '../types/Workloads/WorkloadStatuses';

export const actions: ActionTree<WorkloadsState, RootState> = {
    changeSearchTerm: ({commit}, searchTerm: string) => commit('CHANGE_SEARCH_TERM', searchTerm),
    updateWorkloadStatus:  ({commit}, status: WorkloadStatuses) => commit('UPDATE_WORKLOAD_STATUS', status),
    fetchWorkloads: ({commit}, namespace: string): any => axios.get('/workloads/' + namespace).then(
        (response) => {
            const workloads = workloadsTransformer(response.data);
            commit('UPDATE_WORKLOADS', workloads);
        },
    ),
    releaseVersion: ({dispatch}, workload: Workload): any => {
        const releaseData: any = {
            Cause: {
              Message: "",
              User: "Flux-web"
            },
            Spec: {
              ContainerSpecs: {},
              Kind: "execute",
              SkipMismatches: true
            },
            Type: "containers"
          };
      
          releaseData.Spec.ContainerSpecs[workload.id] = [
            {
              Container: workload.container,
              Current: workload.image + ":" + workload.current_tag.tag,
              Target: workload.image + ":" + workload.selected_tag.tag
            }
          ];

        axios.post('/release', releaseData).then(
            ()  => dispatch('updateWorkloadStatus', workload, WorkloadStatuses.releasing),
       )
    },
};
