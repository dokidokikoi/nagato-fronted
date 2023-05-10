<template>
  <div class="viewer">
    <MdViewer :text="text" />
  </div>
</template>


<script setup>
import { ref } from 'vue'
import MdViewer from '../../../components/MdViewer.vue';
import { getCurrentInstance } from 'vue'
import { ElNotification } from "element-plus"
import { GetBlank } from '@@/go/main/App'

const instance = getCurrentInstance()
const text = ref("")
const loading = ref(true)

function getBlank() {
  loading.value = true
  GetBlank(instance.proxy.$route.params.id).then(res => {
    if (res.code !== 200) {
        ElNotification({
            title: res.msg,
            type: "error"
        })
        loading.value = false
        return
    }
    text.value = res.data.content
    loading.value = false
  })
}

getBlank()
</script>


<style scoped>
.viewer {
  height: calc(100% - 50px);
  overflow-y: auto;
}
</style>