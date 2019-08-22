<template>
  <div class="workloads-wrapper">
    <workloads-search></workloads-search>
    <workloads-list></workloads-list>
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
import { StoreNamespaces } from "../../store/types/StoreNamespaces";
import { Action, Getter } from "vuex-class";

@Component({
  components: { WorkloadsSearch, WorkloadsList }
})
export default class Workloads extends Vue {
  private TAG = `[${Workloads.name}]`;

  @Action("fetchWorkloads", { namespace: StoreNamespaces.workloads })
  public fetchWorkloads: any;

  @Action("fetchNamespaces", { namespace: StoreNamespaces.namespaces })
  public fetchNamespaces: any;

  @Action("setCurrentNamespace", { namespace: StoreNamespaces.namespaces })
  public setCurrentNamespace: any;

  @Getter("currentNamespace", { namespace: StoreNamespaces.namespaces })
  public currentNamespace: any;

  public async mounted() {
    let namespaces = await this.fetchNamespaces();
    if (!namespaces || !namespaces.length) {
      throw "Couldn't fetch namespaces, please check that backend is up and running";
      return;
    }
    this.setCurrentNamespace(namespaces[0]);

    await this.fetchWorkloads(this.currentNamespace);
  }
}
</script>