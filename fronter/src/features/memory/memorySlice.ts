import { createSlice, createAsyncThunk, PayloadAction } from '@reduxjs/toolkit';
import { fetchMemoriesService, fetchMemoryService, addMemoryService, updateMemoryService, deleteMemoryService } from '../../services/memoryApi';
import { AddMemory, Memory } from '../../model/memory';

interface MemoryState {
  memories: Memory[];
  memoryItem: Memory;
  memoryLoading: boolean;
  memoryError: string | null;
}

const initialState: MemoryState = {
  memories: [],
  memoryItem: {
    id: 0,
    user_id: 0,
    title: 'no title',
    notes: '',
    created_at: '',
    updated_at: '',
    source_type: '',
    author: '',
    factor: '',
    process: '',
    evaluation_axis: '',
    information_amount: '',
    tags: '',
    read_status: '',
    read_date: '',
  },
  memoryLoading: false,
  memoryError: null,
}

export const loadMemories = createAsyncThunk(
  'memory/loadMemories',
  async (_, { rejectWithValue }) => {
    const response = await fetchMemoriesService();
    if ('error' in response) {
      return rejectWithValue(response.error);
    }
    return response.value ?? [];
  }
)


export const getMemory = createAsyncThunk(
  'memory/getMemory',
  async (memoryId: number, { rejectWithValue }) => {
    const response = await fetchMemoryService(memoryId);
    if ('error' in response) {
      return rejectWithValue(response.error);
    }
    return response.value;
  }
)


export const createMemory = createAsyncThunk(
  'memory/createMemory',
  async (memoryData: AddMemory, { rejectWithValue }) => {
    const response = await addMemoryService(memoryData);
    if ('error' in response) {
      return rejectWithValue(response.error);
    }
    return response.value;
  }
)

export const updateMemory = createAsyncThunk(
  'memory/updateMemory',
  async (memoryData: Memory, { rejectWithValue }) => {
    const response = await updateMemoryService(memoryData);
    if ('error' in response) {
      return rejectWithValue(response.error);
    }
    return response.value;
  }
)

export const removeMemory = createAsyncThunk(
  'memory/removeMemory',
  async (id: number, { rejectWithValue }) => {
    const response = await deleteMemoryService(id.toString());
    if ('error' in response) {
      return rejectWithValue(response.error);
    }
    return id;
  }
);

const memorySlice = createSlice({
  name: 'memory',
  initialState,
  reducers: {
    clearError: (state) => {
      state.memoryError = null;
    },
  },
  extraReducers: (builder) => {
    builder
      .addCase(loadMemories.pending, (state) => {
        state.memoryLoading = true;
        state.memoryError = null;
      })
      .addCase(loadMemories.fulfilled, (state, action: PayloadAction<Memory[]>) => {
        console.log(action.payload)
        state.memoryLoading = false;
        state.memories = action.payload;
      })
      .addCase(loadMemories.rejected, (state, action) => {
        state.memoryLoading = false;
        state.memoryError = action.payload as string;
      })
      .addCase(getMemory.pending, (state) => {
        state.memoryLoading = true;
        state.memoryError = null;
      })
      .addCase(getMemory.fulfilled, (state, action: PayloadAction<Memory>) => {
        state.memoryLoading = false;
        state.memoryItem = action.payload;
      })
      .addCase(getMemory.rejected, (state, action) => {
        state.memoryLoading = false;
        state.memoryError = action.payload as string;
      })
      .addCase(createMemory.pending, (state) => {
        state.memoryLoading = true;
        state.memoryError = null;
      })
      .addCase(createMemory.fulfilled, (state, action: PayloadAction<Memory>) => {
        state.memoryLoading = false;
        state.memories.push(action.payload);
      })
      .addCase(createMemory.rejected, (state, action) => {
        state.memoryLoading = false;
        state.memoryError = action.payload as string;
      })
      .addCase(removeMemory.fulfilled, (state, action: PayloadAction<number>) => {
        state.memories = state.memories.filter(memory => memory.id !== action.payload);
      });
  },
});

export const { clearError } = memorySlice.actions;
export default memorySlice.reducer;
