<template>
  <el-container class="common-layout">
    <el-header height="50px" class="header">
      <Header />
    </el-header>
    <el-container>
      <el-aside width="200px" class="aside">
        <Aside />
      </el-aside>
      <el-main class="main">
        <!-- <FolderTree /> -->
        <BlankEdit />
      </el-main>
    </el-container>
  </el-container>
</template>


<script setup>
import Header from "./header/index.vue"
import Aside from "./aside/index.vue"
import FolderTree from "../foldertree/index.vue"
import BlankEdit from "../blank/edit/index.vue"
</script>


<style scoped>
.common-layout {
  height: 100%;
}

.header {
  background-color: rgb(31, 31, 31);
}

.aside {
  background-color: rgb(24, 24, 24);
}

.main {
  height: calc(100vh - 50px) !important;
  background-color: rgb(28, 28, 28);
  position: relative;
  overflow: hidden;
  height: 100%;
  padding: 0 !important;
}
</style>