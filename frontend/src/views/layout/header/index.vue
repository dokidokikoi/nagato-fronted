<template>
  <div class="header">
    <h1>nagato</h1>
    <div class="btn-set">
      <div class="text-btn">
        <div class="search">
          <el-input
            class="search_inner"
            v-model="search"
            size="small"
            width="3000"
            placeholder="搜索"
            :prefix-icon="Search"
          />
        </div>
        <!-- <el-button class="btn" link @click.stop>下载</el-button> -->
        <!-- <el-button class="btn" link @click.stop>上传</el-button> -->
        
      </div>
      <Icon name="circle-download" class="icon-btn"/>
      <Icon name="cloud-upload" class="icon-btn"/>
      <Icon name="refresh" class="icon-btn"/>
      <Icon name="setting" class="icon-btn"/>
    </div>
  </div>
</template>


<script setup>
import Icon from "../../../components/Icon.vue"
import { Search } from '@element-plus/icons-vue'
import { ref } from "vue";

let search = ref("")
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
  align-items: center;
}

.text-btn {
  padding: 0 20px;
  display: flex;
  flex-direction: row;
  justify-content: center;
}

.search {
  /* padding: 0 10px 0 20px;  */
}

.search_inner {
  width: 150px;
}
</style>