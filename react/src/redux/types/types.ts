interface Stats {
  latest_member: string;
  total_members: number;
  total_online: number;
  total_posts: number;
  total_topics: number;
}

interface User {
  current_id: number;
  username: string;
  avatar: string;
  days: number;
  posts: number;
  comments?: number;
}

interface Post {
  id: number;
  title: string;
  content_id: number;
  user_id: number;
  category_id: number;
  created?: string;
  updated?: string;
  dislikes?: number;
  likes?: number;
  tags?: string;
}

interface Topic {
  category_id: number;
  content_id: number;
  created: string;
  dislikes: number;
  id: number;
  likes: number;
  tags: string;
  title: string;
  updated: string;
  user_id: number;
}

interface Category {
  id: number;
  intro: string;
  parent_id: number;
  short: string;
  title: string;
  topics: Topic[];
}

export interface HomePageData {
  categories: Category[];
  stats: Stats;
  topCoders: User[];
  latestPosts: {
    post: Post;
    created_raw: string;
    author: User;
  }[];
  hotTopics: {
    post: Topic;
    created_raw: string;
    author: User;
  }[];
}
