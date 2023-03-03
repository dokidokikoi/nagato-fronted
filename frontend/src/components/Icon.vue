<template>
  <component :is="icon" class="fill-current" />
</template>

<script setup>
import { defineAsyncComponent } from 'vue';

const props = defineProps({
  name: {
    type: String,
    required: true,
  },
});

const icon = defineAsyncComponent(() =>
  import(`../assets/svgs/${props.name}.svg`)
);
</script>

<style scoped>
.fill-current {
  width: 18px;
  height: 18px;
  display: flex;
  justify-content: center;
  color: blue;
}
</style>