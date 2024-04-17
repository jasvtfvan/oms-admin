import router from '@/router'
import { dynamicModules, rootLayout, defaultModules } from '@/router/layout'

// 递归的子方法（默认admin都选中，除非特殊notAdminDefault）
const _getAdminMenuNames = (modules = []) => {
  return modules.reduce((prev, curt) => {
    if (curt.children && curt.children.length) {
      return prev.concat(_getAdminMenuNames(curt.children))
    } else {
      if (curt.meta && curt.meta.notAdminDefault) { // 非admin默认选中的
        return prev
      } else {
        return prev.concat(curt.name)
      }
    }
  }, []);
}

// admin（rootAdmin或者普通admin）返回所有菜单名（不包含/home等默认路由）
export const getAdminMenuNames = () => {
  return _getAdminMenuNames(dynamicModules);
}

// 将menus进行排序
const _sortMenus = (menus) => {
  return menus
    .filter((m) => {
      meta = m.meta || {};
      const isShow = !meta.hideInMenu; // 选出不隐藏的
      children = m.children || [];
      if (isShow && children.length) {
        m.children = _sortMenus(children);
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

// 递归的子方法，menuNames不包含则隐藏，虽然包含如何本身隐藏的继续隐藏，通过跳转时判断权限
const _getLayoutMenus = (modules = [], menuNames = []) => {
  const menus = []
  modules.forEach(module => {
    const target = {
      path: module.path || '/', // 菜单路径
      name: module.name || '', // 菜单名称
      meta: module.meta || {}, // 菜单属性
    }
    if (!target.name && !target.meta.hideInMenu) target.meta.hideInMenu = true
    if (!menuNames.includes(target.meta.name)) target.meta.hideInMenu = true

    if (module.children && module.children.length) {
      target.children = _getLayoutMenus(module.children, menuNames)
    }

    menus.push(target)
  })
  return menus
}

// 根据menuNames（不包含/home等默认路由）获取菜单
export const getLayoutMenus = (menuNames = []) => {
  const modules = [...defaultModules, ...dynamicModules]
  const menus = _getLayoutMenus(modules, menuNames)
  return _sortMenus(menus)
}

// 把menu转成路由，menuNames不包含则隐藏，虽然包含如何本身隐藏的继续隐藏，通过跳转时判断权限
const _transformMenuToRoutes = (modules = [], menuNames = []) => {
  modules.forEach(route => {
    if (!route.meta) route.meta = {}
    if (!route.meta.name) route.meta.name = ''
    if (!route.name && !route.meta.hideInMenu) route.meta.hideInMenu = true
    if (!menuNames.includes(route.meta.name)) route.meta.hideInMenu = true
    if (route.children && route.children.length) {
      _transformMenuToRoutes(route.children, menuNames);
    }
  })
  return modules
}

// 根据menuNames（不包含/home等默认路由）添加动态路由
export const addDynamicRoutes = (menuNames) => {
  if (menuNames && menuNames.length) { // 获取到菜单名后，才需要添加
    const transRoutes = _transformMenuToRoutes(dynamicModules, menuNames)
    // 根据menuNames把dynamicModules添加到路由中
    rootLayout.children.push(transRoutes)
    router.addRoute(rootLayout)
  }
};
