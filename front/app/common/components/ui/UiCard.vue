<script setup lang="ts">
import { computed } from '#imports';

interface Props {
  variant?: 'default' | 'elevated' | 'outlined';
  padding?: 'none' | 'sm' | 'md' | 'lg';
}

const props = withDefaults(defineProps<Props>(), {
  variant: 'default',
  padding: 'md',
});

const cardClasses = computed(() => {
  const base = 'r:24 rel';

  const variants = {
    default: 'bg:white',
    elevated: 'bg:white box-shadow:0|0|20|fade-88',
    outlined: 'bg:white border:1|solid|fade-88',
  };

  const paddings = {
    none: '',
    sm: 'p:12',
    md: 'p:20',
    lg: 'p:32',
  };

  return `${base} ${variants[props.variant]} ${paddings[props.padding]}`;
});
</script>

<template>
  <div :class="cardClasses">
    <slot />
  </div>
</template>
