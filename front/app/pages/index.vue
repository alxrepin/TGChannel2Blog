<script setup lang="ts">
import PostGrid from '~/features/PostGrid.vue';

import Hero from '~/modules/hero/components/Hero.vue';
import UiContainer from '~/common/components/ui/UiContainer.vue';
import { useChannelData } from '~/common/composables/useChannelData';
import { usePostsData } from '~/common/composables/usePostsData';
import { useHead } from '#imports';

const channel = useChannelData();
const { posts } = usePostsData();

const postsTotal = posts.length;

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
    <PostGrid :posts="posts" :total="postsTotal" />
  </UiContainer>
</template>
