<template>
  <Icon :name="getIconName(type)" @click="showSetting()" class="f-icon"/>
</template>


<script setup>
import Icon from "../../../components/Icon.vue"

let props = defineProps(["type"])

console.log(props.type)

function getIconName(type) {
  switch (type) {
    case "dir": return "file folder"
    case "txt": return "txt"
    case "png": return "img"
  }
}
</script>


<style scoped>
.f-icon {
  width: 30px;
  height: auto;
}
</style>