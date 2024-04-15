<template>
  <article class="layout-container">
    <a-layout class="layout-main">
      <a-layout-sider class="layout-left">Sider</a-layout-sider>
      <!-- right -->
      <a-layout class="layout-right">
        <a-layout-header>Header</a-layout-header>
        <!-- content -->
        <a-layout-content class="layout-content">
          <router-view v-slot="{ Component }">
            <component :is="Component" />
          </router-view>
        </a-layout-content>
        <!-- /content -->
        <a-layout-footer class="layout-footer">
          <div>{{ footerInfo }}</div>
        </a-layout-footer>
      </a-layout>
      <!-- /right -->
    </a-layout>
  </article>
</template>

<script setup>
import { ref, computed } from 'vue'
import setting from '@/setting.js'
const copyright = ref(setting.websiteInfo.copyright)
const forTheRecord = ref(setting.websiteInfo.forTheRecord)
const footerInfo = computed(() => `${copyright.value} | ${forTheRecord.value}`)
</script>

<style lang="scss" scoped>
.layout-container {
  min-height: 100vh;
  .layout-main {
    height: 100%;
    .layout-right {
      .layout-content {
        flex: none;
      }
      .layout-footer {
        text-align: center;
        font-size: 12px;
        color: rgba(0, 0, 0, 0.45);
      }
    }
  }
}
</style>
