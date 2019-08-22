import { getters } from './getters';
import { actions } from './actions';
import { mutations } from './mutations';
import { Module } from 'vuex';
import { RootState } from '@/store/types/RootState';
import { WorkloadsState } from '@/store/types/Workloads/WorkloadsState';

import { WorkloadStatuses } from '../types/Workloads/WorkloadStatuses';

export const state: WorkloadsState = {
    workloads: [],
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
