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
import { decryptPwd, decryptSecret } from '@/utils/cryptoLoginSecret'
import { rsaEncryptOAEP } from '@/utils/rsaEncrypt'
import { useUserStore } from '@/stores/user'

const userStore = useUserStore()
const username = userStore.userProfile.username

const collapsed = ref(false)
const theme = ref('light')

const judgePwd = async () => {
  const password = decryptPwd()
  const enSecret = userStore.encryptedSecret
  if (!enSecret) {
    console.warn('userStore.encryptedSecret未取到')
  }
  const deSecret = decryptSecret(enSecret)
  const secret = await rsaEncryptOAEP(JSON.stringify({ username, password }))
  const secret1 = await rsaEncryptOAEP(JSON.stringify({ username, password }))
  console.log(secret == secret1)
  if (secret != deSecret) {
    // TODO 可以通过调用后台，验证两次结果是否一致
    console.warn('缓存密码被篡改')
  }
}
judgePwd()
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
