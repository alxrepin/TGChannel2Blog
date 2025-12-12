<template>
  <UiContainer as="main" :padding="false">
    <slot />
  </UiContainer>
</template>

<script setup lang="ts">
import UiContainer from '~/common/components/ui/UiContainer.vue';
</script>
