<script setup lang="ts">
import { computed } from '#imports';

interface Props {
  as?: 'div' | 'section' | 'header' | 'main' | 'footer' | 'nav';
  size?: 'sm' | 'md' | 'lg' | 'xl' | 'full';
  padding?: boolean;
}

const props = withDefaults(defineProps<Props>(), {
  as: 'section',
  size: 'lg',
  padding: true,
});

const containerClasses = computed(() => {
  const base = 'mx:auto font-family:inter';

  const sizes = {
    sm: 'max-w:sm',
    md: 'max-w:md',
    lg: 'max-w:lg max-w:3xl@xl',
    xl: 'max-w:xl',
    full: 'w:full',
  };

  const paddingClass = props.padding ? 'p:24 p:48@sm' : '';

  return `${base} ${sizes[props.size]} ${paddingClass}`;
});
</script>

<template>
  <component :is="as" :class="containerClasses">
    <slot />
  </component>
</template>
