<template>
  <div class="workloads-list">
    <namespace-select></namespace-select>
    <vue-good-table
      :columns="columns"
      :rows="workloads"
      :search-options="{
        enabled: true,
        externalQuery: searchTerm
      }"
    >
      <template slot="table-row" slot-scope="props">
        <span v-if="props.column.field == 'release'">
          <button
            class="release_button"
            @click="releaseVersion(props.row)"
          >{{props.row.releasing ? 'Relasing' : 'Release'}}</button>
        </span>
        <workload-available-tags
          :options-prop="props.row.available_tags"
          :workload="props.row"
          @input="tagChanged"
          v-if="props.column.field == 'available_tags'"
        />
        <span v-else>{{props.formattedRow[props.column.field]}}</span>
      </template>
    </vue-good-table>
  </div>
</template>

<script lang="ts">
import { Component, Vue, Watch } from "vue-property-decorator";
import { StoreNamespaces } from "../../store/types/StoreNamespaces";
import WorkloadAvailableTags from "./WorkloadAvailableTags.vue";
import NamespaceSelect from "./NamespaceSelect.vue";
import { namespace } from "vuex-class";
import { Getter, Action } from "vuex-class";
import { Workload } from "../../store/types/Workloads/Workload";
import { Tag } from "../../store/types/Workloads/Tag";

@Component({
  components: { WorkloadAvailableTags, NamespaceSelect }
})
export default class WorkloadsList extends Vue {
  public columns = [
    {
      label: "Workload",
      field: "workload"
    },
    {
      label: "Container",
      field: "container"
    },
    {
      label: "Image",
      field: "image"
    },
    {
      label: "Current tag",
      field: "current_tag.tag"
    },
    {
      label: "Available tags",
      field: "available_tags"
    },
    {
      label: "Release",
      field: "release"
    },
    {
      label: "Status",
      field: "status"
    }
  ];

  @Getter("workloads", { namespace: StoreNamespaces.workloads })
  protected workloads!: any;

  @Getter("searchTerm", { namespace: StoreNamespaces.workloads })
  protected searchTerm!: any;

  @Action("fetchNamespaces", { namespace: StoreNamespaces.namespaces })
  public fetchNamespaces: any;

  @Action("releaseVersion", { namespace: StoreNamespaces.workloads })
  public releaseVersion: any;

  public tagChanged(workload: Workload, value: Tag) {
    const w = this.workloads.find((w: Workload) => (workload.id = w.id));
    w.selected_tag = value;
  }
}
</script>

<style src="vue-multiselect/dist/vue-multiselect.min.css"></style>

<style lang="scss">
@import "../../assets/scss/include";

.workloads-list {
  height: calc(100% - 110px);
  min-height: 250px;
  padding: 15px 0;
  box-sizing: border-box;
  overflow-y: scroll;

  .vgt-responsive {
    overflow-x: initial;
  }
  .vgt-table {
    border: none;
    font-family: sans-serif;
    &.bordered {
      td,
      th {
        border: none;
        background: none;
      }
      td {
        color: #3c5171;
        font-size: 15px;
        .release_button {
          background: #007efe;
          padding: 8px;
          color: #fff;
          border-radius: 5px;
          font-size: 11px;
          &:hover {
            background: #3190f1;
          }
          &:focus {
            outline: 0;
          }
        }
      }
      th {
        color: #9aa9c2;
        font-weight: 400;
        font-size: 14px;
      }
    }
  }
}
</style>