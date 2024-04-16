<template>
  <a-layout-sider
    class="layout-left"
    :theme="props.theme"
    collapsible
    :width="siderWidth"
    v-model:collapsed="siderCollapsed"
  >
    <section class="logo-wrap">
      <div class="svg-wrap">
        <svg-icon name="logo" color="red"></svg-icon>
      </div>
      <h2 v-show="!siderCollapsed" class="title">{{ nickName }}</h2>
    </section>
    <SideMenu :theme="props.theme" :collapsed="siderCollapsed" />
  </a-layout-sider>
</template>

<script setup>
import { ref, computed, watchEffect } from 'vue'
import setting from '@/setting.js'
import SideMenu from './menu/index.vue'

const props = defineProps({
  collapsed: {
    type: Boolean,
    default: false
  },
  theme: {
    type: String
  }
})

const siderCollapsed = ref(false)
watchEffect(() => {
  siderCollapsed.value = props.collapsed
})
const siderWidth = computed(() => (siderCollapsed.value ? 80 : 220))

const nickName = ref(setting.websiteInfo.nickName)
</script>

<style lang="scss" scoped>
.layout-left {
  .logo-wrap {
    display: flex;
    align-items: center;
    overflow: hidden;
    white-space: nowrap;
    height: var(--app-header-height);
    line-height: var(--app-header-height);
    padding-left: 24px;
    .svg-wrap {
      width: 32px;
      height: 32px;
      margin-right: 8px;
    }
    .title {
      margin-bottom: 0;
      font-size: 20px;
      line-height: 28px;
      color: var(--app-primary-color);
    }
  }
}
</style>
