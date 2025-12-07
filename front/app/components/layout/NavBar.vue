<template>
  <nav class="flex jc:space-between ai:center height:40px mb:32">
    <NuxtLink to="/" class="flex ai:center gap:10 text-decoration:none">
      <NuxtImg
        :src="avatarSrc"
        :alt="channelName"
        :width="40"
        :height="40"
        class="rounded"
        style="border-radius: 50%"
        loading="lazy"
      />
      <div class="f:20 f:bold f:fade-10">
        {{ channelName }}
      </div>
    </NuxtLink>
    
    <UiButton
      tag="a"
      :href="telegramLink"
      target="_blank"
      rel="noopener noreferrer"
      size="sm"
    >
      <Icon name="mdi:telegram" />
      Подписаться
    </UiButton>
  </nav>
</template>

<script setup lang="ts">
interface Props {
  channelName: string
  avatarSrc: string
  telegramHandle: string
}

const props = defineProps<Props>()

const telegramLink = computed(() => `https://t.me/${props.telegramHandle}`)
</script>