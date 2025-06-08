import { configureStore } from '@reduxjs/toolkit'
import { postReducer } from './modules/postReducer'
import { userReducer } from './modules/userReducer'
import { State } from './model/posts'

export type AppState = {
  state: State
}

export const setupStore = configureStore({
  reducer: {
    state: postReducer,
    user: userReducer,
  },
})

export type RootState = ReturnType<typeof setupStore.getState>
export type AppDispatch = typeof setupStore.dispatch
