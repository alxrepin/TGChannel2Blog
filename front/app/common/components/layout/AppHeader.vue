<script setup lang="ts">
import NavBar from '~/modules/navbar/components/NavBar.vue';
import UiContainer from '~/common/components/ui/UiContainer.vue';
import { computed, ref, onMounted, onUnmounted, nextTick, useRoute } from '#imports';

const route = useRoute();
const isMainPage = computed(() => route.path === '/');
const isHeroVisible = ref(true);
const isShow = computed(() => {
  return !isMainPage.value || !isHeroVisible.value;
});

let observer: IntersectionObserver | null = null;

onMounted(() => {
  if (!isMainPage.value) {
    return;
  }

  nextTick(() => {
    const element = document.querySelector('#hero');

    if (!element) {
      return;
    }

    observer = new IntersectionObserver(
      entries => {
        isHeroVisible.value = entries[0].isIntersecting;
      },
      { threshold: 0 }
    );

    observer.observe(element);
  });
});

onUnmounted(() => observer?.disconnect());
</script>

<template>
  <UiContainer
    v-show="isShow"
    as="header"
    class="fixed w:full z:10 contain:layout|style"
    :padding="false"
  >
    <NavBar title="R — Repin" image="/images/avatar.jpeg" url="@allrpn" />
  </UiContainer>
</template>
