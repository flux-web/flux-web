<template>
  <div class="workloads-logs">
    <div class="message-wrapper">
      <input
        type="text"
        name="search"
        class="text-input"
        placeholder="Last log"
        disabled="disabled"
        autocomplete="off"
      />
    </div>
  </div>
</template>

<script lang="ts">
import { Component, Prop, Vue } from "vue-property-decorator";
import { Action, Getter } from "vuex-class";

@Component({})
export default class WorkloadsLogs extends Vue {}
</script>

<style scoped lang="scss">
@import "../../assets/scss/include";

.workloads-logs {
  position: relative;
  width: 100%;
  height: 55px;
  border-top: solid 1px #f0f0fb;
  .message-wrapper {
    padding: 25px 0 0 0;
    height: 55px;
    box-sizing: border-box;
    .text-input {
      height: 42px;
      border-radius: 7px;
      background-color: #f0f0fb;
      border: none;
      width: 100%;
      box-sizing: border-box;
      padding: 0 0 0 30px;
      font-family: SpoilerHE;
      font-weight: 300;
      font-size: 13px;
      color: #b6b9c3;
      float: right;
      &::placeholder {
        color: #b6b9c3;
      }
      &:focus {
        outline: none;
        padding: 0 0 0 29px;
        border: 1px solid #b6b9c3;
        border-left: none;
        & + .send-button {
          border: 1px solid #b6b9c3;
          border-left: none;
        }
      }
    }
  }
}
</style>
