<template>
  <div>
    <home-outlined />

    <div class="svg-wrap">
      <svg-icon name="logo" color="red"></svg-icon>
    </div>

    <a-steps
      :current="1"
      :items="[
        {
          title: 'Finished',
          description
        },
        {
          title: 'In Progress',
          description,
          subTitle: 'Left 00:00:08'
        },
        {
          title: 'Waiting',
          description
        }
      ]"
    ></a-steps>

    <a-button type="primary">你好</a-button>
  </div>
</template>

<script setup>
import { onMounted } from 'vue'
import { HomeOutlined } from '@ant-design/icons-vue'

onMounted(() => {
  console.log('onMounted')
})

const description = 'This is a description.'
</script>

<style lang="scss" scoped>
.svg-wrap {
  width: 30px;
  height: 30px;
  position: relative;
  color: blue;
}
</style>
