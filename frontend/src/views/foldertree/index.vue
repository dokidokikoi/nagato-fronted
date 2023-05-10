<template>
  <div class="data-panel">
    <el-table
    ref="multipleTableRef"
    :data="tableData"
    :default-sort="{ prop: 'date', order: 'descending' }"
    :scroll-load="load"
    infinite-scroll-distance="100"
    class="folder-panel"
    v-if="!loading"
  >
    <el-table-column type="selection" width="30" />
    <el-table-column label="文件名" prop="name" show-overflow-tooltip >
      <template #default="scope">
        <div style="display: flex; flex-direction: row ; align-items: center;">
          <FileIcon :type="scope.row.dir ? 'dir':scope.row.ext"/>
          <a  style="font-size: 1.1em; padding: 3px;" @click="getFolderChildren(scope.row.uuid)">{{ scope.row.name }}</a>
        </div>
      </template>
    </el-table-column>
    <el-table-column prop="created_at" label="修改日期" sortable show-overflow-tooltip />
    <el-table-column prop="size" label="大小" show-overflow-tooltip  width="120">
      <template #default="scope">
        {{ scope.row.dir ? "-":scope.row.size }}
      </template>
    </el-table-column>
    <el-table-column label="操作">
      <template #default="scope">
      
      </template>
    </el-table-column>
  </el-table>
  </div>
</template>


<script setup>
import FileIcon from "./icon/index.vue"
import { ref } from 'vue'
import { GetFolderChildren } from '@@/go/main/App'
import { ElNotification } from "element-plus"

const tableData = ref(null)
const loading = ref(true)

function getFolderChildren(uid) {
  loading.value = true
  GetFolderChildren(uid).then(res => {
    if (res.code !== 200) {
        ElNotification({
            title: res.msg,
            type: "error"
        })
        loading.value = false
        return
    }
    tableData.value = res.data.children
    loading.value = false
  })
}

getFolderChildren()

</script>


<style scoped>
/* .el-table {
  background-color: inherit;
  --el-table-tr-bg-color: rgb(32, 32, 32);
} */
/* ::v-deep .el-table__body-wrapper {
  border-bottom: 0 !important;
} */
/* 更改表头和表格行的背景色 */
/* ::v-deep .el-table__row, ::v-deep th.el-table__cell.is-leaf {
  background-color: rgb(32, 32, 32); 
  color: #797878;
  border-color: aqua !important;
} */
/* ::v-deep .el-table__body td {
  color: #797878 !important;
  border-color: aqua !important;
} */

.folder-panel {
  overflow: auto;
  height: calc(100% + 1px);
}

a {
  cursor: pointer;
}

a:hover {
  
}
</style>