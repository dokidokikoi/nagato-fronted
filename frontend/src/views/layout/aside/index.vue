<template>
  <div class="me">
    <!-- <el-avatar src="../../../assets/images/avatar.png" /> -->
    <div class="avatar" @click="toRoute('/me')">
      <img src="../../../assets/images/avatar.png" style="object-fit: 100%;" alt="">
    </div>
    <span class="name" @click="toRoute('/me')"> 
      <span>nagato</span> <Icon name="arrow-right" />
    </span>
  </div>
  <div class="menu">
    <el-menu
      background-color="#181818"
      text-color="aliceblue"
      router
      >
        <el-menu-item index="/">
          <el-icon><folder /></el-icon>
          <span>我的资源</span>
        </el-menu-item>
        <el-menu-item index="/blank">
          <el-icon><reading /></el-icon>
          <span>Blank</span>
        </el-menu-item>
        <el-menu-item index="/">
          <el-icon><Picture /></el-icon>
          <span>相册</span>
        </el-menu-item>
        <el-menu-item index="/history">
          <Icon name="history" style="padding: 0 7px 0 3px;"/>
          <span>历史</span>
        </el-menu-item>
        <el-menu-item index="/share">
          <el-icon><Share /></el-icon>
          <span>分享</span>
        </el-menu-item>
      </el-menu>
  </div>
</template>


<script setup>
import Icon from "../../../components/Icon.vue"
import {
  Picture,
  Folder,
  Reading,
  Share
} from '@element-plus/icons-vue'
import router from "@/router/router.js"

function toRoute(path) {
  router.push({path: path})
}
</script>


<style scoped>
.me {
  height: 60px;
  background-color: rgb(21, 21, 21);
  box-sizing: border-box;
  padding: 0 10px;
  display: flex;
  flex-direction: row;
  align-items: center;
}

.avatar {
  width: 40px;
  height: 40px;
  background-color: aliceblue;
  border-radius: 50%;
  overflow: hidden;
}

.avatar img {
  width: inherit;
  height: inherit;
}

.avatar, .name:hover {
  cursor: pointer;
}

.name {
  text-align: center;
  margin-left: 10px;
  display: flex;
  align-items: center;
}

.el-menu-item.is-active {
  color: aliceblue;
  background-color: rgb(21, 21, 21) !important;
}

.el-menu-item:hover {
  background-color: rgb(30, 30, 30);
}

.el-input__inner:focus {
  border-color: black !important;
}

.el-menu {
  border: none;
}
</style>