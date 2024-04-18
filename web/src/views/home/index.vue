<template>
  <div>
    <section>
      <img src="@/assets/images/empty.svg" />
      <img src="@/assets/images/empty.svg" />
    </section>
  </div>
</template>

<script setup>
import { onMounted } from 'vue'
import { encryptPwd, decryptPwd } from '@/utils/cryptoLoginSecret'
onMounted(async () => {
  encryptPwd('Oms123Admin456')
  console.log(decryptPwd())
})

</script>

<style lang="scss" scoped>
</style>
