<template>
  <section class="menu-container">
    <a-menu :theme="props.theme" :collapsed="props.collapsed" collapsible> 菜单 </a-menu>
  </section>
</template>

<script setup>
const props = defineProps({
  collapsed: {
    type: Boolean,
    default: false
  },
  theme: {
    type: String
  }
})
</script>

<style lang="scss" scoped>
.menu-container {
  width: 100%;
  overflow: auto;
  &::-webkit-scrollbar {
    width: 0;
    height: 0;
  }
}
</style>
