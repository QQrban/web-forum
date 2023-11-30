interface CategoryMain {
  id: number;
  intro: string;
  parent_id: number;
  short: string;
  title: string;
}

interface StatsAll {
  latest_member: string;
  total_members: number;
  total_online: number;
  total_posts: number;
  total_topics: number;
}

interface LastCommentMain {
  author: string;
  avatar: string;
  comment_id: number;
  content: string;
  created: string;
  dislikes: number;
  likes: number;
}

interface AuthorMain {
  avatar: string;
  username: string;
}

interface PostMain {
  created: string;
  likes: number;
  dislikes: number;
  id: number;
  title: string;
}

export interface HomePageData {
  categories: {
    category: CategoryMain;
    topics: [
      {
        category: CategoryMain;
        comments_number: number;
        last_comment: LastCommentMain;
      }
    ];
  }[];
  hotTopics: {
    topic: {
      author: AuthorMain;
      post: PostMain;
    };
  }[];
  latestPosts: {
    author: AuthorMain;
    post: PostMain;
  }[];
  stats: StatsAll;
  topCoders: {
    author: AuthorMain;
  }[];
}
