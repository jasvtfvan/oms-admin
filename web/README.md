# web

This template should help get you started developing with Vue 3 in Vite.

## Recommended IDE Setup

[VSCode](https://code.visualstudio.com/) + [Volar](https://marketplace.visualstudio.com/items?itemName=Vue.volar) (and disable Vetur).

## Customize configuration

See [Vite Configuration Reference](https://vitejs.dev/config/).

## Project Setup

```sh
npm install
```

### Compile and Hot-Reload for Development

```sh
npm run dev
```

### Compile and Minify for Production

```sh
npm run build
```

### Lint with [ESLint](https://eslint.org/)

```sh
npm run lint
```

# 测试

### build测试
1. 安装**Live Server**工具
2. 修改`vite.config.js`
```conf
base: './',
```
3. 修改`src/router/index.js`
```js
createWebHistory --替换成--> createWebHashHistory
```
4. 执行build
```sh
npm run build:staging
```
5. 打开dist_staging/执行右键index.html
**Open With Live Server**
6. 还原代码
`vite.config.js`和`src/router/index.js`
