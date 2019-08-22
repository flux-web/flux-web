import { ActionTree } from 'vuex';
import { RootState } from '../types/RootState';
import { NamespacesState } from '../types/Namespaces/NamespacesState';
import axios from 'axios';
import { Namespace } from '../types/Namespaces/Namespace';

export const actions: ActionTree<NamespacesState, RootState> = {
    fetchNamespaces: ({commit}): any => axios.get('/namespaces').then(
        (response) => commit('UPDATE_NAMESPACES', response.data),
    ),
    setCurrentNamespace: ({commit}, namespace: Namespace): any => commit('SET_CURRENT_NAMESPACE', namespace)
};
