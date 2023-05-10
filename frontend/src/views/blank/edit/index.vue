<template>
  <div class="data-panel edit-div">
    <MdEditer class="editer" ref="editerRef" @saveEvent="shouDrawer" />
    <el-drawer v-model="drawer">
      <template #header>
        <h4>save artcle</h4>
      </template>
      <template #default>
        <el-form :model="blank" class="blank-form">
          <el-form-item label="type">
            <el-input v-model="blank.type" placeholder="type" />
          </el-form-item>
          <el-form-item label="title">
            <el-input v-model="blank.title" placeholder="title" />
          </el-form-item>
          <el-form-item label="tags">
            <el-input v-model="blank.tags" placeholder="tags" />
          </el-form-item>
          <el-form-item label="matters">
            <el-input v-model="blank.matters" placeholder="matters" />
          </el-form-item>
        </el-form>
        {{ blank.tags }}
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
import { CreateBlank } from '@@/go/main/App'
import { ref, reactive } from "vue"
import { ElNotification } from "element-plus"

const drawer = ref(false)
const loading = ref(true)
const editerRef = ref(null)
let blank = reactive({})

function shouDrawer() {
  drawer.value = true
}

function confirmClick() {
  blank.content = editerRef.value.getText()
  CreateBlank(blank).then(res => {
    if (res.code !== 200) {
        ElNotification({
            title: res.msg,
            type: "error"
        })
        loading.value = false
        drawer.value = false
        return
    }
    ElNotification({
        title: "保存成功",
        type: "success"
    })
    loading.value = false
    drawer.value = false
  })
}

function cancelClick() {
  drawer.value = false
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