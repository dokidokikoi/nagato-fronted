<template>
  <div class="parent">
    <ul v-infinite-scroll="load" class="infinite-list fill" style="overflow: auto">
      <li v-for="i in count" :key="i" class="infinite-list-item" @click="toRoute('/blank/view')">
        <router-link to="/blank/view">{{i}}</router-link>
      </li>
    </ul>
  </div>
</template>


<script setup>
import { ref } from 'vue'
import router from '@/router/router.js';
const count = ref(20)
const load = () => {
  count.value += 5
}
function toRoute(path) {
  router.push({path: path})
}
</script>


<style scoped>
.parent {
  display: flex;
  flex-direction: column;
  height: calc(100% - 50px);
  overflow: hidden; /* Hide overflowing content from the parent */
}
.fill {
  list-style: none;
  padding: 10px;
  margin: 0;
  flex-grow: 1;
  overflow: auto; /* Enable scrolling when content exceeds the height */
}
.infinite-list .infinite-list-item {
  box-sizing: border-box;
  padding: 10px;
  display: flex;
  align-items: center;
  color: var(--el-color-primary);
}

.infinite-list .infinite-list-item:hover {
  background-color: rgb(24, 24, 24);
  cursor: pointer;
}

.infinite-list-item {
  border-radius: 5px;
}

</style>