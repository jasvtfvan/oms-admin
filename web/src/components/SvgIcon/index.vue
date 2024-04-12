<template>
  <svg :class="svgClass" aria-hidden="true" v-bind="$attrs">
    <use :xlink:href="symbolId" :fill="color" />
  </svg>
</template>

<script setup name="SvgIcon">
import { computed } from 'vue'

const props = defineProps({
  prefix: {
    type: String,
    default: 'icon'
  },
  color: {
    type: String,
    default: '',
  },
  name: {
    type: String,
    required: true
  },
  className: {
    type: String,
    default: ''
  }
})

const symbolId = computed(() => `#${props.prefix}-${props.name}`)

const svgClass = computed(() => {
  if (props.className) {
    return `svg-icon ${props.className}`
  }
  return 'svg-icon'
})
</script>

<style scoped>
.svg-icon {
  width: 100%;
  height: 100%;
  fill: currentColor;
  overflow: hidden;
  display: block;
}
</style>
