<template>
  <div class="me">
    <!-- <el-avatar src="../../../assets/images/avatar.png" /> -->
    <div class="avatar">
      <img src="../../../assets/images/avatar.png" style="object-fit: 100%;" alt="">
    </div>
    <span class="name">
      <span>nagato</span> <Icon name="arrow-right" />
    </span>
  </div>
</template>


<script setup>
import Icon from "../../../components/Icon.vue"
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

.name {
  text-align: center;
  margin-left: 10px;
  display: flex;
  align-items: center;
}
</style>