import { GetterTree } from 'vuex';
import { RootState } from '../types/RootState';
import { NamespacesState } from '../types/Namespaces/NamespacesState';
import { Namespace } from '../types/Namespaces/Namespace';

export const getters: GetterTree<NamespacesState, RootState> = {
    namespaces: (state: NamespacesState): Namespace[]  => state.namespaces,
    currentNamespace:  (state: NamespacesState): Namespace | null  => state.currentNamespace,
};
