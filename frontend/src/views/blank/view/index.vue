<template>
  <div class="viewer">
    <MdViewer />
  </div>
</template>


<script setup>
import MdViewer from '../../../components/MdViewer.vue';
</script>


<style scoped>
.viewer {
  height: calc(100% - 50px);
  overflow-y: auto;
}
</style>