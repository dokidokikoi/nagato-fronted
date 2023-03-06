<template>
  <MainHeader />
  <div class="data-panel edit-div">
    <div>
      
    </div>
    <MdEditer class="editer" />
  </div>
</template>


<script setup>
import MainHeader from '../../../components/MainHeader.vue';
import MdEditer from '../../../components/MdEditer.vue';
</script>


<style scoped>
.data-panel.edit-div {
  background-color: #fff;
}
.editer {
  /* height: inherit; */
}
</style>