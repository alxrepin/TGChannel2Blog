<script setup lang="ts">
import AppHeader from '~/common/components/layout/AppHeader.vue';
import AppContent from '~/common/components/layout/AppContent.vue';
import AppFooter from '~/common/components/layout/AppFooter.vue';
import { useAsyncData } from '#app';
import { fetchChannel } from '~/modules/channel/api/channel';
import NavBar from '~/modules/navbar/components/NavBar.vue';
import Hero from '~/modules/hero/components/Hero.vue';

const { data: channel } = await useAsyncData('channel', fetchChannel);
</script>

<template>
  <AppHeader>
    <NavBar />
    <Hero />
  </AppHeader>
  <AppContent>
    <slot />
  </AppContent>
  <AppFooter :text="channel.title" />
</template>
