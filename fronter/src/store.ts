import { configureStore } from '@reduxjs/toolkit'
import { postReducer } from './modules/postReducer'
import { State } from './model/posts'
import userReducer from './features/user/userSlice'
import memoryReducer from './features/memory/memorySlice'

export type AppState = {
  state: State
}

export const setupStore = configureStore({
  reducer: {
    state: postReducer,
    user: userReducer,
    memory: memoryReducer
  },
})

export type RootState = ReturnType<typeof setupStore.getState>
export type AppDispatch = typeof setupStore.dispatch
