<template>
  <a-layout-header class="layout-header">
    <div class="container">
      <div>header</div>
      <div class="btn-wrap">
        <a-button type="error" size="small" block @click="confirmLogout"> 退出登录 </a-button>
        <a-button type="primary" size="small" block @click="refreshAuth"> 刷新权限 </a-button>
      </div>
    </div>
  </a-layout-header>
</template>

<script setup>
import { nextTick } from 'vue'
import { Modal, message } from 'ant-design-vue'
import { useUserStore } from '@/stores/user'
const userStore = useUserStore()

// 是否退出
const confirmLogout = async () => {
  Modal.confirm({
    title: '确定退出系统？',
    okText: '确定',
    cancelText: '取消',
    onOk: () => {
      doLogout()
    }
  })
}

// 退出
const doLogout = async () => {
  await userStore.Logout()
  window.location.reload()
}

// 刷新权限
const refreshAuth = async () => {
  try {
    message.loading('登录中...', 0)
    await userStore.RefreshAuth()
  } catch (error) {
    console.warn(error)
    nextTick(() => {
      Modal.confirm({
        title: '刷新失败',
        content: '可选择【继续使用】或【退出系统】',
        okText: '退出系统',
        cancelText: '继续使用',
        onOk: () => {
          doLogout()
        }
      })
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

    .btn-wrap {
      display: flex;
    }
  }
}
</style>
