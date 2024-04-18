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
import { rsaEncryptOAEP } from '@/utils/rsaEncrypt'
import { aesEncryptCBC, aesDecryptCBC } from '@/utils/aesCrypto'

console.log(aesEncryptCBC('hello world ga ga'))
console.log(aesDecryptCBC('012f7bf73e5b190d135f878514dbab2a44f6db23dd0e6f1b7319bbcd149ab94f'))

onMounted(async () => {
  console.log('onMounted')
  // generateKeyPair()
  const user = '{"username":"oms_admin","password":"Oms123Admin456"}'
  const res = await rsaEncryptOAEP(user)
  console.log(res)
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
