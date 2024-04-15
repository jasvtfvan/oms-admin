import { createRouter, createWebHistory } from 'vue-router'
import Index from '@/views/index.vue'
import Nprogress from 'nprogress'

const moduleDir = import.meta.glob('./modules/*.js', { eager: true })
const modules = [];
for (const path in moduleDir) {
  const module = moduleDir[path];
  modules.push(...module.default);
}

const routes = [
  {
    path: '/:pathMatch(.*)*',
    name: '404',
    component: () => import('@/views/404.vue'),
    hidden: true
  },
  {
    path: '/',
    name: 'Index',
    component: Index,
  },
  ...modules,
]

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes,
})


router.beforeEach(async (to, from, next) => {
  // 可以跳转后，开启Nprogress，并且执行next
  Nprogress.start()
  next()
})

router.afterEach(async () => {
  Nprogress.done()
})

router.onError((error) => {
  console.error('路由错误:', error)
})


export default router
