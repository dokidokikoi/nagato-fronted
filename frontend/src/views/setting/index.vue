<template>
  <el-drawer v-model="settingVisable">
    <template #header>
      <h4>edit setting</h4>
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
</template>


<script setup>
import { ref } from 'vue';

const props = defineProps({settingVisable: Boolean})
const settingVisable = ref(false) 

function showSetting() {
  settingVisable.value = true
} 

defineExpose({
  showSetting
})
</script>


<style scoped>

</style>