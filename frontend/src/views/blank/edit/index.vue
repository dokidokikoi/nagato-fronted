<template>
  <div class="data-panel edit-div">
    <MdEditer class="editer" @saveEvent="shouDrawer" />
    <el-drawer v-model="drawer">
      <template #header>
        <h4>save artcle</h4>
      </template>
      <template #default>
        <div>
          <el-radio v-model="radio1" label="Option 1" size="large"
            >Option 1</el-radio
          >
          <el-radio v-model="radio1" label="Option 2" size="large"
            >Option 2</el-radio
          >
        </div>
      </template>
      <template #footer>
        <div style="flex: auto">
          <el-button @click="cancelClick">取消</el-button>
          <el-button type="primary" @click="confirmClick">保存</el-button>
        </div>
      </template>
    </el-drawer>
  </div>
</template>


<script setup>
import MdEditer from '../../../components/MdEditer.vue';
import { ref } from "vue"
const drawer = ref(false)
function shouDrawer() {
  console.log("hello")
  drawer.value = true
}
</script>


<style scoped>
.data-panel.edit-div {
  background-color: #fff;
  overflow-y: auto;
}
.editer {
  height: 100vh;
}
</style>