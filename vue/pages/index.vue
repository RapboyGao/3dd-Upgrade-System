<template>
  <div>
    <NuxtWelcome />
  </div>
</template>

<script setup lang="ts">
useSeoMeta({
  title: 'Hello World',
  applicationName: 'Hello World',
});
</script>

<style lang="scss">
.bg-black {
  background-color: black;
  color: #cccccc;
}
</style>
