<template>
  <div style="display: flex; align-items: center;">
    <h3>文件路径</h3>
    <span>当前路径：{{ current ? current.path:"/" }}</span>
  </div>
  <ul>
    <li v-for="item in dirs" :key="item.id">
      <div style="display: flex; flex-direction: row ; align-items: center;">
        <FileIcon v-if="!loading" :type="item.dir ? 'dir':item.ext"/>
        <a  style="font-size: 1.1em; padding: 3px;" @click="getFolderChildren(item)">{{item.name}}</a>
      </div>
    </li>
  </ul>
</template>


<script setup>
import { GetFolderChildren } from '@@/go/main/App'
import { ref } from 'vue'

const dirs = ref(null)
const loading = ref(true)
const current = ref(null)

function getFolderChildren(item) {
  loading.value = true
  let uid = ""
  if (item) {
    uid = item.uuid
  }
  current.value = item
  GetFolderChildren(uid).then(res => {
    if (res.code !== 200) {
        ElNotification({
            title: res.msg,
            type: "error"
        })
        loading.value = false
        return
    }
    dirs.value = res.data.children
    loading.value = false
  })
}

function getUUID() {
  if (current) {
    return ""
  }
  return current.value.uuid
}

getFolderChildren()

defineExpose({
  getUUID
})
</script>


<style scoped>
ul {
  list-style: none;
  margin: 0;
  padding: 0;
  height: 100px;
  overflow: auto;
  background-color: rgb(249, 247, 247);
}

li {
  margin: 0;
  padding: 10px;
}

li:hover {
  background-color: rgb(229, 226, 226);
}
</style>