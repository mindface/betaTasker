import { configureStore } from '@reduxjs/toolkit'
import { postReducer } from './modules/postReducer'
import { State } from './model/posts'

export type AppState = {
  state: State
}

export const setupStore = configureStore({
  reducer: {
    state: postReducer,
  },
})

export type RootState = ReturnType<typeof setupStore.getState>
export type AppDispatch = typeof setupStore.dispatch
