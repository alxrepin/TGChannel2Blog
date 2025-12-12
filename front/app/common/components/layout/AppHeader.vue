<script setup lang="ts">
import NavBar from '~/modules/navbar/components/NavBar.vue';
import UiContainer from '~/common/components/ui/UiContainer.vue';
import { computed, ref, onMounted, onUnmounted, nextTick, useRoute } from '#imports';

const route = useRoute();
const isHome = computed(() => route.path === '/');
const isHeroVisible = ref(true);
const isShow = computed(() => {
  return !isHome.value || !isHeroVisible.value;
});

let observer: IntersectionObserver | null = null;

onMounted(() => {
  if (!isHome.value) {
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
    v-if="isShow"
    as="header"
    class="flex flex:column fixed w:full z:10 contain:layout|style"
  >
    <NavBar title="R — Repin" image="/images/avatar.jpeg" url="@allrpn" class="absolute" />
  </UiContainer>
</template>
