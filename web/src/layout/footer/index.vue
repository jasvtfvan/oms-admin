<template>
  <a-layout-footer class="layout-footer">
    <div>{{ copyright }}</div>
  </a-layout-footer>
</template>

<script setup>
import { ref } from 'vue'
import setting from '@/setting.js'

const copyright = ref(setting.websiteInfo.copyright)
</script>

<style lang="scss" scoped>
.layout-footer {
  height: var(--app-footer-height);
  padding: 0px;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 12px;
  color: rgba(0, 0, 0, 0.45);
}
</style>