import { createSlice, createAsyncThunk, PayloadAction } from "@reduxjs/toolkit";
import {
  fetchTasksClient,
  addTaskClient,
  updateTaskClient,
  deleteTaskClient,
  SuccessResponse,
  getTasksLimitClient,
} from "../../client/taskApi";
import { AddTask, Task } from "../../model/task";
import { LimitResponse } from "@/model/respose";

interface TaskState {
  tasks: Task[];

  tasksPage: number;
  tasksLimit: number;
  tasksTotal: number;
  tasksTotalPages: number;

  taskLoading: boolean;
  taskError: Error | null;
}

const initialState: TaskState = {
  tasks: [],

  tasksPage: 1,
  tasksLimit: 1,
  tasksTotal: 1,
  tasksTotalPages: 1,

  taskLoading: false,
  taskError: null,
};

export const loadTasks = createAsyncThunk(
  "task/loadTasks",
  async (_, { rejectWithValue }) => {
    const response = await fetchTasksClient();
    if ("error" in response) {
      return rejectWithValue(response.error);
    }
    return { tasks: response.tasks, meta: response.meta };
  },
);

export const getTasksLimit = createAsyncThunk(
  "task/getTasksLimit",
  async (payload: { page: number; limit: number }, { rejectWithValue }) => {
    const response = await getTasksLimitClient(
      payload.page,
      payload.limit,
    );
    if ("error" in response) {
      return rejectWithValue({
        message: response.error.message,
        name: response.error.name,
      });
    }
    return { tasks: response.tasks, meta: response.meta };
  },
);

export const createTask = createAsyncThunk(
  "task/createTask",
  async (taskData: AddTask, { rejectWithValue }) => {
    const response = await addTaskClient(taskData);
    if ("error" in response) {
      return rejectWithValue(response.error);
    }
    return response.value;
  },
);

export const updateTask = createAsyncThunk(
  "task/updateTask",
  async (taskData: Task, { rejectWithValue }) => {
    const response = await updateTaskClient(taskData);
    if ("error" in response) {
      return rejectWithValue(response.error);
    }
    return response.value;
  },
);

export const removeTask = createAsyncThunk(
  "task/removeTask",
  async (id: number, { rejectWithValue }) => {
    const response = await deleteTaskClient(id);
    if ("error" in response) {
      return rejectWithValue(response.error);
    }
    return { id };
  },
);

const taskSlice = createSlice({
  name: "task",
  initialState,
  reducers: {},
  extraReducers: (builder) => {
    builder
      .addCase(loadTasks.pending, (state) => {
        state.taskLoading = true;
        state.taskError = null;
      })
      .addCase(loadTasks.fulfilled, (state, action: PayloadAction<LimitResponse<Task,"tasks">>) => {
        state.taskLoading = false;
        state.tasks = action.payload.tasks;
      })
      .addCase(loadTasks.rejected, (state, action) => {
        state.taskLoading = false;
        state.taskError = action.payload as Error;
      })
      .addCase(getTasksLimit.pending, (state) => {
        state.taskLoading = true;
        state.taskError = null;
      })
      .addCase(getTasksLimit.fulfilled, (state, action: PayloadAction<LimitResponse<Task,"tasks">>) => {
        state.taskLoading = false;
        state.tasks = action.payload.tasks;
        state.tasksLimit = action.payload.meta.limit;
        state.tasksTotal = action.payload.meta.total;
        state.tasksTotalPages = action.payload.meta.total_pages;
      })
      .addCase(getTasksLimit.rejected, (state, action) => {
        console.error("----------------");
        console.error("Failed to load tasks:", action.payload);
        state.taskLoading = false;
        state.taskError = action.payload as Error;
      })
      .addCase(createTask.fulfilled, (state, action: PayloadAction<Task>) => {
        state.tasks.push(action.payload);
      })
      .addCase(updateTask.fulfilled, (state, action: PayloadAction<Task>) => {
        const idx = state.tasks.findIndex((t) => t.id === action.payload.id);
        if (idx !== -1) state.tasks[idx] = action.payload;
      })
      .addCase(
        removeTask.fulfilled,
        (state, action: PayloadAction<{ id: number }>) => {
          state.tasks = state.tasks.filter((t) => t.id !== action.payload.id);
        },
      );
  },
});

export default taskSlice.reducer;
