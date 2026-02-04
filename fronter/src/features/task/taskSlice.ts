import { createSlice, createAsyncThunk, PayloadAction } from "@reduxjs/toolkit";
import {
  fetchTasksClient,
  addTaskClient,
  updateTaskClient,
  deleteTaskClient,
  SuccessResponse,
} from "../../client/taskApi";
import { AddTask, Task } from "../../model/task";

interface TaskState {
  tasks: Task[];
  taskLoading: boolean;
  taskError: string | null;
}

const initialState: TaskState = {
  tasks: [],
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
    return response.value;
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
      .addCase(loadTasks.fulfilled, (state, action: PayloadAction<Task[]>) => {
        state.taskLoading = false;
        state.tasks = action.payload;
      })
      .addCase(loadTasks.rejected, (state, action) => {
        state.taskLoading = false;
        state.taskError = action.payload as string;
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
