import router from '@/router'
import { layoutModules, rootLayout } from '@/router/layout'
import globRoutes from '@/router/unLayout'

const sortRoutes = (menus) => {
  return menus
    .filter((m) => {
      meta = m.meta || {};
      const isShow = !meta.hideInMenu; // 选出不隐藏的
      children = m.children || [];
      if (isShow && children.length) {
        m.children = sortRoutes(children);
      }
      return isShow;
    })
    .sort((a, b) => {
      const meta1 = a.meta || {};
      const meta2 = b.meta || {};
      const sortMenu1 = Number(meta1.sortMenu) || 0;
      const sortMenu2 = Number(meta2.sortMenu) || 0;
      return sortMenu1 - sortMenu2;
    });
};

// 把menu转成路由
const transformMenuToRoutes = (menus) => {
  // TODO 通过menus过滤layoutModules
  return layoutModules
}

export const addDynamicRoutes = (menus) => {
  const routes = [
    ...globRoutes,
    rootLayout,
  ]
  if (menus && menus.length) {
    // 根据menus把layoutModules用到的添加到路由中
    const transRoutes = transformMenuToRoutes(menus)
    const sortedRoutes = sortRoutes(transRoutes)
    rootLayout.children.push(sortedRoutes)
    router.addRoute(rootLayout);
  } else {
    router.addRoute(sortRoutes(routes))
  }
};
