import { configureStore } from '@reduxjs/toolkit'
import { State } from './model/posts'
import userReducer from './features/user/userSlice'
import memoryReducer from './features/memory/memorySlice'
import taskReducer from './features/task/taskSlice'
import assessmentReducer from './features/assessment/assessmentSlice'
import learningReducer from './features/learningData/learningDataSlice'
import memoryAidReducer from './features/memoryAid/memoryAidSlice';
import { postReducer } from './modules/postReducer';
import { TypedUseSelectorHook, useDispatch, useSelector } from 'react-redux';

export type AppState = {
  state: State
}

export const setupStore = configureStore({
  reducer: {
    user: userReducer,
    post: postReducer,
    memory: memoryReducer,
    task: taskReducer,
    assessment: assessmentReducer,
    learning: learningReducer,
    memoryAid: memoryAidReducer,
  },
})

export type RootState = ReturnType<typeof setupStore.getState>
export type AppDispatch = typeof setupStore.dispatch

export const useAppDispatch = () => useDispatch<AppDispatch>();
export const useAppSelector: TypedUseSelectorHook<RootState> = useSelector;
