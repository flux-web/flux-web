import { GetterTree } from 'vuex';
import { RootState } from '../types/RootState';
import { NamespacesState } from '../types/Namespaces/NamespacesState';

export const getters: GetterTree<NamespacesState, RootState> = {
    namespaces: (state: NamespacesState): string[]  => state.namespaces,
    currentNamespace:  (state: NamespacesState): string | null  => state.currentNamespace,
};
