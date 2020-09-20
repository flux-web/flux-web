<template>
  <div class="container">
    <div class="main-wrapper">
      <div class="left-side">
        <navbar></navbar>
      </div>
      <div class="right-side">
        <slot name="right-side"></slot>
      </div>
    </div>
  </div>
</template>


<script lang="ts">
import { Component, Vue } from "vue-property-decorator";
import Navbar from "~/components/Navbar.vue";

@Component({
  components: {
    Navbar
  }
})
export default class CallScreenLayout extends Vue {}
</script>

<style scoped lang="scss">
@import "../assets/scss/include";

.container {
  height: 100%;
  width: 100%;
  height: 100%;
  left: 0;
  top: 0;

  box-sizing: border-box;
  position: relative;
  @include respond-above("Lmedium") {
    flex-direction: column;
  }

  .main-wrapper {
    display: flex;
    flex-direction: row;
    justify-content: center;
    z-index: 9999;
    position: relative;
    height: 100%;
    min-height: 450px;
    @include respond-above("Lmedium") {
      display: block;
      height: 88vh;
    }
    .left-side {
      @include respond-above("Lmedium") {
        height: 60px;
        width: 100%;
      }
      width: 60px;
      background: rgb(38, 50, 80);
      background: linear-gradient(
        180deg,
        rgba(38, 50, 80, 1) 0%,
        rgba(59, 68, 106, 1) 100%
      );
    }
    .right-side {
      position: relative;
      background: #fff;
      width: calc(100% - 60px);
      height: 100%;
      padding: 15px;
      box-sizing: border-box;
    }
  }
}
</style>
