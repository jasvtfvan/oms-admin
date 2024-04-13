<template>
  <a-config-provider :locale="locale">
    <loading>
      <router-view />
    </loading>
  </a-config-provider>
</template>

<script setup>
import { ref } from 'vue'
import zhCN from 'ant-design-vue/es/locale/zh_CN'
import dayjs from 'dayjs'
import 'dayjs/locale/zh-cn'
import Loading from '@/components/Loading/index.vue'

dayjs.locale('zh-cn')
const locale = ref(zhCN)
</script>

<style scoped>
</style>
