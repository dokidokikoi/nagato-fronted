<template>
  <v-md-editor 
  class="editer"
  :disabled-menus="[]" 
  v-model="text" 
  @upload-image="handleUploadImage"
  @save="saveText"
  @image-click="imageFocus"
  tab-size="4"
  :include-level="[1, 2, 3,4]"
  ></v-md-editor>
  <el-dialog v-model="dialogImageVisible" center>
    <el-carousel :initial-index="dialogImageIndex" :autoplay="false" height="400px">
      <el-carousel-item v-for="item in dialogImages" :key="item">
        <el-image :src="item" fit="cover"/>
      </el-carousel-item>
    </el-carousel>
  </el-dialog>
</template>

<script setup>
import { ref } from "vue"
const emit = defineEmits(['saveEvent'])

let text = ref("")

let dialogImageVisible = ref(false)
let dialogImages = ref([])
let dialogImageIndex = ref(0)

function handleUploadImage(event, insertImage, files) {
  // 拿到 files 之后上传到文件服务器，然后向编辑框中插入对应的内容
  console.log(files);

  // 此处只做示例
  insertImage({
    url:
      'https://www.freecodecamp.org/news/content/images/2021/10/golang.png',
    desc: 'golang',
    // width: 'auto',
    // height: 'auto',
  });
}

function saveText() {
  // localStorage.setItem('edit', text.value)
  // console.log(localStorage.getItem('edit'))
  emit('saveEvent')
  console.log("hi")
}

function imageFocus(images, currentIndex) {
  // console.log(images, currentIndex)
  dialogImageVisible.value = true
  dialogImages.value = images
  dialogImageIndex = currentIndex
}

function getText() {
  return text.value
}

defineExpose({
  getText
})
</script>


<style scoped>
.editer {
  height: 900px;
  /* overflow: auto; */
}
</style>