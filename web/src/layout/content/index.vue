<template>
  <a-layout-content class="layout-content">
    <router-view v-slot="{ Component }">
      <component :is="Component" />
    </router-view>
  </a-layout-content>
</template>

<style lang="scss" scoped>
.layout-content {
  flex: none;
  height: calc(100vh - var(--app-header-height) - var(--app-footer-height));
  overflow: hidden;
  overflow-y: auto;
}
</style>