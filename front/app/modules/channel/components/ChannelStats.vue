<script setup lang="ts">
interface Props {
  subscriberCount: number
}

defineProps<Props>()
</script>

<template>
  <div class="flex gap:8 f:14 f:fade-30">
    <div class="bg:fade-94 px:12 py:4 r:12 flex ai:center gap:6">
      <Icon name="mdi:users" />
      {{ subscriberCount }} подписчиков
    </div>
  </div>
</template>
