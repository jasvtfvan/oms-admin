<template>
  <section>
    keepAlive
    <a-steps :current="1" :items="[
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
        ]">
    </a-steps>
    {{ now }}
  </section>
</template>

<script setup>
import { ref } from 'vue';

const description = 'This is a description.'
const now = ref(new Date().getTime());
</script>
