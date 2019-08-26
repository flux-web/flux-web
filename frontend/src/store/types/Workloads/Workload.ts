import { Tag } from './Tag';
import { WorkloadStatuses } from './WorkloadStatuses';

export interface Workload {
    id: string;
    workload: string;
    container: string;
    image: string;
    status: WorkloadStatuses;
    current_tag: Tag;
    available_tags: Tag[];
    selected_tag: Tag;
}
