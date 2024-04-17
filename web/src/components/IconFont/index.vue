<template>
  <render></render>
</template>

<script setup>
import { h, ref, defineProps } from 'vue'

const props = defineProps({
  iconName: {
    type: String,
    default: ''
  }
})

const render = () =>
  h('<div></div>')
</script>

