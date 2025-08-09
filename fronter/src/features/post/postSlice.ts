import { createSlice, PayloadAction } from '@reduxjs/toolkit';
import { State, Posts } from '../../model/posts'

interface UserState {
  loading: boolean;
  isAuthenticated: boolean;
  error: string | null;
  post: Posts | null;
  rePosts: Posts[];
}

const initialState: UserState = {
  loading: false,
  isAuthenticated: false,
  error: null,
  post: null,
  rePosts: [{
    id: 1,
    title: "test01",
    text: "test01",
    name: "test01",
    disc: "test01",
    imgPath: "test01",
  }],
};

const userSlice = createSlice({
  name: 'posts',
  initialState,
  reducers: {
    postRequest: (state) => {
      state.loading = true;
      state.error = null;
    },
    postSuccess: (state, action: PayloadAction<Posts>) => {
      state.loading = false;
      state.isAuthenticated = true;
      state.post = action.payload;
    },
    postFailure: (state, action: PayloadAction<string>) => {
      state.loading = false;
      state.isAuthenticated = false;
      state.error = action.payload;
    },
  },
});

export const { postRequest, postSuccess, postFailure } = userSlice.actions;
export default userSlice.reducer;
