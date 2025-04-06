<script lang="ts" setup></script>

<template>
  <div class="vertical-menu">
    <h3>Desktop Cheat Sheet</h3>
    <a href="#tabela1">Tabela 1</a>
    <a href="#tabela2">Tabela 2</a>
    <a href="#tabela3">Tabela 3</a>
  </div>
</template>

<style scoped>
.vertical-menu {
  width: 200px;
  height: 100vh;
  background-color: #f8f9fa;
  padding: 10px;
  position: fixed;
  overflow-x: hidden;
  box-shadow: 2px 0 5px rgba(0,0,0,0.1);
}

.vertical-menu a {
  display: block;
  padding: 10px;
  color: #333;
  text-decoration: none;
}

.vertical-menu a:hover {
  background-color: #ddd;
}
</style>