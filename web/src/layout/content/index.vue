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
}
</style>