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
    if (!workloads) {
        return []
    }
    return workloads.reduce((accWorkloads: any, workload: any) => {
        if (!workload.Containers) {
            return accWorkloads;
        }

        return accWorkloads.concat(workload.Containers.reduce((containerWorkloads: any, container: any) => {
            const current = (container.Current.ID && container.Current) || (container.Available && container.Available[0]) || {}
            const currentTag = parseCurrentTag(current.ID)

            let availableTagList = []
            if (container.Available) {

                availableTagList = container.Available.map((available: any) => {
                    const availableTag = available.ID.split(':').pop();
                    return {
                        tag: available.ID.split(':').pop(),
                        date: available.CreatedAt,
                        current: availableTag === currentTag,
                    };
                })

                if (container.Name === "chart-image") {
                    let availableTag = JSON.parse(JSON.stringify(availableTagList[0]));
                    availableTag.automated = true;
                    availableTagList.unshift(availableTag)
                }

            }
            const availableTags = availableTagList

            const isStatusUpToDate = (availableTags: Tag[], currentTag: string) => currentTag == 'latest' || ( availableTags.length && currentTag == availableTags[0].tag)

            const temp = {
                id: workload.ID,
                workload: workload.ID.split(':').pop(),
                container: container.Name,
                image: getImageFromUrl(current.ID),
                status: isStatusUpToDate(availableTags, currentTag) ? WorkloadStatuses.upToDate : WorkloadStatuses.behind,
                available_tags: availableTags,
                current_tag: {
                    tag: currentTag,
                    current: true,
                    date: current.CreatedAt || null,
                    automated: workload.Automated,
                },
                selected_tag: {},
            };
            containerWorkloads.push(temp);
            return containerWorkloads;
        }, []));
    }, []);
};
