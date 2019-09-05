<template>
  <div class="navbar">
    <div class="info">
      <div class="details">
        <span class="details__item details__label cr__label">Flux Web</span>
      </div>
      <ul class="icons">
        <li class="icon">
          <router-link to="/">
            <i class="fas fa-rocket"></i>
          </router-link>
        </li>
        <li class="icon">
          <router-link to="/about">
            <i class="fas fa-info"></i>
          </router-link>
        </li>
        <li class="icon">
          <a :href="githubUrl" target="__blank">
            <i class="fab fa-github"></i>
          </a>
        </li>
      </ul>
    </div>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from "vue-property-decorator";
import Config from "~/config";

@Component({})
export default class Navbar extends Vue {
  public githubUrl = Config.config.githubUrl;
}
</script>

<style scoped lang="scss">
@import "../assets/scss/include";

.navbar {
  padding: 20px;
  height: 100%;
  align-items: center;
  box-sizing: border-box;
  display: flex;
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
    .icons {
      list-style-type: none;
      padding: 0px;
      margin: 20px 0 0 0;
      @include respond-above("Lmedium") {
        margin: 10px 0 0 0;
      }
      .icon {
        color: #99a9c2;
        @include respond-above("Lmedium") {
          display: inline-block;
        }
        &:hover {
          color: #fff;
        }
        a {
          padding: 25px 0;
          display: block;
          color: inherit;
          @include respond-above("Lmedium") {
            padding: 0 15px;
          }
        }
      }
    }
  }
}
</style>
