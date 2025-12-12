<script setup lang="ts">
interface Props {
  telegramHandle: string;
  size?: string;
}

const props = withDefaults(defineProps<Props>(), {
  size: 'md',
});

const telegramLink = computed(() => `https://t.me/${props.telegramHandle}`);
</script>

<template>
  <UiButton tag="a" :href="telegramLink" target="_blank" rel="noopener noreferrer" :size="size">
    <Icon name="mdi:telegram" />
    Подписаться
  </UiButton>
</template>
