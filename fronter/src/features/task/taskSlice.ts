import { createSlice, createAsyncThunk, PayloadAction } from '@reduxjs/toolkit';
import { fetchTasksService, addTaskService, updateTaskService, deleteTaskService, SuccessResponse } from '../../services/taskApi';
import { AddTask, Task } from '../../model/task';
import { ErrorCode } from '../../errors/errorCodes';

interface TaskState {
  tasks: Task[];
  taskLoading: boolean;
  taskError: string | null;
}

const initialState: TaskState = {
  tasks: [],
  taskLoading: false,
  taskError: null,
}

export const loadTasks = createAsyncThunk(
  'task/loadTasks',
  async (_, { rejectWithValue }) => {
    try {
      const response = await fetchTasksService();
      if ('error' in response) {
        return rejectWithValue(response.error);
      }
      console.log(response)
      if ('data' in response) {
        return response.data || [];
      }
      return response as Task[];
    } catch (error: any) {
      return rejectWithValue(error.message);
    }
  }
)

export const createTask = createAsyncThunk(
  'task/createTask',
  async (taskData: AddTask, { rejectWithValue }) => {
    try {
      const response = await addTaskService(taskData);
      if ('error' in response) {
        return rejectWithValue(response.error);
      }
      // Check if response is SuccessResponse<Task> or Task
      if ('data' in response) {
        return response.data as Task;
      }
      return response as Task;
    } catch (error: any) {
      return rejectWithValue(error.message);
    }
  }
)

export const updateTask = createAsyncThunk(
  'task/updateTask',
  async (taskData: Task, { rejectWithValue }) => {
    try {
      const response = await updateTaskService(taskData);
      if ('error' in response) {
        return rejectWithValue(response.error);
      }
      // Check if response is SuccessResponse<Task> or Task
      if ('data' in response) {
        return response.data as Task;
      }
      return response as Task;
    } catch (error: any) {
      return rejectWithValue(error.message);
    }
  }
)

export const removeTask = createAsyncThunk(
  'task/removeTask',
  async (id: number, { rejectWithValue }) => {
    try {
      const response = await deleteTaskService(String(id));
      if ('error' in response) {
        return rejectWithValue(response.error);
      }
      return { id };
    } catch (error: any) {
      return rejectWithValue(error.message);
    }
  }
)

const taskSlice = createSlice({
  name: 'task',
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
        console.log('Tasks loaded:', action.payload);
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
        const idx = state.tasks.findIndex(t => t.id === action.payload.id);
        if (idx !== -1) state.tasks[idx] = action.payload;
      })
      .addCase(removeTask.fulfilled, (state, action: PayloadAction<{ id: number }>) => {
        state.tasks = state.tasks.filter(t => t.id !== action.payload.id);
      });
  },
});

export default taskSlice.reducer;
