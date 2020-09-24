<template>
  <div class="namespace-select" v-show="namespaces.length">
    <multiselect
      v-model="namespace"
      :options="options"
      label="name"
      track-by="name"
      :placeholder="namespace ? namespace.name : 'Select namespace'"
      :allow-empty="false"
      deselect-label="Selected"
      @input="selectNamespace"
    >
      <template slot="singleLabel" slot-scope="{ option }">
        <strong>{{ option.name }}</strong>
      </template>
      <template slot="option" slot-scope="props">
        <div class="option__desc">
          <span class="option__tag">{{ props.option.name }}</span>
        </div>
      </template>
    </multiselect>
    <div
      class="namespace-status"
    >{{loading ? 'Loading, please wait' : ''}}</div>
  </div>
</template>

<script lang="ts">
import Multiselect from "vue-multiselect";
import { Component, Vue } from "vue-property-decorator";
import { StoreNamespaces } from "../../store/types/StoreNamespaces";
import { Action, Getter } from "vuex-class";

@Component({
  components: {
    Multiselect
  }
})
export default class NamespaceSelect extends Vue {
  public namespace: any = {};

  public options: Array<any> = [];

  public loading: boolean = false;

  @Action("setCurrentNamespace", { namespace: StoreNamespaces.namespaces })
  public setCurrentNamespace: any;

  @Action("fetchNamespaces", { namespace: StoreNamespaces.namespaces })
  public fetchNamespaces: any;

  @Action("fetchWorkloads", { namespace: StoreNamespaces.workloads })
  public fetchWorkloads: any;

  @Getter("currentNamespace", { namespace: StoreNamespaces.namespaces })
  public currentNamespace: any;

  @Getter("namespaces", { namespace: StoreNamespaces.namespaces })
  public namespaces: any;

  public async mounted() {
    try {
      await this.fetchNamespaces();
    } catch (e) {
      throw "Error fetching namespaces";
    }
    this.options = this.namespaces.map(n => ({name: n}))
    this.namespace = this.currentNamespace;
    if (this.namespace) {
      this.selectNamespace();
    }
  }

  public async selectNamespace() {
    if (this.loading) {
      return;
    }
    this.setCurrentNamespace(this.namespace);

    this.loading = true;

    try {
      await this.fetchWorkloads(this.currentNamespace.name);
    } catch (e) {
      alert(
        "Error when retrieving workloads for namespace: " +
          this.currentNamespace.name
      );
    }

    this.loading = false;
  }
}
</script>

<style scoped lang="scss">
@import "../../assets/scss/include";

.namespace-select {
  margin-bottom: 10px;
  display: inline-block;
  display: flex;
  align-items: center;
  align-content: center;
  .multiselect {
    width: 300px;
  }
  .namespace-status {
    margin-left: 10px;
  }
  .namespace-input {
    height: 30px;
    border-radius: 7px;
    background-color: #f0f0fb;
    border: none;
    width: 200px;
    box-sizing: border-box;
    padding: 0 0 0 15px;
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

  .namespace-button {
    margin-left: 10px;
    background-color: #f0f0fb;
    border: none;
    font-family: sans-serif;
    font-weight: 100;
    font-size: 13px;
    color: #97989c;
    height: 30px;
    border-radius: 7px;
    padding: 8px;
  }
}
</style>
