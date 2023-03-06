<template>
<MainHeader />
  <ul v-infinite-scroll="load" class="infinite-list" style="overflow: auto">
    <li v-for="i in count" :key="i" class="infinite-list-item">{{ i }}</li>
  </ul>
</template>


<script setup>
import MainHeader from "../../components/MainHeader.vue"
import { ref } from 'vue'
const count = ref(5)
const load = () => {
  if (count.value < 15) {
    count.value += 1
  }
  console.log("#")
}
</script>


<style scoped>
.infinite-list .infinite-list-item {
  box-sizing: border-box;
  padding: 10px;
  display: flex;
  align-items: center;
  height: 50px;
  color: var(--el-color-primary);
}

.infinite-list .infinite-list-item:hover {
  background-color: rgb(24, 24, 24);
}

.infinite-list .infinite-list-item + .list-item {
  margin-top: 10px;
}
</style>