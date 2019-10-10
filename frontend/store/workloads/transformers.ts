import { WorkloadStatuses } from '../types/Workloads/WorkloadStatuses';
import { Tag } from '../types/Workloads/Tag';

function getImageFromUrl(url: any) {
    url = url.split(':');
    if (url.length > 1) {
        url.pop();
    }
    return url.join(':');
}

function parseCurrentTag(currentTag: string): string {
    const tagParts = currentTag.split(':');
    return tagParts.length == 1 ? 'latest' : (tagParts.pop() || 'unknown')
}

export const workloadsTransformer = (workloads: any[]) => {
    return workloads.reduce((accWorkloads: any, workload: any) => {
        if (!workload.Containers) {
            return accWorkloads;
        }

        return accWorkloads.concat(workload.Containers.reduce((containerWorkloads: any, container: any) => {
            const currentTag = parseCurrentTag(container.Current.ID)

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
                status: container.Status,
                available_tags: availableTags,
                current_tag: {
                    tag: currentTag,
                    current: true,
                    date: container.Current.CreatedAt || null,
                },
                selected_tag: {}
              };
            containerWorkloads.push(temp);
            return containerWorkloads;
        }, []));
    }, []);
};
