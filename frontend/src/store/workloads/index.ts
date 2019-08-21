import { getters } from './getters';
import { actions } from './actions';
import { mutations } from './mutations';
import { Module } from 'vuex';
import { RootState } from '@/store/types/RootState';
import { WorkloadsState } from '@/store/types/Workloads/WorkloadsState';

import { WorkloadStatuses } from '../types/Workloads/WorkloadStatuses';

export const state: WorkloadsState = {
    workloads: [
        {
          id: 'sdpofiapedjisdasdlak',
          workload: 'production',
          container: 'frontend',
          image: 'frontend',
          status: WorkloadStatuses.updating,
          available_tags: [
              {
                tag: 'v1.1',
                date: '12:13 12/08/19',
                current: false,
              },
              {
                tag: 'v1.0',
                date: '02:23 13/08/19',
                current: true,
            },
          ],
        },
        {
          id: 'e09qudwijooiajlskcmsasdkn',
          workload: 'staging',
          container: 'backend',
          image: 'frontend',
          status: WorkloadStatuses.updating,
          available_tags: [
              {
                tag: '3refkp',
                date: '23ropkew',
                current: false,
              },
              {
                tag: '2wsefijo',
                date: 'zxlkn39',
                current: true,
            },
          ],
        },
      ],
    searchTerm: '',
};

const namespaced: boolean = true;

export const workloads: Module<WorkloadsState, RootState> = {
    namespaced,
    state,
    getters,
    actions,
    mutations,
};
