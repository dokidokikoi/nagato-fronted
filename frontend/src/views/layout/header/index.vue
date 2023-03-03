<template>
  <div class="header">
    <h1>nagato</h1>
    <div class="btn-set">
      <div class="text-btn">
        <el-input
          v-model="input1"
          size="small"
          placeholder="搜索"
          :prefix-icon="Search"
          class="search"
        />
        <el-button class="btn" link @click.stop>下载</el-button>
        <el-button class="btn" link @click.stop>上传</el-button>
      </div>
      <Icon name="setting" class="icon-btn"/>
    </div>
  </div>
</template>


<script setup>
import Icon from "../../../components/Icon.vue"
import { Search } from '@element-plus/icons-vue'
</script>


<style scoped>
.header {
  height: 100%;
  width: 100%;
  display: flex;
  flex-direction: row;
  justify-content: space-between;
  align-items: center;
}

.btn-set {
  display: flex;
  flex-direction: row;
  justify-content: center;
}

.text-btn {
  width: 40%;
  padding: 0 23px;
  display: flex;
  flex-direction: row;
  justify-content: center;
}

.icon-btn:hover {
  cursor: pointer;
}

/* .search ::v-deep .el-input__wrapper {
  background-clip: padding-box;
  background-color: blue;
} */
</style>