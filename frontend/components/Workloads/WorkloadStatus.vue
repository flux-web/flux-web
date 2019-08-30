<template>
  <div>
    <div class="icon" :style="iconStyle"></div>
  </div>
</template>

<style scoped lang="scss">
.icon {
  background-position: center center;
  background-size: contain;
  width: 35px;
  height: 35px;
}
</style>

<script lang="ts">
import { Component, Vue, Prop } from "vue-property-decorator";
import { Getter } from "vuex-class";
import { StoreNamespaces } from "../../store/types/StoreNamespaces";
import { WorkloadStatuses } from "../../store/types/Workloads/WorkloadStatuses";

@Component({})
export default class WorkloadStatus extends Vue {
  @Prop() protected workload: any;

  @Getter("getWorkload", { namespace: StoreNamespaces.workloads })
  public getWorkload: any;

  get iconStyle() {
    const status = this.getWorkload(this.workload).status;
    const imageName = status;
    const imageExtension = status == WorkloadStatuses.releasing ? "gif" : "png";
    return {
      backgroundImage: `url(${require(`@/assets/images/${imageName}.${imageExtension}`)})`
    };
  }
}
</script>