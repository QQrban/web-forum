import { createSlice, PayloadAction } from "@reduxjs/toolkit";
import { HomePageData } from "../types/types";

const initialState: HomePageData = {
  categories: [
    {
      category: {
        id: 0,
        intro: "",
        parent_id: 0,
        short: "",
        title: "",
      },
      topics: [
        {
          category: {
            id: 0,
            intro: "",
            parent_id: 0,
            short: "",
            title: "",
          },
          comments_number: 0,
          last_comment: {
            author: "",
            avatar: "",
            comment_id: 0,
            content: "",
            created: "",
            dislikes: 0,
            likes: 0,
          },
        },
      ],
    },
  ],
  hotTopics: [],
  latestPosts: [],
  stats: {
    latest_member: "",
    total_members: 0,
    total_online: 0,
    total_posts: 0,
    total_topics: 0,
  },
  topCoders: [],
};

const homePageSlice = createSlice({
  name: "homePageData",
  initialState,
  reducers: {
    setHomePageData: (state, action: PayloadAction<HomePageData>) => {
      return { ...state, ...action.payload };
    },
  },
});

export const { setHomePageData } = homePageSlice.actions;
export default homePageSlice.reducer;
