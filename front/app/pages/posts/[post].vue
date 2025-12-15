<script setup lang="ts">
import PostNavigation from '~/modules/posts/components/PostNavigation.vue';
import PostContent from '~/modules/posts/components/PostContent.vue';
import UiContainer from '~/common/components/ui/UiContainer.vue';
import { useHead, useRoute, createError, useAsyncData, computed } from '#imports';
import { useChannelData } from '~/modules/channel/composables/use-channel-data';
import { fetchPosts, mapPostSnippetToPost } from '~/modules/posts/api/posts';
import PostCard from '~/modules/posts/components/PostCard.vue';

const route = useRoute();
const channel = useChannelData();
const { data: postsResponse } = useAsyncData('posts', () => fetchPosts(1, 1000));
const posts = computed(() => postsResponse.value?.items.map(mapPostSnippetToPost) || []);

const getPostById = (id: number) => posts.value.find((post) => post.id === id) || null;
const getNextPost = (currentId: number) => {
  const currentIndex = posts.value.findIndex((post) => post.id === currentId);
  if (currentIndex === -1 || currentIndex === posts.value.length - 1) return null;
  return posts.value[currentIndex + 1] || null;
};
const getPreviousPost = (currentId: number) => {
  const currentIndex = posts.value.findIndex((post) => post.id === currentId);
  if (currentIndex === -1 || currentIndex === 0) return null;
  return posts.value[currentIndex - 1] || null;
};

// Получаем ID поста из URL
const postId = parseInt(route.params.post as string);

// Получаем текущий пост
const currentPost = getPostById(postId) || posts.value[0];

if (!currentPost) {
  throw createError({ statusCode: 404, statusMessage: 'Post not found' });
}

// Получаем связанные посты
const nextPost = getNextPost(currentPost.id);
const previousPost = getPreviousPost(currentPost.id);

// SEO оптимизация
useHead({
  title: `${currentPost.title} - ${channel.name}`,
  meta: [
    { name: 'description', content: currentPost.excerpt },
    { name: 'author', content: channel.name },

    // Open Graph
    { property: 'og:type', content: 'article' },
    { property: 'og:title', content: currentPost.title },
    { property: 'og:description', content: currentPost.excerpt },
    { property: 'og:image', content: currentPost.thumbnail || channel.avatar },
    { property: 'article:published_time', content: currentPost.publishedAt },

    // Twitter Card
    { name: 'twitter:card', content: 'summary_large_image' },
    { name: 'twitter:title', content: currentPost.title },
    { name: 'twitter:description', content: currentPost.excerpt },
    { name: 'twitter:image', content: currentPost.thumbnail || channel.avatar },
  ],
  link: [{ rel: 'canonical', href: `https://repin.pw/posts/${postId}` }],
});
</script>

<template>
  <UiContainer size="sm">
    <PostNavigation :published-at="currentPost.publishedAt" />

    <PostContent :title="currentPost.title" class="mb:50">
      <h3 class="f:24 f:semibold pb:8">🚌 Общественный транспорт</h3>
      Сравнивая с Японией, мы не испытали абсолютно никакого дискомфорта при его использовании. Нет
      проблем с навигацией, нет пропадающих указателей в метро, нет лотереи, в какой поезд ты сейчас
      сел — local или rapid. Если мне нужно попасть из точки А в точку Б, я просто открываю
      навигатор (лучше скачать местный — Naver), прокладываю маршрут и… просто еду.
      <div class="height:16px"></div>
      К тому же между пересадками с условного аэроэкспресса из аэропорта в само метро мне не нужно
      доплачивать. Однако до метро Москвы всё ещё не дотягивает.
      <div class="height:16px"></div>
      Единственный минус — автобусы. Во-первых, как и в Японии, они неразумно используют
      пространство внутри. Во-вторых — водители. Если в Москве наши комфортные электробусы плавно,
      считай по рельсам, переносят нашу тушку, то здешние — это какие-то раллийные гонки. Эти парни
      подрезают, перестраиваются через несколько полос сразу и топят тапку в пол со старта. Внутри
      салона физически тяжело держаться за поручни.
      <div class="height:16px"></div>
      <h3 class="f:24 f:semibold pb:8 pt:12 lh:1.4">🚗 Аренда авто</h3>
      Если Япония не поддерживает наши международные права (которые делаются в ГИБДД в течение
      часа), то Южная Корея, наоборот, прекрасно с ними считается.
      <div class="height:16px"></div>
      Мы нашли аренду автомобиля на острове Чеджу, которая не требует депозита и кредитной карты.
      Вот этот микроб на фото — Hyundai Casper — обошёлся нам в 12 тысяч рублей на 4 дня.
      <div class="height:16px"></div>
      Машины здесь, к слову, все чистые, затонированные (тонировка передней полусферы — зло) и часто
      имеют торчащие пупырки на дверях, чтобы не бить их при открытии.
      <div class="height:16px"></div>
      Удивило, что на заправках всего два вида топлива: бензин и дизель. Нет 92, 95 и 100 — просто
      бензин. Цена, кстати, ~98р. за литр.
      <div class="height:16px"></div>
      <h3 class="f:24 f:semibold pb:8 pt:12 lh:1.4">🚥 Вождение и ПДД</h3>
      ПДД, разметка и знаки здесь схожи, но есть и свои отличия. К примеру, светофоры имеют 4
      световых сигнала: 2 зелёных, 1 красный и 1 жёлтый.
      <div class="height:16px"></div>
      Зачем два зелёных? Один для движения прямо и направо, а второй — для поворота налево и
      разворота. Не знаю, как бы это работало на нашем потоке, но отдельная стрелка для поворота
      налево кажется более удобным решением. Все машины пересекают встречку по очереди, а не как у
      нас, выезжая на перекрёсток, из-за чего и ввели вафельную разметку.
      <div class="height:16px"></div>
      Ограничения скорости также имеют свои небольшие особенности: 30 — возле школ, 50 — в
      населённых пунктах, 70 — за городом и 110 — на магистралях. Про порог допустимого превышения я
      так и не понял: интернет говорит, что на всех типах камер всё по-разному. По водителям я не
      смог выявить корреляцию допустимого порога — все под камерой едут по-разному.
      <div class="height:16px"></div>
      Местная культура вождения всё ещё сдержанная по сравнению с нашей, но не такая душная, как в
      той же Японии.
    </PostContent>

    <section v-if="nextPost" class="mb:50">
      <h2 class="f:24 f:bold pb:32">Следующий пост</h2>
      <div class="grid gap:20 grid-cols:1">
        <PostCard :post="nextPost" />
      </div>
    </section>
    <section v-if="previousPost" class="mb:50">
      <h2 class="f:24 f:bold pb:32">Предыдущий пост</h2>
      <div class="grid gap:20 grid-cols:1">
        <PostCard :post="previousPost" />
      </div>
    </section>
  </UiContainer>
</template>
