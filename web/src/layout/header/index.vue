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
import { useUserStore } from '@/stores/user'
const userStore = useUserStore()

const doSubmit = () => {
  userStore.GetUserProfile()
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
