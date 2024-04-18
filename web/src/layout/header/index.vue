<template>
  <a-layout-header class="layout-header">
    <div class="container">
      <div>header</div>
      <div>
        <a-button type="primary" size="small" block @click="doSubmit"> 刷新权限 </a-button>
      </div>
    </div>
  </a-layout-header>
</template>

<script setup>
import { nextTick } from 'vue';
import { message } from 'ant-design-vue'
import { useUserStore } from '@/stores/user'
const userStore = useUserStore()

const doSubmit = async () => {
  message.loading('登录中...', 0)
  try {
    await userStore.RefreshAuth()
  } catch (error) {
    nextTick(() => {
      message.error(error.msg, 2, () => {})
    })
  } finally {
    message.destroy()
  }
}
</script>

<style lang="scss" scoped>
.layout-header {
  background: rgba(255, 255, 255, 0.85);
  .container {
    display: flex;
    justify-content: space-between;
    align-items: center;
  }
}
</style>
