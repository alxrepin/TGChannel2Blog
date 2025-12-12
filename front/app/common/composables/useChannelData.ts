import type { ChannelInfo } from '~/modules/hero/types';

export function useChannelData(): ChannelInfo {
  return {
    name: 'R — Repin',
    description:
      'Живу в Екатеринбурге, занимаюсь бэкенд разработкой. Тут делюсь опытом, мыслями и заметками из прочитанных книг. Написать мне в личку - @ainr_c Мой блог - https://ainur-khakimov.ru',
    avatar: '/images/avatar.jpeg',
    subscriberCount: 35,
    telegramHandle: 'ainr_c',
    blogUrl: 'https://ainur-khakimov.ru',
  } as ChannelInfo;
}
