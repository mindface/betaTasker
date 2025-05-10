import { createSlice, PayloadAction } from '@reduxjs/toolkit'
import { State, Posts } from '../model/posts'

// 初期状態
export const initialState: State = {
  loading: false,
  status: "init",
  posts: [{
    id: 1,
    title: "test01",
    text: "test01",
    name: "test01",
    disc: "test01",
    imgPath: "test01",
  }],
  rePosts: [{
    id: 1,
    title: "test01",
    text: "test01",
    name: "test01",
    disc: "test01",
    imgPath: "test01",
  }],
  searchText: "",
  searchCategpryText: "",
  lastUpdated: 0
}

export const postSlice = createSlice({
  name: 'post',
  initialState,
  reducers: {
    getPostRequest(state) {
      state.status = "Fetching"
      state.posts = []
      state.rePosts = []
      state.searchText = ""
      state.searchCategpryText = ""
      state.lastUpdated = Date.now()
      state.loading = true
    },
    getPostSuccess(state, action: PayloadAction<{ payload: Posts[]; rePayload: Posts[] }>) {
      state.status = "Success"
      state.posts = action.payload.payload
      state.rePosts = action.payload.rePayload
      state.searchText = ""
      state.searchCategpryText = ""
      state.lastUpdated = Date.now()
      state.loading = false
    },
    getPostFailure(state) {
      state.status = "Failure"
      state.posts = []
      state.rePosts = []
      state.searchText = ""
      state.searchCategpryText = ""
      state.lastUpdated = Date.now()
      state.loading = false
    }
  }
})

// Actionsをエクスポート
export const { getPostRequest, getPostSuccess, getPostFailure } = postSlice.actions

// Reducerをエクスポート
export const postReducer = postSlice.reducer