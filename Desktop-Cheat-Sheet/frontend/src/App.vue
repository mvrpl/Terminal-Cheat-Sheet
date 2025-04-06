<script lang="ts" setup>
import VerticalMenu from './components/Menu.vue';
import ContentArea from './components/ContentArea.vue';
</script>

<template>
  <div id="app" class="app-container">
    <VerticalMenu />
    <ContentArea />
  </div>
</template>

<style>
.app-container {
  height: 100vh;
}
</style>