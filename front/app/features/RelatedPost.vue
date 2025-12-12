<template>
  <section v-if="post" class="mb:50">
    <h2 class="f:24 f:bold pb:32">
      {{ title }}
    </h2>
    <div class="grid gap:20 grid-cols:1">
      <PostCard :post="post" />
    </div>
  </section>
</template>

<script setup lang="ts">
import type { Post } from '~/modules/posts/types';
import PostCard from '~/modules/posts/components/PostCard.vue';

interface Props {
  post: Post | null;
  title: string;
}

defineProps<Props>();
</script>
