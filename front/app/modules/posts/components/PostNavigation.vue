<script setup lang="ts">
import { computed } from '#imports';

interface Props {
  publishedAt: string;
}

const props = defineProps<Props>();

const formattedDate = computed(() => {
  const date = new Date(props.publishedAt);
  return new Intl.DateTimeFormat('ru-RU', {
    day: 'numeric',
    month: 'short',
    year: 'numeric',
  }).format(date);
});
</script>

<template>
  <div class="flex jc:space-between pb:10 f:16 color:fade-64">
    <NuxtLink
      to="/"
      class="flex ai:center gap:6 text-decoration:none color:fade-64 hover:color:fade-10 transition"
    >
      <Icon name="mdi:arrow-left" />
      Назад
    </NuxtLink>
    <time :datetime="publishedAt" class="block">
      {{ formattedDate }}
    </time>
  </div>
</template>
