import { configureStore } from "@reduxjs/toolkit";
import { State } from "./model/posts";
import userReducer from "./features/user/userSlice";
import memoryReducer from "./features/memory/memorySlice";
import taskReducer from "./features/task/taskSlice";
import postReducer from "./features/post/postSlice";
import assessmentReducer from "./features/assessment/assessmentSlice";
import learningReducer from "./features/learning_data/learningDataSlice";
import memoryAidReducer from "./features/memoryAid/memoryAidSlice";
import heuristicsReducer from "./features/heuristics/heuristicsSlice";
import knowledgePatternReducer from "./features/knowledge_pattern/knowledgePatternSlice";
import languageOptimizationReducer from "./features/language_optimization/languageOptimizationSlice";
import teachingFreeControlReducer from "./features/teaching_free_control/teachingFreeControlSlice";
import learningDataReducer from "./features/learning_data/learningDataSlice";
import qualitativeLabelReducer from "./features/qualitative_label/qualitativeLabelSlice";
import processOptimizationReducer from "./features/process_optimization/processOptimizationSlice";
import { TypedUseSelectorHook, useDispatch, useSelector } from "react-redux";

export type AppState = {
  state: State;
};

export const setupStore = configureStore({
  reducer: {
    user: userReducer,
    memory: memoryReducer,
    post: postReducer,
    task: taskReducer,
    assessment: assessmentReducer,
    learning: learningReducer,
    memoryAid: memoryAidReducer,
    heuristics: heuristicsReducer,
    knowledgePattern: knowledgePatternReducer,
    languageOptimization: languageOptimizationReducer,
    teachingFreeControl: teachingFreeControlReducer,
    learningData: learningDataReducer,
    qualitativeLabel: qualitativeLabelReducer,
    processOptimization: processOptimizationReducer,
  },
});

export type RootState = ReturnType<typeof setupStore.getState>;
export type AppDispatch = typeof setupStore.dispatch;

export const useAppDispatch = () => useDispatch<AppDispatch>();
export const useAppSelector: TypedUseSelectorHook<RootState> = useSelector;
