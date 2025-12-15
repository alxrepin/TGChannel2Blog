import { ref, onMounted, onUnmounted, nextTick, type Ref } from 'vue';

interface UseIntersectionObserverOptions extends IntersectionObserverInit {
  enabled?: Ref<boolean> | boolean;
}

export function useIntersectionObserver(
  target: Ref<Element | null> | string,
  options: UseIntersectionObserverOptions = {}
) {
  const isIntersecting = ref(false);
  let observer: IntersectionObserver | null = null;

  const { enabled = true, ...observerOptions } = options;

  onMounted(async () => {
    const isEnabled = typeof enabled === 'boolean' ? enabled : enabled.value;

    if (!isEnabled) {
      return;
    }

    await nextTick();

    const element = typeof target === 'string' ? document.querySelector(target) : target.value;

    if (!element) {
      return;
    }

    observer = new IntersectionObserver(entries => {
      isIntersecting.value = entries[0].isIntersecting;
    }, observerOptions);

    observer.observe(element);
  });

  onUnmounted(() => {
    observer?.disconnect();
  });

  return {
    isIntersecting,
  };
}
