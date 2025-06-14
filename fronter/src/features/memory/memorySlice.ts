import { createSlice, createAsyncThunk, PayloadAction } from '@reduxjs/toolkit';
import { fetchMemoriesService, addMemoryService, updateMemoryService, deleteMemoryService } from '../../services/memoryApi';
import { AddMemory, Memory } from '../../model/memory';

interface MemoryState {
  memories: Memory[];
  loading: boolean;
  error: string | null;
}

const initialState: MemoryState = {
  memories: [],
  loading: false,
  error: null,
}

export const loadMemories = createAsyncThunk(
  'memory/loadMemories',
  async (_, { rejectWithValue }) => {
    try {
      const response = await fetchMemoriesService();
      if (response.error) {
        return rejectWithValue(response.error);
      }
      return response.memories;
    } catch (error: any) {
      return rejectWithValue(error.message);
    }
  }
)

export const createMemory = createAsyncThunk(
  'memory/createMemory',
  async (memoryData: AddMemory, { rejectWithValue }) => {
    try {
      const response = await addMemoryService(memoryData);
      if (response.error) {
        return rejectWithValue(response.error);
      }
      return response;
    } catch (error: any) {
      return rejectWithValue(error.message);
    }
  }
)

export const updateMemory = createAsyncThunk(
  'memory/updateMemory',
  async (memoryData: Memory, { rejectWithValue }) => {
    try {
      const response = await updateMemoryService(memoryData);
      if (response.error) {
        return rejectWithValue(response.error);
      }
      return response;
    } catch (error: any) {
      return rejectWithValue(error.message);
    }
  }
)

export const removeMemory = createAsyncThunk(
  'memory/removeMemory',
  async (id: number, { rejectWithValue }) => {
    try {
      const response = await deleteMemoryService(id.toString());
      if (response.error) {
        return rejectWithValue(response.error);
      }
      return id;
    } catch (error: any) {
      return rejectWithValue(error.message);
    }
  }
);

const memorySlice = createSlice({
  name: 'memory',
  initialState,
  reducers: {
    clearError: (state) => {
      state.error = null;
    },
  },
  extraReducers: (builder) => {
    builder
      .addCase(loadMemories.pending, (state) => {
        state.loading = true;
        state.error = null;
      })
      .addCase(loadMemories.fulfilled, (state, action: PayloadAction<Memory[]>) => {
        state.loading = false;
        state.memories = action.payload;
      })
      .addCase(loadMemories.rejected, (state, action) => {
        state.loading = false;
        state.error = action.payload as string;
      })
      .addCase(createMemory.pending, (state) => {
        state.loading = true;
        state.error = null;
      })
      .addCase(createMemory.fulfilled, (state, action: PayloadAction<Memory>) => {
        state.loading = false;
        state.memories.push(action.payload);
      })
      .addCase(createMemory.rejected, (state, action) => {
        state.loading = false;
        state.error = action.payload as string;
      })
      .addCase(removeMemory.fulfilled, (state, action: PayloadAction<number>) => {
        state.memories = state.memories.filter(memory => memory.id !== action.payload);
      });
  },
});

export const { clearError } = memorySlice.actions;
export default memorySlice.reducer;
