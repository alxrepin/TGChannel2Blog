import { computed, useRoute } from '#imports';

export function usePageData() {
  const route = useRoute();
  const isMainPage = computed(() => route.path === '/');

  return {
    isMainPage,
  };
}
