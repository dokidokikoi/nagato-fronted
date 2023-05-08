<template>
  <!-- <MainHeader /> -->
  <!-- <ul v-infinite-scroll="load" class="data-panel infinite-list" style="overflow: auto">
    <div class="folder-header">
      <el-checkbox
        v-model="checkAll"
        :indeterminate="isIndeterminate"
        @change="handleCheckAllChange"
        style="width: 100px;"
        >文件名</el-checkbox
      >
      <span>修改时间</span>
      <span style="width: 100px;">大小</span>
    </div>
    <li v-for="i in count" :key="i" class="infinite-list-item">
      <el-checkbox :label="i">
        {{ i }}
      </el-checkbox>
    </li>
  </ul> -->
  <div class="data-panel">
    <el-table
    ref="multipleTableRef"
    :data="tableData"
    :default-sort="{ prop: 'date', order: 'descending' }"
    :scroll-load="load"
    infinite-scroll-distance="100"
    class="folder-panel"
  >
    <el-table-column type="selection" width="55" />
    <el-table-column label="文件名" prop="name" show-overflow-tooltip />
    <el-table-column prop="date" label="修改日期" sortable show-overflow-tooltip />
    <el-table-column prop="size" label="大小" show-overflow-tooltip  width="120"/>
    <el-table-column label="操作">
      <template #default="scope">
      
      </template>
    </el-table-column>
  </el-table>
  </div>
</template>


<script setup>
import MainHeader from "../../components/MainHeader.vue"
import { ref } from 'vue'

const load = () => {
  // if (count.value < 15) {
  //   count.value += 1
  // }
  if (tableData.value.length < 30) {
    tableData.value.push({
      date: '2016-05-03',
      name: 'Tom',
      address: 'No. 189, Grove St, Los Angeles',
    })
    console.log("#")
  }
  console.log("#")
}

const tableData = ref([
  {
    name: 'video',
    date: '2023-02-23 12:34:22',
    size: 2345455,
  },
  {
    name: 'video',
    date: '2023-02-23 12:34:22',
    size: 2345455,
  },
])

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

/* 更改表头和表格行的文字颜色 */
/* ::v-deep .el-table__header th,
::v-deep .el-table__body td {
  color: #303133; 
} */
</style>