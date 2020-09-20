<template>
  <div class="workloads-wrapper">
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
import WorkloadsList from "~/components/Workloads/WorkloadsList.vue";
import { StoreNamespaces } from "../../store/types/StoreNamespaces";
import { Action, Getter } from "vuex-class";

@Component({
  components: { WorkloadsList }
})
export default class Workloads extends Vue {}
</script>