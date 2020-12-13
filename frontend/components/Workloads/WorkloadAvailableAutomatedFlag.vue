<template>
  <div class="available-automated-flag-select-wrapper">
    <multiselect
      v-model="value"
      :options="options"
      :placeholder="palceholder"
      label="automated"
      track-by="automated"
      :allow-empty="false"
      :disabled="(workload.selected_tag.tag) ? isDisabled : currentTag.tag != workload.available_tags[0].tag"
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
  protected isDisabled: Boolean = true;

  protected possibleOptions: any = [
    { automated: true },
    { automated: false },
  ];

  public updated() {
    this.refreshSelectedValue();
  }

  public refreshSelectedValue() {
    const {selected_tag, available_tags} = this.workload;
    this.isDisabled = (selected_tag.tag && selected_tag.tag != available_tags[0].tag) ? true : false;
  }

  get options() {
    return this.possibleOptions;
  }

  get palceholder() {
    const {current_tag, selected_tag, available_tags} = this.workload;
    let placeholder = '';

    if (current_tag && !selected_tag){
      placeholder = current_tag.automated ? 'true' : 'false';
    } else {
      placeholder = (selected_tag.tag && selected_tag.tag != available_tags[0].tag) ? 'false' : (available_tags[0].automated) ? 'true' : 'false';
    }

    return placeholder;
  }

  public valueChanged() {
    this.$emit("input", this.workload, this.value);
  }
}
</script>
