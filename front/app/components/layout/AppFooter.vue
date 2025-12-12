<script setup lang="ts">
interface Props {
  test: string;
}

defineProps<Props>();

const currentYear = new Date().getFullYear();
</script>

<template>
  <footer>
    <div class="f:14 f:fade-50">© {{ currentYear }} {{ text }}</div>
  </footer>
</template>
