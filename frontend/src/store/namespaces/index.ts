import { getters } from './getters';
import { actions } from './actions';
import { mutations } from './mutations';
import { Module } from 'vuex';
import { RootState } from '@/store/types/RootState';
import { NamespacesState } from '@/store/types/Namespaces/NamespacesState';

export const state: NamespacesState = {
    namespaces: [
        {
          name: 'production',
        },
    ],
    currentNamespace: null,
};

const namespaced: boolean = true;

export const namespaces: Module<NamespacesState, RootState> = {
    namespaced,
    state,
    getters,
    actions,
    mutations,
};
