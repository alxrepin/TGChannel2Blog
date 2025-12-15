<script setup lang="ts">
import type { TPostSnippet } from '../types';
import UiCard from '~/common/components/ui/UiCard.vue';
import { computed, toRefs } from '#imports';

interface Props {
  post: TPostSnippet;
}

const props = defineProps<Props>();

const { title, text, thumbnail, createdAt } = toRefs(props.post);

const formattedDate = computed(() => {
  const date = new Date(createdAt.value);

  return new Intl.DateTimeFormat('ru-RU', {
    day: 'numeric',
    month: 'short',
    year: 'numeric',
  }).format(date);
});
</script>

<template>
  <UiCard
    variant="elevated"
    padding="none"
    class="flex flex:column flex:row@xs min-height:240px max-height:480px height:240px@xs"
  >
    <div v-if="thumbnail" class="my:4 mx:4 r:24 overflow:hidden block hide@xs">
      <NuxtImg
        :src="thumbnail"
        :alt="title"
        width="100%"
        class="h:100% w:100% object:contain scale(120%) scale(280%)@ms"
        loading="lazy"
      />
    </div>
    <div
      :class="[
        'jc:space-between flex flex:column h:full',
        thumbnail ? 'width:70%@xs p:20' : 'p:20 w:full',
      ]"
    >
      <div>
        <h2 class="f:20 f:semibold f:fade-10 max-height:56px overflow:hidden lines:2 text:ellipsis">
          {{ title }}
        </h2>
        <p
          class="f:14 pt:8 color:fade-54 lh:1.6 max-height:112px overflow:hidden lines:5 text:ellipsis box:content"
          v-html="text"
        />
      </div>
      <time :datetime="formattedDate" class="f:14 color:fade-64">
        {{ formattedDate }}
      </time>
    </div>

    <div v-if="thumbnail" class="my:4 mr:4 r:24 width:30% overflow:hidden hide block@xs">
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
