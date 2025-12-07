<script setup lang="ts">
interface Props {
  name: string
  description?: string
}

defineProps<Props>()
</script>

<template>
  <div class="flex flex:column ai:center text:center">
    <h1 class="f:28 f:bold pb:10 f:fade-10">
      {{ name }}
    </h1>
    <p v-if="description" class="f:16 w:full w:2/3@sm pb:20 color:fade-50 lh:1.6">
      {{ description }}
    </p>
  </div>
</template>
