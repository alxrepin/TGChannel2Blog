<script setup lang="ts">
import type { Post } from '../types';

interface Props {
  post: Post;
}

const props = defineProps<Props>();

const { title, excerpt, publishedAt, thumbnail, hasImage } = toRefs(props.post);

const formattedDate = computed(() => {
  const date = new Date(props.post.publishedAt);
  return new Intl.DateTimeFormat('ru-RU', {
    day: 'numeric',
    month: 'short',
    year: 'numeric',
  }).format(date);
});
</script>

<template>
  <UiCard variant="elevated" padding="none" class="flex height:240px">
    <div
      :class="['jc:space-between flex flex:column', hasImage ? 'width:70% p:20' : 'p:20 w:full']"
    >
      <div>
        <h2 class="f:20 f:semibold f:fade-10 max-height:56px overflow:hidden lines:2 text:ellipsis">
          {{ title }}
        </h2>
        <p
          class="f:14 pt:8 color:fade-54 lh:1.6 max-height:112px overflow:hidden lines:5 text:ellipsis box:content"
        >
          {{ excerpt }}
        </p>
      </div>
      <time :datetime="publishedAt" class="f:14 color:fade-64">
        {{ formattedDate }}
      </time>
    </div>

    <div v-if="hasImage && thumbnail" class="my:4 mr:4 r:24 width:30% overflow:hidden">
      <NuxtImg
        :src="thumbnail"
        :alt="title"
        width="100%"
        class="h:100% w:100% object:contain scale(120%) scale(220%)@ms"
        loading="lazy"
      />
    </div>
  </UiCard>
</template>
