import { configureStore } from '@reduxjs/toolkit'
import { State } from './model/posts'
import userReducer from './features/user/userSlice'
import memoryReducer from './features/memory/memorySlice'
import taskReducer from './features/task/taskSlice'
import assessmentReducer from './features/assessment/assessmentSlice'

export type AppState = {
  state: State
}

export const setupStore = configureStore({
  reducer: {
    user: userReducer,
    memory: memoryReducer,
    task: taskReducer,
    assessment: assessmentReducer,
  },
})

export type RootState = ReturnType<typeof setupStore.getState>
export type AppDispatch = typeof setupStore.dispatch
