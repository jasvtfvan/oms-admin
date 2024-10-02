<template>
  <div>user: {{ now }}</div>
</template>

<script setup>
import { ref } from 'vue';

const now = ref(new Date().getTime());
</script>
