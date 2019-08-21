import { Workload } from './Workload';

export interface WorkloadsState {
    workloads: Workload[];
    searchTerm: string;
}
