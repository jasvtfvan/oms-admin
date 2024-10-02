<template>
  <div>
    keepAlive
    <a-icon name="ReloadOutlined" /> 角色页面
    {{ now }}
    <a-button type="primary" html-type="submit" size="large" block>
      登录
    </a-button>
  </div>
</template>

<script setup>
import { ref } from 'vue';

const now = ref(new Date().getTime());
</script>
