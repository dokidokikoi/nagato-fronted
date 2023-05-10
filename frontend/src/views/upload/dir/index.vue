<template>
  <el-dialog
    v-model="dialogVisible"
    title="Create Dir"
    width="50%"
    destroy-on-close
  >
    <SelectDir ref="selectDirRef" />
    <el-input type="text" v-model="dirName" />
    <template #footer>
      <span class="dialog-footer">
        <el-button @click="dialogVisible = false">Cancel</el-button>
        <el-button type="primary" @click="create">
          Confirm
        </el-button>
      </span>
    </template>
  </el-dialog>
</template>


<script setup>
import SelectDir from "../../../components/SelectDir.vue"
import { CreateDir } from '@@/go/main/App'
import { ref } from 'vue';
import { ElNotification } from "element-plus"

const dirName = ref("")
const dialogVisible = ref(false)
const selectDirRef = ref(null)

function showDialog() {
  dialogVisible.value = true
}

function create() {
  CreateDir(dirName.value, selectDirRef.value.getUUID()).then(res => {
    if (res.code !== 200) {
        ElNotification({
            title: res.msg,
            type: "error"
        })
        return
    }
  })
  dialogVisible.value = false
}

defineExpose({
  showDialog
})
</script>


<style scoped>

</style>