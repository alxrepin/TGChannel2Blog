<script setup lang="ts">
import HeroTitle from '~/modules/hero/components/HeroTitle.vue';
import HeroStatistics from '~/modules/hero/components/HeroStatistics.vue';
import HeroAvatar from '~/modules/hero/components/HeroAvatar.vue';
import { useChannelData } from '~/modules/channel/composables/use-channel-data';
import UiButton from '~/common/components/ui/UiButton.vue';
import UiContainer from '~/common/components/ui/UiContainer.vue';

const { data: channel } = useChannelData();
</script>

<template>
  <UiContainer
    id="hero"
    as="div"
    class="flex flex:column ai:center text:center my:88"
    :padding="false"
  >
    <HeroAvatar v-if="channel.avatar" :src="channel.avatar" :alt="channel.title" class="mb:20" />
    <HeroTitle :name="channel.title" :description="channel.description" />
    <div class="pb:20">
      <HeroStatistics :subscribers="channel.subscriptions" />
    </div>
    <UiButton tag="a" :href="channel.url" target="_blank" rel="noopener noreferrer" size="md">
      <Icon name="mdi:telegram" />
      Подписаться
    </UiButton>
  </UiContainer>
</template>
