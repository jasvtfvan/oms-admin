<template>
  <a-sub-menu v-if="showSubMenu(item)" :key="`${item.name}_1`">
    <template #icon>
      <a-icon :name="item.meta.icon || 'LayoutOutlined'" />
    </template>
    <template #title>
      <span>{{ item.meta.title }}</span>
    </template>
    <template v-for="child in item.children || []" :key="child.name">
      <MySubMenu :item="child" />
    </template>
  </a-sub-menu>
  <a-menu-item v-else :key="`${item.name}_2`" @click="linkTo(item)">
    <a-icon :name="item.meta.icon || 'FileOutlined'" />
    <span>{{ item.meta.title }}</span>
  </a-menu-item>
</template>

<script setup>
import { useRouter } from 'vue-router';

defineOptions({
  name: 'MySubMenu'
})
defineProps({
  item: {
    type: Object,
    default: () => ({})
  }
})
const router = useRouter()

const showSubMenu = (item) => {
  return (item.children && item.children.length) || item.meta.hideChildrenInMenu
}
const linkTo = (item) => {
  router.push(item.path)
}
</script>
