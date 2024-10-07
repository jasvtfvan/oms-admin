<template>
  <section>
    keepAlive
    <a-steps :current="1" :items="
      [
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
      ]
    ">
    </a-steps>
    <a-input v-model:value="value" placeholder="BasicUsage" />
  </section>
</template>

<script setup>
import { watch, ref } from 'vue';

const value = ref('');

const description = 'This is a description.';
</script>
