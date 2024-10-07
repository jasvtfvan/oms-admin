<template>
  <a-breadcrumb>
    <template v-for="(routeItem, routeIndex) in menus" :key="routeItem.name">
      <a-breadcrumb-item>
        {{ routeItem.meta.title || "" }}
      </a-breadcrumb-item>
    </template>
  </a-breadcrumb>
</template>

<script setup>
import { computed } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import { useUserStore } from '@/stores/user';

defineOptions({
  name: 'LayoutBreadcrumb',
});

const router = useRouter();
const route = useRoute();
const userStore = useUserStore();

const menus = computed(() => {
  return route.matched.filter(item => item.path && item.path != '/');
});

const getSelectKeys = (routeIndex) => {
  return [menus.value[routeIndex + 1]?.name];
};
</script>

<style lang="scss" scoped>

</style>
