<script setup lang="ts">
interface Props {
  variant?: 'primary' | 'secondary' | 'ghost'
  size?: 'sm' | 'md' | 'lg'
  tag?: 'button' | 'a'
  href?: string
  type?: 'button' | 'submit' | 'reset'
  disabled?: boolean
}

const props = withDefaults(defineProps<Props>(), {
  variant: 'primary',
  size: 'md',
  tag: 'button',
  type: 'button',
  disabled: false,
})

const buttonClasses = computed(() => {
  const base = 'font:medium r:12 flex gap:10 ai:center jc:center transition cursor:pointer'

  const variants = {
    primary: 'bg:sky-2 font:white hover:bg:sky-3',
    secondary: 'bg:fade-94 font:fade-10 hover:bg:fade-88',
    ghost: 'bg:transparent font:fade-10 hover:bg:fade-94',
  }

  const sizes = {
    sm: 'f:14 px:16 pl:10 py:12',
    md: 'f:16 px:18 pl:12 py:14',
    lg: 'f:20 px:20 pl:14 py:16',
  }

  return `${base} ${variants[props.variant]} ${sizes[props.size]}`
})
</script>

<template>
  <component
    :is="tag"
    :href="href"
    :type="type"
    :class="buttonClasses"
    :disabled="disabled"
  >
    <slot />
  </component>
</template>
