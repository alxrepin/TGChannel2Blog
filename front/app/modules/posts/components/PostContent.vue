<template>
  <article class="f:16 color:fade-10 lh:1.5 box:content">
    <h1 class="f:28 f:bold lh:1.4 mb:24">
      {{ title }}
    </h1>
    <div class="post-content">
      <slot />
    </div>
  </article>
</template>

<script setup lang="ts">
interface Props {
  title: string;
}

defineProps<Props>();
</script>

<style>
.post-content h3 {
  font-size: 24px;
  font-weight: 600;
  padding-bottom: 8px;
  padding-top: 12px;
  line-height: 1.4;
}

.post-content .height\:16px {
  height: 16px;
}
</style>
