<script setup lang="ts">
import UiContainer from '~/common/components/ui/UiContainer.vue';
</script>

<template>
  <UiContainer as="header" :padding="false">
    <slot />
  </UiContainer>
</template>
