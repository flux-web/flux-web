<template>
  <dashboard-layout>
    <template slot="right-side">Flux web about page</template>
  </dashboard-layout>
</template>

<script lang="ts">
import { Component, Vue } from "vue-property-decorator";
import DashboardLayout from "@/layouts/DashboardLayout.vue";
import Workloads from "@/components/Workloads/Workloads.vue";

@Component({
  components: {
    DashboardLayout,
    Workloads
  }
})
export default class About extends Vue {}
</script>

<style lang="scss"></style>
