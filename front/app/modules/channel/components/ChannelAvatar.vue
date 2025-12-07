<script setup lang="ts">
interface Props {
  src: string
  alt: string
  size?: number
}

withDefaults(defineProps<Props>(), {
  size: 200,
})
</script>

<template>
  <NuxtImg
    :src="src"
    :alt="alt"
    :width="size"
    :height="size"
    class="rounded"
    loading="lazy"
  />
</template>
