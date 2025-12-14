import type { Paginate } from '~/common/types';

export type TPostSnippet = {
  id: number;
  url: string;
  title?: string;
  text?: string;
  thumbnail?: string;
  createdAt: string;
  groupId: number;
};

export type TPostSnippetCollection = {
  items: TPostSnippet[];
  paginate: Paginate;
};
