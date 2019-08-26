<template>
  <button
    class="release_button"
    @click="release(workload)"
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

  public release() {
    const releaseData: any = {
      Cause: {
        Message: "",
        User: "Flux-web"
      },
      Spec: {
        ContainerSpecs: {},
        Kind: "execute",
        SkipMismatches: true
      },
      Type: "containers"
    };

    releaseData.Spec.ContainerSpecs[this.workload.id] = [
      {
        Container: this.workload.container,
        Current: this.workload.image + ":" + this.workload.current_tag.tag,
        Target: this.workload.image + ":" + this.workload.selected_tag.tag
      }
    ];

    this.releaseVersion({ workload: this.workload, releaseData });
  }
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