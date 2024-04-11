import { fileURLToPath, URL } from 'node:url'

import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import vueJsx from '@vitejs/plugin-vue-jsx'

// https://vitejs.dev/config/
export default () =>
  defineConfig({
    plugins: [vue(), vueJsx()],
    base: '/oms-admin/',
    server: {
      //本地服务器主机名 配置后可以使用本地网络访问
      host: '0.0.0.0',
      //指定启动端口号
      port: 8080,
      //设为 true 时若端口已被占用则会直接退出，而不是尝试下一个可用端口
      strictPort: false,
      //服务器启动时自动在浏览器中打开应用程序,当此值为字符串时，会被用作 URL 的路径名
      open: true,
      proxy: {
        '/v1': {
          target: ' http://localhost:8888',
          changeOrigin: true,
          // rewrite: (path) => path.replace(/^\/api/, ''),
        }
      }
    },
    resolve: {
      alias: {
        '@': fileURLToPath(new URL('./src', import.meta.url))
      }
    },
  })
