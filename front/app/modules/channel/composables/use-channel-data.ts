import type { TChannel } from '../types';
import { useNuxtData } from '#app';

export function useChannelData() {
  return useNuxtData<TChannel>('channel');
}
