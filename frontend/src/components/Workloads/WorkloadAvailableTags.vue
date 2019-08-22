<template>
  <div>
    <multiselect
      v-model="value"
      :options="options"
      placeholder="Select one"
      label="tag"
      track-by="tag"
      @input="valueChanged"
    >
      <template slot="singleLabel" slot-scope="{ option }">
        <strong>{{ option.tag }}</strong> -
        <strong>{{ moment(option.date).fromNow() }}</strong>
      </template>
      <template slot="option" slot-scope="props">
        <div class="option__desc">
          <span class="option__tag">{{ props.option.tag }}</span>
        </div>
      </template>
    </multiselect>
  </div>
</template>

<style lang="scss">
@import "../../assets/scss/include";
</style>

<script lang="ts">
import { Component, Vue, Prop } from "vue-property-decorator";
import Multiselect from "vue-multiselect";
import moment from "moment";

@Component({
  components: {
    Multiselect
  }
})
export default class WorkloadAvailableTags extends Vue {
  @Prop() protected optionsProp: any;
  @Prop() protected workload: any;

  protected value: any = null;

  protected moment = moment;

  get options() {
    return this.optionsProp.map((option: any) => {
      option.$isDisabled = option.current;
      return option;
    });
  }

  public valueChanged() {
    this.$emit("input", this.workload, this.value);
  }
}
</script>