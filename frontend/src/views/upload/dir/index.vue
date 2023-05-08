<template>
  <el-dialog
    v-model="dialogVisible"
    title="Create Dir"
    width="50%"
    :before-close="handleClose"
  >
    <span>This is a message</span>
    <template #footer>
      <span class="dialog-footer">
        <el-button @click="dialogVisible = false">Cancel</el-button>
        <el-button type="primary" @click="dialogVisible = false">
          Confirm
        </el-button>
      </span>
    </template>
  </el-dialog>
</template>


<script setup>
import { ref } from 'vue';

const dialogVisible = ref(false)

function showDialog() {
  dialogVisible.value = true
}

defineExpose({
  showDialog
})
</script>


<style scoped>

</style>