
<template>
  <a-dropdown :trigger="[isExtra ? 'click' : 'contextmenu']">
    <a v-if="isExtra" class="ant-dropdown-link" @click.prevent>
      <a-icon name="DownOutlined" style="font-size: 20px;" />
    </a>
    <div v-else style="display: inline-block">
      {{ (tabItem.meta && tabItem.meta.title) || tabItem.name }}
    </div>
    <template #overlay>
      <a-menu style="user-select: none">
        <a-menu-item key="1" :disabled="activeKey !== tabItem.fullPath" @click="reloadPage">
          <a-icon name="ReloadOutlined" /> 重现加载
        </a-menu-item>
        <a-menu-item key="2" @click="removeTab">
          <a-icon name="CloseOutlined" /> 关闭标签页
        </a-menu-item>
        <a-menu-divider />
        <a-menu-item key="3" @click="closeLeft">
          <a-icon name="VerticalRightOutlined" /> 关闭左侧标签页
        </a-menu-item>
        <a-menu-item key="4" @click="closeRight">
          <a-icon name="VerticalLeftOutlined" /> 关闭右侧标签页
        </a-menu-item>
        <a-menu-divider />
        <a-menu-item key="5" @click="closeOther">
          <a-icon name="ColumnWidthOutlined" /> 关闭其他标签页
        </a-menu-item>
        <a-menu-item key="6" @click="closeAll">
          <a-icon name="MinusOutlined" /> 关闭全部标签页
        </a-menu-item>
      </a-menu>
    </template>
  </a-dropdown>
</template>

<script setup>
import { computed, unref } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import { message } from 'ant-design-vue';
import { useTabsViewStore } from '@/stores/tabsView';
import { REDIRECT_NAME } from '@/router/constant';

defineOptions({
  name: 'TabOperator',
});

const props = defineProps({
  tabItem: {
    type: Object,
    required: true,
  },
  isExtra: Boolean,
});

const route = useRoute();
const router = useRouter();
const tabsViewStore = useTabsViewStore();

const activeKey = computed(() => {
  const current = tabsViewStore.getCurrentTab || {};
  return current.fullPath || "";
});

/** 标签页列表 */
const tabsList = computed(() => tabsViewStore.getTabsList);

/** 目标路由是否等于当前路由 */
const isCurrentRoute = (route) => {
  return router.currentRoute.value.matched.some((item) => item.name === route.name);
};

/** 关闭当前页面 */
const removeTab = () => {
  if (tabsList.value.length === 1) {
    return message.warning('这已经是最后一页，不能再关闭了！');
  }
  tabsViewStore.CloseCurrentTab(props.tabItem);
};

/** 刷新页面 */
const reloadPage = () => {
  router.replace({
    name: REDIRECT_NAME,
    params: {
      path: unref(route).fullPath,
    },
  });
};

/** 关闭左侧 */
const closeLeft = () => {
  tabsViewStore.CloseLeftTabs(props.tabItem);
  !isCurrentRoute(props.tabItem) && router.replace(props.tabItem.fullPath);
};

/** 关闭右侧 */
const closeRight = () => {
  tabsViewStore.CloseRightTabs(props.tabItem);
  !isCurrentRoute(props.tabItem) && router.replace(props.tabItem.fullPath);
};

/** 关闭其他 */
const closeOther = () => {
  tabsViewStore.CloseOtherTabs(props.tabItem);
  !isCurrentRoute(props.tabItem) && router.replace(props.tabItem.fullPath);
};

/** 关闭全部 */
const closeAll = () => {
  tabsViewStore.CloseAllTabs();
  router.replace('/');
};

defineExpose({
  removeTab,
});
</script>

<style lang="scss" scoped>

</style>
