<template>
  <a-layout-content class="layout-content">
    <section class="tabs-view">
      <a-tabs
        :active-key="activeKey"
        :tabBarGutter="5"
        hide-add
        type="editable-card"
        class="tabs"
        @change="changePage"
        @edit="editTabItem"
        size="small"
      >
        <a-tab-pane v-for="tabItem in tabsViewStore.getTabsList" :key="tabItem.fullPath">
          <template #tab>
            <TabsOperator :ref="(ins) => (itemRefs[tabItem.fullPath] = ins)" :tab-item="tabItem" />
          </template>
        </a-tab-pane>
      </a-tabs>
    </section>
    <section class="tabs-view-content" :style="{ overflow }">
      <div class="tabs-view-pane">
        <router-view v-slot="{ Component }">
          <template v-if="Component">
            <Suspense>
              <Transition
                name="fade-slide"
                mode="out-in"
                appear
                @before-leave="overflow = 'hidden'"
                @after-leave="overflow = 'auto'"
              >
                <keep-alive :include="keepAliveComponents">
                  <component :is="keepAliveWrap($route.name, Component)" :key="$route.fullPath" />
                </keep-alive>
              </Transition>
              <template #fallback> 正在加载... </template>
            </Suspense>
          </template>
        </router-view>
        <slot></slot>
      </div>
    </section>
  </a-layout-content>
</template>

<script setup>
import { computed, ref, h } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import { useTabsViewStore } from '@/stores/tabsView';
import { useKeepAliveStore } from '@/stores/keepAlive';
import TabsOperator from './tabs-operator.vue';

const route = useRoute();
const router = useRouter();
const tabsViewStore = useTabsViewStore();
const keepAliveStore = useKeepAliveStore();

const itemRefs = {};

// 解决路由切换动画出现滚动条闪烁问题
const overflow = ref('auto');
const activeKey = computed(() => {
  const current = tabsViewStore.getCurrentTab || {};
  return current.fullPath || "";
});

// 缓存的路由组件列表
const keepAliveComponents = computed(() => keepAliveStore.list);

const wrapperMap = new Map()
const keepAliveWrap = (name, component) => {
  let wrapper
  const wrapperName = name
  if (wrapperMap.has(wrapperName)) {
    wrapper = wrapperMap.get(wrapperName)
  } else {
    wrapper = {
      name: wrapperName,
      render() {
        return component
      },
    }
    wrapperMap.set(wrapperName, wrapper)
  }
  return h(wrapper)
}

// tabs 编辑（remove || add）
const editTabItem = (targetKey, action) => {
  if (action == 'remove') {
    itemRefs[targetKey].removeTab();
  }
};

// 切换页面
const changePage = (key) => {
  Object.is(route.fullPath, key) || router.push(key);
};
</script>

<style lang="scss" scoped>
.layout-content {
  flex: none;
  padding: 4px 4px 0px 4px;

  .tabs-view {
    :deep(.ant-tabs) {
      .ant-tabs-nav {
        --un-bg-opacity: 1;
        background-color: rgb(255 255 255 / var(--un-bg-opacity));
        margin: 0;
        padding: 4px;
        user-select: none;
        height: 44px;

        .ant-tabs-tab {
          border-radius: 4px;
          border: 1px solid rgba(5, 5, 5, 0.06);
          background: rgba(0, 0, 0, 0);

          &.ant-tabs-tab-active {
            border-color: #1677ff;

            .ant-tabs-tab-remove {
              .anticon-close {
                color: #1677ff;
              }
            }
          }
        }
      }
    }
  }

  .tabs-view-content {
    padding-top: 4px;
    .tabs-view-pane{
      height: calc(100vh - var(--app-header-height) - 48px - 4px);
      overflow: hidden;
      overflow-y: auto;
      >*{
        --un-bg-opacity: 1;
        background-color: rgb(255 255 255 / var(--un-bg-opacity));
        min-height: calc(100vh - var(--app-header-height) - var(--app-footer-height) - 48px - 4px);
      }
    }
  }
}
</style>
