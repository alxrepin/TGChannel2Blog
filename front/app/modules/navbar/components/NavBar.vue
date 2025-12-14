<script setup lang="ts">
import UiButton from '~/common/components/ui/UiButton.vue';
import UiContainer from '~/common/components/ui/UiContainer.vue';
import { useRoute, computed } from '#imports';

interface Props {
  title: string;
  image: string;
  url: string;
}

defineProps<Props>();

const route = useRoute();
const isMainPage = computed(() => route.path === '/');
</script>

<template>
  <UiContainer
    :size="!isMainPage ? 'sm' : null"
    class="flex jc:space-between ai:center height:40px z:10 rel px:24 py:48 px:48@sm"
    aria-label="Основная навигация"
    :padding="false"
  >
    <NuxtLink to="/" class="flex ai:center gap:10 text-decoration:none">
      <NuxtImg
        :src="image"
        :alt="title"
        :width="40"
        :height="40"
        class="rounded"
        style="border-radius: 50%"
        loading="lazy"
      />
      <div class="f:20 f:bold f:fade-10">
        {{ title }}
      </div>
    </NuxtLink>

    <UiButton
      tag="a"
      :href="'https://t.me/' + url"
      target="_blank"
      rel="noopener noreferrer"
      size="sm"
    >
      <Icon name="mdi:telegram" />
      Подписаться
    </UiButton>
  </UiContainer>
  <div class="abs h:full w:full bd:blur(5px) top:0 left:0"></div>
</template>
