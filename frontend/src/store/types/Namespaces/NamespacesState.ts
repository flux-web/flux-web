import { Namespace } from './Namespace';

export interface NamespacesState {
    namespaces: Namespace[];
    currentNamespace: Namespace | null;
}
