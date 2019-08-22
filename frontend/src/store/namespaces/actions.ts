import { ActionTree } from 'vuex';
import { RootState } from '../types/RootState';
import { NamespacesState } from '../types/Namespaces/NamespacesState';
import axios from 'axios';

export const actions: ActionTree<NamespacesState, RootState> = {
    fetchNamespaces: ({commit}): any => axios.get('/namespaces').then(
        (response) => commit('UPDATE_NAMESPACES', response.data),
    ),
    setCurrentNamespace: ({commit}, namespace: string): any => commit('SET_CURRENT_NAMESPACE', namespace),
};
