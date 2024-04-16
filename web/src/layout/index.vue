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
  </article>
</template>

<script setup>
import { ref } from 'vue'
import Sider from './sider/index.vue'
import Header from './header/index.vue'
import Content from './content/index.vue'
import Footer from './footer/index.vue'

const collapsed = ref(false)
const theme = ref('light')
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
