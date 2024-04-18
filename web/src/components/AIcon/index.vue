<template>
  <component :is="iconComponent" />
</template>

<script setup>
import { defineComponent, h } from 'vue'
import * as Icons from '@ant-design/icons-vue'

defineOptions({
  name: 'AIcon'
})

const props = defineProps({
  name: {
    type: String,
    required: true
  }
})

const iconComponent = defineComponent({
  render() {
    const IconComponent = Icons[props.name]
    return IconComponent ? h(IconComponent) : null
  }
})
</script>
