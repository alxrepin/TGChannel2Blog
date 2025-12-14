<script setup lang="ts">
import AppHeader from '~/common/components/layout/AppHeader.vue';
import AppContent from '~/common/components/layout/AppContent.vue';
import AppFooter from '~/common/components/layout/AppFooter.vue';
</script>

<template>
  <AppHeader />
  <AppContent class="py:88">
    <slot />
  </AppContent>
  <AppFooter text="R — Repin" />
</template>
