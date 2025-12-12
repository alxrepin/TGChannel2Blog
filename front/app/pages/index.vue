<script setup lang="ts">
import AppHeader from '~/components/layout/AppHeader.vue';
import ChannelAvatar from '~/modules/hero/components/ChannelAvatar.vue';
import ChannelInfo from '~/modules/hero/components/ChannelInfo.vue';
import ChannelStats from '~/modules/hero/components/ChannelStats.vue';
import SubscribeButton from '~/features/SubscribeButton.vue';
import PostGrid from '~/features/PostGrid.vue';
import AppFooter from '~/components/layout/AppFooter.vue';
import UiContainer from '~/components/ui/UiContainer.vue';

const { channel } = useChannelData();
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
  <UiContainer>
    <AppHeader>
      <ChannelAvatar :src="channel.avatar" :alt="channel.name" class="mb:20" />
      <ChannelInfo :name="channel.name" :description="channel.description" />
      <div class="pb:20">
        <ChannelStats :subscriber-count="channel.subscriberCount" />
      </div>
      <SubscribeButton :telegram-handle="channel.telegramHandle" />
    </AppHeader>

    <PostGrid :posts="posts" :total="postsTotal" />

    <AppFooter :text="channel.name" />
  </UiContainer>
</template>
