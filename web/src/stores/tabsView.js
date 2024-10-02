import { computed, ref, watch } from 'vue';
import { defineStore } from 'pinia';
import { useKeepAliveStore } from './keepAlive';
import router from '@/router';
import { useRoute } from 'vue-router';
import { REDIRECT_NAME, LOGIN_NAME, PAGE_NOT_FOUND_NAME } from '@/router/constant';

// 不需要出现在标签页中的路由
export const routeExcludes = [REDIRECT_NAME, LOGIN_NAME, PAGE_NOT_FOUND_NAME];

export const useTabsViewStore = defineStore('tabs-view', () => {
  const tabsList = ref([]);
  const currentRoute = useRoute();

  const getTabsList = computed(() => {
    return tabsList.value.filter((item) => {
      return item && !_isInRouteExcludes(item) && router.hasRoute(item.name);
    });
  });

  const getCurrentTab = computed(() => {
    return tabsList.value.find((item) => {
      return item && !_isInRouteExcludes(item) && item.fullPath === currentRoute.fullPath;
    });
  });

  /** 给定的路由是否在排除名单里面 */
  const _isInRouteExcludes = (route) => {
    return (route.meta && route.meta.hideInTabs) || routeExcludes.some((n) => n === route.name);
  };

  const _getRawRoute = (route) => {
    return {
      ...route,
      matched: route.matched.map((item) => {
        const { meta, path, name } = item;
        return { meta, path, name };
      }),
    };
  };

  const _addTabs = (route) => {
    if (_isInRouteExcludes(route)) {
      return false;
    }
    const isExists = tabsList.value.some((item) => item.fullPath == route.fullPath);
    if (!isExists) {
      tabsList.value.push(_getRawRoute(route));
    }
    return true;
  };

  /** 将已关闭的标签页的组件从keep-alive中移除 */
  const _delCompFromClosedTabs = (closedTabs) => {
    const keepAliveStore = useKeepAliveStore();
    const routes = router.getRoutes();
    const compNames = closedTabs.reduce((prev, curr) => {
      if (curr.name && router.hasRoute(curr.name)) {
        const componentName = routes.find((n) => n.name === curr.name).name;
        componentName && prev.push(componentName);
      }
      return prev;
    }, []);
    keepAliveStore.remove(compNames);
    keepAliveStore.add("home");
  };
  /** 关闭左侧 */
  const CloseLeftTabs = (route) => {
    const index = tabsList.value.findIndex((item) => item.fullPath == route.fullPath);
    _delCompFromClosedTabs(tabsList.value.splice(0, index));
  };
  /** 关闭右侧 */
  const CloseRightTabs = (route) => {
    const index = tabsList.value.findIndex((item) => item.fullPath == route.fullPath);
    _delCompFromClosedTabs(tabsList.value.splice(index + 1));
  };
  /** 关闭其他 */
  const CloseOtherTabs = (route) => {
    const targetIndex = tabsList.value.findIndex((item) => item.fullPath === route.fullPath);
    if (targetIndex !== -1) {
      const current = tabsList.value.splice(targetIndex, 1);
      _delCompFromClosedTabs(tabsList.value);
      tabsList.value = current;
    }
  };
  /** 关闭当前页 */
  const CloseCurrentTab = (route) => {
    const index = tabsList.value.findIndex((item) => item.fullPath == route.fullPath);
    const isDelCurrentTab = Object.is(getCurrentTab.value, tabsList.value[index]);
    _delCompFromClosedTabs(tabsList.value.splice(index, 1));
    // 如果关闭的tab就是当前激活的tab，则重定向页面
    if (isDelCurrentTab) {
      const currentRoute = tabsList.value[Math.max(0, tabsList.value.length - 1)];
      router.push(currentRoute);
    }
  };
  /** 关闭全部 */
  const CloseAllTabs = () => {
    _delCompFromClosedTabs(tabsList.value);
    const home = tabsList.value.find(item => item.name == 'home');
    tabsList.value = [home];
  };

  // 更新tab标题
  const UpdateTabTitle = (title) => {
    const currentRoute = router.currentRoute.value;
    const upTarget = tabsList.value.find((item) => item.fullPath === currentRoute.fullPath);
    if (upTarget) {
      upTarget.meta.title = title;
    }
  };

  watch(
    () => currentRoute.fullPath,
    () => {
      _addTabs(currentRoute);
    },
    { immediate: true },
  );

  return {
    getTabsList,
    getCurrentTab,
    CloseAllTabs,
    CloseCurrentTab,
    CloseLeftTabs,
    CloseOtherTabs,
    CloseRightTabs,
    UpdateTabTitle,
  }
});
