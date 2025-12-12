<script setup lang="ts">
import UiButton from '~/common/components/ui/UiButton.vue';

interface Props {
  title: string;
  image: string;
  url: string;
}

defineProps<Props>();
</script>

<template>
  <nav
    class="flex jc:space-between ai:center height:40px mb:32 z:10"
    aria-label="Основная навигация"
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
  </nav>
  <div class="abs h:70% w:full bd:blur(5px) top:0 left:0"></div>
</template>
