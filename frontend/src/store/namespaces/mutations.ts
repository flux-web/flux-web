import { MutationTree } from 'vuex';
import { NamespacesState } from '../types/Namespaces/NamespacesState';

export const mutations: MutationTree<NamespacesState> = {
    UPDATE_NAMESPACES: (state: NamespacesState, namespaces: string[]) => state.namespaces = namespaces,
    SET_CURRENT_NAMESPACE: (state: NamespacesState, namespace: string) => state.currentNamespace = namespace,
};
