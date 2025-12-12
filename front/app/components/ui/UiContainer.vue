<script setup lang="ts">
interface Props {
  size?: 'sm' | 'md' | 'lg' | 'xl' | 'full';
  padding?: boolean;
}

const props = withDefaults(defineProps<Props>(), {
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

  const paddingClass = props.padding ? 'p:48' : '';

  return `${base} ${sizes[props.size]} ${paddingClass}`;
});
</script>

<template>
  <div :class="containerClasses">
    <slot />
  </div>
</template>
