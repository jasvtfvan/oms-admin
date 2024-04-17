<template>
  <a-sub-menu v-if="showSubMenu(item)" :key="`${item.name}_1`">
    <!-- <template v-if="item.meta.icon" #icon>
      <icon-font :type="`icon-${item.meta.icon}`" />
    </template> -->
    <template #title>
      <span>{{ item.meta.title }}</span>
    </template>
    <template v-for="child in item.children || []" :key="child.name">
      <MySubMenu :item="child" />
    </template>
  </a-sub-menu>
  <a-menu-item v-else :key="`${item.name}_2`">
    <iCon />
    <!-- <icon-font v-if="item.meta.icon" :type="`icon-${item.meta.icon}`" /> -->
    <span>{{ item.meta.title }}</span>
  </a-menu-item>
</template>

<script setup>
import iCon from '@/components/IconFont/index.vue'

defineOptions({
  name: 'MySubMenu'
})

defineProps({
  item: {
    type: Object,
    default: () => ({})
  }
})

const showSubMenu = (item) => {
  return (item.children && item.children.length) || item.meta.hideChildrenInMenu
}
</script>
