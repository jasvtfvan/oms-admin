<template>
  <a-spin :spinning="loading" :tip="tip" :indicator="indicator">
    <slot></slot>
  </a-spin>
</template>

<script setup>
import { Loading3QuartersOutlined } from '@ant-design/icons-vue'
import { h, ref, watchEffect } from 'vue'
import { useLoading } from '@/hooks/useLoading'

const { getLoading, getLoadingTip } = useLoading()
const loading = ref(false)
const tip = ref('')
watchEffect(() => {
  loading.value = getLoading.value
  tip.value = getLoadingTip.value
})

const indicator = h(Loading3QuartersOutlined, {
  style: {
    fontSize: '24px'
  },
  spin: true
})
</script>
