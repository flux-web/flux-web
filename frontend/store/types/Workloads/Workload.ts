import { Tag } from './Tag';
import { WorkloadStatuses } from './WorkloadStatuses';
import { WorkloadTypes } from "~/store/types/Workloads/WorkloadTypes";

export interface Workload {
    id: string;
    workload: string;
    container: string;
    type: WorkloadTypes;
    image: string;
    status: WorkloadStatuses;
    current_tag: Tag;
    available_tags: Tag[];
    selected_tag: Tag;
}
