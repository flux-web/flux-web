import { MutationTree } from 'vuex';
import { NamespacesState } from '../types/Namespaces/NamespacesState';
import { Namespace } from '../types/Namespaces/Namespace';

export const mutations: MutationTree<NamespacesState> = {
    UPDATE_NAMESPACES: (state: NamespacesState, namespaces: Namespace[]) => state.namespaces = namespaces,
    SET_CURRENT_NAMESPACE: (state: NamespacesState, namespace: Namespace) => state.currentNamespace = namespace,
};
