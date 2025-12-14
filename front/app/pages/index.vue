<script setup lang="ts">
import Hero from '~/modules/hero/components/Hero.vue';
import UiContainer from '~/common/components/ui/UiContainer.vue';
import { useChannelData } from '~/common/composables/useChannelData';
import { fetchPosts } from '~/modules/posts/api/posts';
import { computed, useHead, useAsyncData } from '#imports';
import PostCard from '~/modules/posts/components/PostCard.vue';
import type { TPostSnippetCollection } from '~/modules/posts/types';

const channel = useChannelData();
const { data: posts } = useAsyncData<TPostSnippetCollection>(() => fetchPosts(1, 6));
const total = computed(() => posts.value?.paginate.count || 0);

// SEO оптимизация
useHead({
  title: channel.name,
  meta: [
    {
      name: 'description',
      content:
        'Живу в Екатеринбурге, занимаюсь бэкенд разработкой. Тут делюсь опытом, мыслями и заметками из прочитанных книг.',
    },
    { name: 'author', content: 'Ainur Khakimov' },
    { name: 'keywords', content: 'backend, разработка, программирование, блог, заметки' },

    // Open Graph / Facebook
    { property: 'og:type', content: 'website' },
    { property: 'og:title', content: channel.name },
    {
      property: 'og:description',
      content:
        'Живу в Екатеринбурге, занимаюсь бэкенд разработкой. Тут делюсь опытом, мыслями и заметками из прочитанных книг.',
    },
    { property: 'og:image', content: channel.avatar },

    // Twitter
    { name: 'twitter:card', content: 'summary_large_image' },
    { name: 'twitter:title', content: channel.name },
    {
      name: 'twitter:description',
      content:
        'Живу в Екатеринбурге, занимаюсь бэкенд разработкой. Тут делюсь опытом, мыслями и заметками из прочитанных книг.',
    },
    { name: 'twitter:image', content: channel.avatar },
  ],
  link: [{ rel: 'canonical', href: 'https://repin.pw' }],
});
</script>

<template>
  <UiContainer class="flex flex:column ai:center">
    <Hero />
  </UiContainer>
  <UiContainer>
    <h2 class="f:24 f:bold pb:32">
      Посты <span class="pl:6 color:fade-70">{{ total }}</span>
    </h2>
    <div class="grid gap:20 grid-cols:1 grid-cols:2@md grid-cols:3@2xl">
      <PostCard v-for="post in posts.items" :key="post.id" :post="post" />
    </div>
  </UiContainer>
</template>
