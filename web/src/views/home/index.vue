<template>
  <div>
    <section>
      <a-input v-model:value="value" placeholder="BasicUsage" />
      <img src="@/assets/images/empty.svg" />
    </section>
  </div>
</template>

<script setup>
import { watch, ref } from 'vue';

const value = ref('');
</script>

<style lang="scss" scoped>
</style>
