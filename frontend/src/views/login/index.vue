<template>
  <el-form
    ref="ruleFormRef"
    :model="loginForm"
    status-icon
    label-width="120px"
    class="demo-ruleForm"
  >
    <el-form-item label="account" prop="account">
      <el-input v-model="loginForm.account" autocomplete="off" />
    </el-form-item>
    <el-form-item label="password" prop="password">
      <el-input
        v-model="loginForm.password"
        type="password"
        autocomplete="off"
      />
    </el-form-item>
    <el-form-item label="email" prop="email">
      <el-input v-model="loginForm.email" />
    </el-form-item>
    
    <el-form-item>
      <el-button type="primary"
        >Submit</el-button
      >
      <el-button>Reset</el-button>
    </el-form-item>
  </el-form>
</template>


<script setup>
import { reactive, ref } from 'vue'

const loginForm = reactive({
  account: 'admin',
  password: '123456',
  email: "",
  key: '',
  code: ''
})

</script>


<style scoped>

</style>