import type { TPostSnippetCollection } from '../types';
import type { Paginate } from '~/common/types';

interface IPost {
  id: number;
  title?: string;
  url: string;
  text?: string;
  created_at: string;
}

interface IPostsResponse {
  items: Post[];
  paginate: Paginate;
}

export const fetchPosts = async (
  page: number = 1,
  limit: number = 100
): Promise<TPostSnippetCollection> => {
  const response = await $fetch<IPostsResponse>('/api/v1/posts', {
    baseURL: 'http://localhost:8080',
    query: { page, limit },
  });

  return {
    items: response.items.map((i: IPost) => ({
      id: i.id,
      url: i.url,
      title: i?.title,
      text: i?.text,
      thumbnail: null,
      createdAt: i.created_at,
    })),
    paginate: response.paginate,
  };
};
