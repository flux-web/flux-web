<template>
  <div class="workloads-list"></div>
</template>

<script lang="ts">
import { Component, Vue, Watch } from "vue-property-decorator";

@Component({})
export default class WorkloadsList extends Vue {}
</script>

<style scoped lang="scss">
@import "../../assets/scss/include";

.workloads-list {
  height: calc(100% - 110px);
  min-height: 250px;
  padding: 15px 0;
  box-sizing: border-box;
  overflow-y: scroll;
}
</style>
