<script setup lang="ts">
interface Props {
  variant?: 'default' | 'elevated' | 'outlined'
  padding?: 'none' | 'sm' | 'md' | 'lg'
}

const props = withDefaults(defineProps<Props>(), {
  variant: 'default',
  padding: 'md',
})

const cardClasses = computed(() => {
  const base = 'r:24'
  const hover = 'outline:1|solid|fade-80:hover box-shadow:0|0|20|10|fade-88:hover ~ease-in|300ms ~property:all'

  const variants = {
    default: 'bg:white',
    elevated: 'bg:white box-shadow:0|0|20|fade-88',
    outlined: 'bg:white border:1|solid|fade-88',
  }

  const paddings = {
    none: '',
    sm: 'p:12',
    md: 'p:20',
    lg: 'p:32',
  }

  return `${base} ${variants[props.variant]} ${paddings[props.padding]} ${hover}`
})
</script>

<template>
  <div :class="cardClasses">
    <slot />
  </div>
</template>
