<template>
  <div>
    <div>
      <a-button type="primary">占位</a-button>
    </div>
    <div>
      <a-button type="primary">占位</a-button>
    </div>
    <div>
      <a-button type="primary">占位</a-button>
    </div>
    <div>
      <a-button type="primary">占位</a-button>
    </div>
    <div>
      <a-button type="primary" @click="onSubmit">登录</a-button>
    </div>
  </div>
</template>

<script setup>
import { onMounted } from 'vue'
import { login } from '@/api/common/user'

onMounted(() => {
  console.log('onMounted')
})

const onSubmit = () => {
  login({ username: 'nihao' })
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
