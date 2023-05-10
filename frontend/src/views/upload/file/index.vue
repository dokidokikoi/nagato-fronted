<template>
  <el-dialog
    v-model="dialogVisible"
    title="Upload File"
    width="50%"
    destroy-on-close
  >
    <SelectDir ref="selectDirRef" />
    <el-input id="upload" type="text" v-model="filePath" />
    <!-- <el-upload
      class="upload-file"
      drag
      action="https://run.mocky.io/v3/9d059bf9-4660-45f2-925d-ce80ad6c4d15"
      multiple
      :before-upload="beforeUpload"
    >
      <el-icon class="el-icon--upload"><upload-filled /></el-icon>
      <div class="el-upload__text">
        Drop file here or <em>click to upload</em>
      </div>
      <template #tip>
        <div class="el-upload__tip">
          jpg/png files with a size less than 500kb
        </div>
      </template>
    </el-upload> -->
    <template #footer>
      <span class="dialog-footer">
        <el-button @click="dialogVisible = false">Cancel</el-button>
        <el-button type="primary" @click="upload">
          Confirm
        </el-button>
      </span>
    </template>
  </el-dialog>
</template>


<script setup>
import SelectDir from "../../../components/SelectDir.vue"
import { UploadFilled } from '@element-plus/icons-vue'
import { Upload } from '@@/go/main/App'
import { ref } from 'vue';
import { ElNotification } from "element-plus"

const dialogVisible = ref(false)
const filePath = ref("")
const selectDirRef = ref(null)

function showDialog() {
  dialogVisible.value = true
}

function upload() {
  Upload(filePath.value, selectDirRef.value.getUUID()).then(res => {
    if (res.code !== 200) {
        ElNotification({
            title: res.msg,
            type: "error"
        })
        return
    }
  })
  dialogVisible.value = false;
}

defineExpose({
  showDialog
})
</script>


<style scoped>

</style>