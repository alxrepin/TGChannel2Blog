<script setup lang="ts">
import type { Post } from '~/modules/posts/types'
import PostCard from "~/modules/posts/components/PostCard.vue";

interface Props {
  posts: Post[]
  total: number
}

defineProps<Props>()
</script>

<template>
  <section class="mb:50">
    <h2 class="f:24 f:bold pb:32">
      Посты <span class="pl:6 color:fade-70">{{ total }}</span>
    </h2>
    <div class="grid gap:20 grid-cols:1 grid-cols:2@md grid-cols:3@2xl">
      <PostCard 
        v-for="post in posts" 
        :key="post.id" 
        :post="post" 
      />
    </div>
  </section>
</template>
