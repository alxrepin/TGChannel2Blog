export interface Post {
  id: string;
  title: string;
  excerpt: string;
  publishedAt: string;
  thumbnail?: string;
  hasImage: boolean;
}

export interface PostsCollection {
  items: Post[];
  total: number;
}
