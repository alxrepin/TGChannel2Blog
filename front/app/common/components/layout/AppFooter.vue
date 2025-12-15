<script setup lang="ts">
import UiContainer from '~/common/components/ui/UiContainer.vue';
import { usePageData } from '~/common/composables/use-page-data';

interface Props {
  text: string;
}

defineProps<Props>();

const { isMainPage } = usePageData();
</script>

<template>
  <UiContainer as="footer" :size="!isMainPage ? 'sm' : null">
    <div class="f:14 f:fade-50">© {{ new Date().getFullYear() }} {{ text }}</div>
  </UiContainer>
</template>
