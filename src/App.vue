<template>
  <base-layout>
    <router-view />
  </base-layout>
</template>

<style></style>

<script lang="ts">
import { Component, Vue } from "vue-property-decorator";
import BaseLayout from "@/layouts/BaseLayout.vue";

@Component({
  name: "App",
  components: { BaseLayout }
})
export default class App extends Vue {}
</script>
