import '@/assets/styles/index.scss'

import { createApp } from 'vue'
// 状态管理
import { createPinia } from 'pinia'
import App from './App.vue'
import router from './router'
// 引入svg-icons
import 'virtual:svg-icons-register'

// 导入加载进度条，防止首屏加载时间过长
import Nprogress from 'nprogress'
import 'nprogress/nprogress.css'
Nprogress.configure({ showSpinner: false, ease: 'ease', speed: 500 })
Nprogress.start()

const app = createApp(App)

app.use(createPinia())
app.use(router)

app.mount('#app')
