<template>
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
  >
  </a-steps>
</template>
<script setup>
const description = 'This is a description.'
</script>

