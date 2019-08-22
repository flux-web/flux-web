import { WorkloadStatuses } from '../types/Workloads/WorkloadStatuses';

function getImageFromUrl(url: any) {
    url = url.split(':');
    url.pop();
    return url.join(':');
}

export const workloadsTransformer = (workloads: any[]) => {
    return workloads.reduce((accWorkloads: any, workload: any) => {
        if (!workload.Containers) {
            return accWorkloads;
        }

        return accWorkloads.concat(workload.Containers.reduce((containerWorkloads: any, container: any) => {
            const currentTag = container.Current.ID.split(':').pop() || 'latest';
            containerWorkloads.push({
                id: workload.ID,
                workload: workload.ID.split(':').pop(),
                container: container.Name,
                image: getImageFromUrl(container.Current.ID),
                status: WorkloadStatuses.upToDate,
                current_tag: currentTag,
                available_tags: container.Available ? container.Available.map((available: any) => {
                    const availableTag = available.ID.split(':').pop();
                    return {
                        tag: available.ID.split(':').pop(),
                        date: available.CreatedAt,
                        current: availableTag == currentTag,
                      };
                }) : [],
              });
            return containerWorkloads;
        }, []));
    }, []);
};
