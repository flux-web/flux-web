<template>
  <div class="workloads-wrapper">
    <workloads-search></workloads-search>
    <workloads-list></workloads-list>
    <workloads-logs></workloads-logs>
  </div>
</template>

<style scoped lang="scss">
@import "../../assets/scss/include";

.workloads-wrapper {
  height: 100%;
}
</style>

<script lang="ts">
import { Component, Vue, Watch } from "vue-property-decorator";
import WorkloadsSearch from "@/components/Workloads/WorkloadsSearch.vue";
import WorkloadsList from "@/components/Workloads/WorkloadsList.vue";
import WorkloadsLogs from "@/components/Workloads/WorkloadsLogs.vue";

@Component({
  components: { WorkloadsSearch, WorkloadsList, WorkloadsLogs }
})
export default class Workloads extends Vue {
  private TAG = `[${Workloads.name}]`;
}
</script>