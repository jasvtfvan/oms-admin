<template>
  <article class="layout-container">
    <a-layout class="layout-main">
      <Sider :theme="theme" :collapsed="collapsed"></Sider>
      <!-- right -->
      <a-layout class="layout-right">
        <Header></Header>
        <!-- content -->
        <Content></Content>
        <!-- /content -->
        <Footer></Footer>
      </a-layout>
      <!-- /right -->
    </a-layout>
    <ChangePwd></ChangePwd>
  </article>
</template>

<script setup>
import { onMounted, ref } from 'vue'
import Sider from './sider/index.vue'
import Header from './header/index.vue'
import Content from './content/index.vue'
import Footer from './footer/index.vue'
import ChangePwd from './ChangePwd.vue'
import { useUserStore } from '@/stores/user'

const userStore = useUserStore()
const { menuNames } = userStore

const collapsed = ref(false) // 菜单折叠
const theme = ref('light') // 主题

onMounted(() => {
  userStore.AddDynamicRoutes(menuNames)
})
</script>

<style lang="scss" scoped>
.layout-container {
  min-height: 100vh;
  .layout-main {
    height: 100%;
    .layout-right {
    }
  }
}
</style>
