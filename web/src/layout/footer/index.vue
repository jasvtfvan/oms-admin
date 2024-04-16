<template>
  <a-layout-footer class="layout-footer">
    <div>{{ footerInfo }}</div>
  </a-layout-footer>
</template>

<script setup>
import { ref, computed } from 'vue'
import setting from '@/setting.js'

const copyright = ref(setting.websiteInfo.copyright)
const forTheRecord = ref(setting.websiteInfo.forTheRecord)
const footerInfo = computed(() => `${copyright.value} | ${forTheRecord.value}`)
</script>

<style lang="scss" scoped>
.layout-footer {
  text-align: center;
  font-size: 12px;
  color: rgba(0, 0, 0, 0.45);
}
</style>