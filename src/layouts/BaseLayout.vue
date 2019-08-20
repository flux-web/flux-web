<template>
  <div id="app" class="main-container">
    <slot></slot>
  </div>
</template>

<style scoped lang="scss">
.main-container {
  margin: 0 auto;
  width: 100%;
  background-image: #fff;
}
</style>

<script lang="ts">
import { Component, Vue } from "vue-property-decorator";

@Component
export default class BaseLayout extends Vue {}
</script>
