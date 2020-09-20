<template>
  <div class="workloads-search">
    <input
      type="text"
      name="search"
      class="text-input"
      placeholder="Search"
      autocomplete="off"
      v-model="searchTerm"
      @input="searchTermChanged"
    />
  </div>
</template>

<style scoped lang="scss">
@import "../../assets/scss/include";

.workloads-search {
  position: relative;
  float: right;
  width: 300px;
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
    color: #6a6c71;
    &::placeholder {
      color: #b6b9c3;
    }
    &:focus {
      outline: none;
      padding: 0 0 0 29px;
      color: #6a6c71;
      border: 1px solid #b6b9c3;
    }
  }
  .input-icon {
    position: absolute;
    right: 10px;
    top: 50%;
    color: #9095a4;
    transform: translateY(-40%);
  }
}
</style>

<script lang="ts">
import { Component, Vue } from "vue-property-decorator";
import { StoreNamespaces } from "../../store/types/StoreNamespaces";
import { Action } from "vuex-class";

@Component({})
export default class WorkloadSearch extends Vue {
  public searchTerm: string = "";

  @Action("changeSearchTerm", { namespace: StoreNamespaces.workloads })
  protected changeSearchTerm: any;

  public searchTermChanged() {
    this.changeSearchTerm(this.searchTerm);
  }
}
</script>