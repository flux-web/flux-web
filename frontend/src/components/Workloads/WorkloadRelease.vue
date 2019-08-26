<template>
  <button
    class="release_button"
    @click="releaseVersion(workload)"
  >{{workload.releasing ? 'Relasing' : 'Release'}}</button>
</template>

<script lang="ts">
import { Component, Vue, Prop } from "vue-property-decorator";
import { StoreNamespaces } from "../../store/types/StoreNamespaces";
import { Getter, Action } from "vuex-class";

@Component({})
export default class WorkloadRelease extends Vue {
  @Prop() protected workload: any;

  @Action("releaseVersion", { namespace: StoreNamespaces.workloads })
  public releaseVersion: any;
}
</script>

<style lang="scss">
@import "../../assets/scss/include";
.release_button {
  background: #007efe;
  padding: 8px;
  color: #fff;
  border-radius: 5px;
  font-size: 11px;
  &:hover {
    background: #3190f1;
  }
  &:focus {
    outline: 0;
  }
}
</style>