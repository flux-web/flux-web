<template>
  <div class="navbar">
    <div class="info">
      <div class="details">
        <span class="details__item details__label cr__label">Flux Web</span>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from "vue-property-decorator";

@Component({})
export default class Navbar extends Vue {}
</script>

<style scoped lang="scss">
@import "../assets/scss/include";

.navbar {
  padding: 20px;
  height: 10%;
  align-items: center;
  box-sizing: border-box;
  display: flex;
  justify-content: center;
  flex-direction: column;
  text-align: center;
  @include respond-above("Lmedium") {
    padding: 3px;
    height: 100%;
  }
  .info {
    .details {
      display: flex;
      align-items: center;
      justify-content: center;

      .cr__label {
        color: $gray-light-color;
        border-radius: 15px;
        padding: 5px 15px 5px 0px;
        font-size: 13px;
        font-weight: 500;
        user-select: none;
        margin: 0 0 0 15px;
        color: #bbc7dc;
        font-family: SpoilerHE;
      }
    }
  }
  .actions {
    margin: 0 15px 0 auto;
    @include respond-above("Smedium") {
      margin: 15px 0 0 0;
    }
  }
}
</style>
