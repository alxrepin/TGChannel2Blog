import type { TChannel } from '~/modules/channel/types';

interface IChannelResponse {
  id: number;
  name: string;
  title: string;
  description?: string;
  avatar?: string;
  subscriptions: number;
}

export const fetchChannel = async (): Promise<TChannel> => {
  const data = await $fetch<IChannelResponse>('/api/v1/channel', {
    baseURL: 'http://localhost:8080',
  });

  return {
    name: data.name,
    title: data.title,
    description: data.description,
    avatar: data.avatar,
    subscriptions: data.subscriptions,
    url: 'https://t.me/' + data.name,
  };
};
