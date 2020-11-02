<template>
  <div class="available-automated-flag-select-wrapper">
    <multiselect
      v-model="value"
      :options="options"
      :placeholder="currentTag ? (currentTag.automated ? 'true' : 'false') : ''"
      label="automated"
      track-by="automated"
      :allow-empty="false"
      :disabled="isDisabled"
      deselect-label="Can't remove this value"
      @input="valueChanged"
    >
      <template slot="singleLabel" slot-scope="{ option }">
        <strong>{{ option.automated }}</strong>
      </template>
      <template slot="option" slot-scope="props">
        <div class="option__desc">
          <span>{{ props.option.automated }}</span>
        </div>
      </template>
    </multiselect>
  </div>
</template>

<style lang="scss">
@import "../../assets/scss/include";
.available-automated-flag-select-wrapper {
  min-width: max-content;
  .multiselect__content-wrapper {
    min-width: max-content;
    right: 0;
  }
}
</style>

<script lang="ts">
import { Component, Vue, Prop } from "vue-property-decorator";
import Multiselect from "vue-multiselect";

@Component({
  components: {
    Multiselect
  }
})
export default class WorkloadAvailableAutomatedFlag extends Vue {
  @Prop() protected currentTag: any;
  @Prop() protected workload: any;

  protected value: any = null;
  protected isDisabled: Boolean = false;

  protected possibleOptions: any = [
    { automated: true },
    { automated: false },
  ];

  public updated() {
    this.refreshSelectedValue();
  }

  public refreshSelectedValue() {
    console.dir(this.workload);
    if (this.workload.selected_tag.tag && this.workload.selected_tag.tag != this.workload.available_tags[0].tag) {
      this.isDisabled = true;
    } else {
      this.isDisabled = false;
    }
  }

  get options() {
    return this.possibleOptions;
  }

  public valueChanged() {
    this.$emit("input", this.workload, this.value);
  }
}
</script>
