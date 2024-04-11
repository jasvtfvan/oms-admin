import { createRouter, createWebHistory } from 'vue-router'
import Index from '@/views/index.vue'

const moduleDir = import.meta.glob('./modules/*.js')
const modules = [];
for (const path in moduleDir) {
  const module = await moduleDir[path](); // TODO 需要改造成动态router，否则await在build时报错
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
  next()
})

router.onError((error) => {
  console.error('路由错误:', error)
})


export default router
