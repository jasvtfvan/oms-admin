<template>
  <div>
    <div>
      <a-button>占位</a-button>
    </div>
    <div>
      <a-button danger @click="onLogout">退出</a-button>
    </div>
    <div>
      <a-button>占位</a-button>
    </div>
    <div>
      <a-button>占位</a-button>
    </div>
    <div>
      <a-button type="primary" @click="onSubmit">登录</a-button>
    </div>
  </div>
</template>

<script setup>
import { onMounted } from 'vue'
import { useUserStore } from '@/stores/user'
const userStore = useUserStore()

onMounted(() => {
  console.log('onMounted')
})

const onSubmit = async () => {
  await userStore.Login({})
}
const onLogout = async () => {
  await userStore.Logout()
}
</script>

<style lang="scss" scoped>
.svg-wrap {
  width: 30px;
  height: 30px;
  position: relative;
  color: blue;
}
</style>
