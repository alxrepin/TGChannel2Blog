<script setup lang="ts">
import HeroTitle from '~/modules/hero/components/HeroTitle.vue';
import SubscribeButton from '~/features/SubscribeButton.vue';
import HeroStatistics from '~/modules/hero/components/HeroStatistics.vue';
import HeroAvatar from '~/modules/hero/components/HeroAvatar.vue';
import { useChannelData } from '~/common/composables/useChannelData';

const data = useChannelData();
</script>

<template>
  <div id="hero" class="flex flex:column ai:center text:center">
    <HeroAvatar :src="data.avatar" :alt="data.name" class="mb:20" />
    <HeroTitle :name="data.name" :description="data.description" />
    <div class="pb:20">
      <HeroStatistics :subscriber-count="data.subscriberCount" />
    </div>
    <SubscribeButton :telegram-handle="data.telegramHandle" />
  </div>
</template>
