<template>
  <div class="main-header">
    <div class="breadcrumb">
      <span class="breadcrumb-item">
        <span class="breadcrumb-item is-active">/</span>
      </span>
      <span style="color: #797878;"> > </span>
      <span class="breadcrumb-item">
        <span class="breadcrumb-item">我的资源</span>
      </span>
    </div>
    <div>
      <div class="btn-group">
        <el-button type="primary" :icon="Upload" size="small" round @click="showUploadDailog()">上传文件</el-button>
        <el-button color="w" :icon="FolderAdd" size="small" link @click="showDirDailog()">新建文件夹</el-button>
        <el-button color="w" :icon="DocumentAdd" size="small" link @click="toRoute('/blank/edit')">新建Blank</el-button>
        <el-button color="w" :icon="Download" size="small" link @click="toRoute('/blank/edit')">离线下载</el-button>
      </div>
    </div>
  </div>
  <DirDialog ref="dirDialogRef" />
  <FileDialog ref="fileDialogRef" />
</template>


<script setup>
import DirDialog from '@/views/upload/dir/index.vue'
import FileDialog from '@/views/upload/file/index.vue'
import { DocumentAdd, FolderAdd, Download, Upload } from '@element-plus/icons-vue'
import router from '@/router/router.js';
import { ref } from 'vue';

const dirDialogRef = ref(null)
const fileDialogRef = ref(null)

function toRoute(path) {
  router.push({path: path})
}

function showDirDailog() {
  dirDialogRef.value.showDialog()
}

function showUploadDailog() {
  fileDialogRef.value.showDialog()
}
</script>


<style scoped>
.main-header {
  width: 100%;
  height: 50px;
  background-color: rgb(24, 24, 24);
  box-sizing: border-box;
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 0 20px;
  background-clip: padding-box;
}

.breadcrumb-item {
  color: #797878;
  font-size: 0.9em;
}

.breadcrumb-item:hover {
  cursor: pointer;
  color: #b5b4b4;
}

.breadcrumb-item.is-active {
  color: #fff;
}


</style>