<template>
  <dashboard-layout>
    <template slot="right-side">
      <workloads></workloads>
    </template>
  </dashboard-layout>
</template>

<script lang="ts">
import { Component, Vue } from "vue-property-decorator";
import DashboardLayout from "~/layouts/DashboardLayout.vue";
import Workloads from "~/components/Workloads/Workloads.vue";

@Component({
  components: {
    DashboardLayout,
    Workloads
  }
})
export default class extends Vue {}
</script>

<style lang="scss"></style>
