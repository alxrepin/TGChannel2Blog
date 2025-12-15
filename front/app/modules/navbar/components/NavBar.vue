<script setup lang="ts">
import UiButton from '~/common/components/ui/UiButton.vue';
import UiContainer from '~/common/components/ui/UiContainer.vue';
import { useChannelData } from '~/modules/channel/composables/use-channel-data';
import { usePageData } from '~/common/composables/use-page-data';
import { useIntersectionObserver } from '~/common/composables/use-intersection-observer';

const { data: channel } = useChannelData();
const { isMainPage } = usePageData();
const { isIntersecting: isHeroVisible } = useIntersectionObserver('#hero', {
  threshold: 0,
  enabled: isMainPage,
});
</script>

<template>
  <UiContainer
    v-show="!isMainPage || !isHeroVisible"
    as="nav"
    aria-label="Основная навигация"
    class="fixed top:0 w:full z:10 contain:layout|style"
    :padding="false"
  >
    <UiContainer
      as="div"
      :size="!isMainPage ? 'sm' : null"
      class="flex jc:space-between ai:center height:40px z:10 rel px:24 py:48 px:48@sm"
      :padding="false"
    >
      <NuxtLink to="/" class="flex ai:center gap:10 text-decoration:none">
        <NuxtImg
          v-if="channel.avatar"
          :src="channel.avatar"
          :alt="channel.title"
          :width="40"
          :height="40"
          class="rounded"
          style="border-radius: 50%"
          loading="lazy"
        />
        <div class="f:20 f:bold f:fade-10">
          {{ channel.title }}
        </div>
      </NuxtLink>

      <UiButton tag="a" :href="channel.url" target="_blank" rel="noopener noreferrer" size="sm">
        <Icon name="mdi:telegram" />
        Подписаться
      </UiButton>
    </UiContainer>
    <div class="abs h:full w:full bd:blur(5px) top:0 left:0"></div>
  </UiContainer>
</template>
