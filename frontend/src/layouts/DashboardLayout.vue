<template>
  <div class="container">
    <div class="top-left-circle">
      <div class="inner-circle">
        <div class="core-circle"></div>
      </div>
    </div>
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
import Navbar from "@/components/Navbar.vue";

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
  height: 100vh;
  box-sizing: border-box;
  padding: 5vh 5vw;
  box-shadow: 15px 20px 99px 0 rgba(82, 80, 94, 0.33);
  position: relative;
  @include respond-above("Lmedium") {
    padding: 3vh 3vw;
  }
  .top-left-circle {
    width: 600px;
    height: 600px;
    background-color: rgba(255, 255, 255, 0.02);
    position: absolute;
    top: -300px;
    left: -300px;
    border-radius: 300px;
    .inner-circle {
      width: 400px;
      height: 400px;
      top: 100px;
      left: 100px;
      border-radius: 200px;
      position: absolute;
      background-color: rgba(255, 255, 255, 0.02);
      .core-circle {
        width: 200px;
        height: 200px;
        top: 100px;
        left: 100px;
        border-radius: 100px;
        position: absolute;
        background-color: rgba(255, 255, 255, 0.02);
      }
    }
  }
  @include respond-above("Lmedium") {
    flex-direction: column;
  }
  .main-wrapper {
    display: flex;
    flex-direction: row;
    justify-content: center;
    z-index: 9999;
    position: relative;
    box-shadow: 15px 20px 99px 0 rgba(82, 80, 94, 0.33);
    height: 90vh;
    min-height: 450px;
    border-radius: 13px;
    @include respond-above("Lmedium") {
      display: block;
      height: 88vh;
    }
    .left-side {
      @include respond-above("Lmedium") {
        height: 50px;
        border-radius: 13px 13px 0 0;
        width: 100%;
      }
      width: 60px;
      background: rgb(38, 50, 80);
      background: linear-gradient(
        180deg,
        rgba(38, 50, 80, 1) 0%,
        rgba(59, 68, 106, 1) 100%
      );
      border-radius: 13px 0 0 13px;
    }
    .right-side {
      @include respond-above("Lmedium") {
        border-radius: 0 0 13px 13px;
      }
      @include respond-above("Lmedium") {
      }
      border-radius: 0 13px 13px 0;
      position: relative;
      background: #fff;
      width: 100%;
      height: 100%;
      padding: 15px;
      box-sizing: border-box;
    }
  }
}
</style>
