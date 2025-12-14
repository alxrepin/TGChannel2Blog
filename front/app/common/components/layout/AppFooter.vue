<script setup lang="ts">
import UiContainer from '~/common/components/ui/UiContainer.vue';
import { useRoute, computed } from '#imports';

interface Props {
  text: string;
}

defineProps<Props>();

const route = useRoute();
const isMainPage = computed(() => route.path === '/');
</script>

<template>
  <UiContainer as="footer" :size="!isMainPage ? 'sm' : null">
    <div class="f:14 f:fade-50">© {{ new Date().getFullYear() }} {{ text }}</div>
  </UiContainer>
</template>
