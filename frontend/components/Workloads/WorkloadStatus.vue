<template>
  <div>
    <div
      :class="{ releasing: getWorkload(workload).status == 'releasing'}"
    >{{getWorkload(workload).status !== 'releasing' ? getWorkload(workload).status : ''}}</div>
  </div>
</template>

<style lang="scss">
.releasing,
.releasing:after {
  border-radius: 50%;
  width: 15px;
  height: 16px;
}
.releasing {
  margin: 5px auto;
  font-size: 10px;
  position: relative;
  text-indent: -9999em;
  border-top: 5px solid rgb(210, 213, 224);
  border-right: 5px solid rgb(210, 213, 224);
  border-bottom: 5px solid rgb(210, 213, 224);
  border-left: 5px solid #2f3b5c;
  -webkit-transform: translateZ(0);
  transform: translateZ(0);
  -webkit-animation: load8 1.1s infinite linear;
  animation: load8 1.1s infinite linear;
  background: #ffffff;
}
@-webkit-keyframes load8 {
  0% {
    -webkit-transform: rotate(0deg);
    transform: rotate(0deg);
  }
  100% {
    -webkit-transform: rotate(360deg);
    transform: rotate(360deg);
  }
}
@keyframes load8 {
  0% {
    -webkit-transform: rotate(0deg);
    transform: rotate(0deg);
  }
  100% {
    -webkit-transform: rotate(360deg);
    transform: rotate(360deg);
  }
}
</style>

<script lang="ts">
import { Component, Vue, Prop } from "vue-property-decorator";
import { Getter } from "vuex-class";
import { StoreNamespaces } from "../../store/types/StoreNamespaces";

@Component({})
export default class WorkloadStatus extends Vue {
  @Prop() protected workload: any;

  @Getter("getWorkload", { namespace: StoreNamespaces.workloads })
  public getWorkload: any;
}
</script>