<template>
  <div class="parent">
    <ul v-infinite-scroll="load" class="infinite-list fill" style="overflow: auto">
      <li v-for="item in blanks" :key="item.id" class="infinite-list-item" @click="toRoute('/blank/view/'+item.id)">
        <router-link :to="'/blank/view/'+item.id">{{item.title}}</router-link>
      </li>
    </ul>
  </div>
</template>


<script setup>
import { ref } from 'vue'
import router from '@/router/router.js';
import { BlankList } from '@@/go/main/App'
import { ElNotification } from "element-plus"

const blanks = ref(null)

const load = () => {
}
function toRoute(path) {
  router.push({path: path})
}
function blankList() {
  BlankList().then(res => {
    if (res.code !== 200) {
        ElNotification({
            title: res.msg,
            type: "error"
        })
        return
    }
    blanks.value = res.data
  })
}

blankList()
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