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

            const availableTags = container.Available ? container.Available.map((available: any) => {
                const availableTag = available.ID.split(':').pop();
                return {
                    tag: available.ID.split(':').pop(),
                    date: available.CreatedAt,
                    current: availableTag === currentTag,
                  };
            }) : [];

            const temp = {
                id: workload.ID,
                workload: workload.ID.split(':').pop(),
                container: container.Name,
                image: getImageFromUrl(container.Current.ID),
                status: WorkloadStatuses.upToDate,
                available_tags: availableTags,
                current_tag: {
                    tag: currentTag,
                    current: true,
                    date: container.Current.CreatedAt || null,
                },
              };
            containerWorkloads.push(temp);
            return containerWorkloads;
        }, []));
    }, []);
};
