import '@/assets/styles/index.scss'

import { createApp } from 'vue'
// 状态管理
import { createPinia } from 'pinia'
import App from './App.vue'
import router from './router'
// 引入svg-icons
import 'virtual:svg-icons-register'

const app = createApp(App)

app.use(createPinia())
app.use(router)

app.mount('#app')
