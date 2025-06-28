import { createSlice, PayloadAction } from '@reduxjs/toolkit'
import { UserInfo } from "../model/user";

export interface UserState {
  loading: boolean;
  isAuthenticated: boolean;
  token: string | null;
  user: UserInfo | null;
  error: string | null;
}

const initialState: UserState = {
  loading: false,
  isAuthenticated: false,
  token: null,
  user: null,
  error: null,
}

export const userSlice = createSlice({
  name: 'user',
  initialState,
  reducers: {
    loginRequest(state) {
      state.loading = true;
      state.error = null;
    },
    loginSuccess(state, action: PayloadAction<{ token: string; user: UserInfo }>) {
      state.loading = false;
      state.isAuthenticated = true;
      state.token = action.payload.token;
      state.user = action.payload.user;
      state.error = null;
    },
    loginFailure(state, action: PayloadAction<string>) {
      state.loading = false;
      state.isAuthenticated = false;
      state.token = null;
      state.user = null;
      state.error = action.payload;
    },
    logout(state) {
      state.loading = false;
      state.isAuthenticated = false;
      state.token = null;
      state.user = null;
      state.error = null;
    },
  },
});

export const { loginRequest, loginSuccess, loginFailure, logout } = userSlice.actions;
export const userReducer = userSlice.reducer;
