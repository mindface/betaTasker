import { createSlice, createAsyncThunk, PayloadAction } from "@reduxjs/toolkit";
import {
  fetchMemoriesClient,
  getMemoryClient,

  addMemoryClient,
  updateMemoryClient,
  deleteMemoryClient,
  getMemoriesLimitClient,
} from "@/client/memoryApi";
import { AddMemory, Memory } from "@/model/memory";
import { LimitResponse } from "@/model/respose";

interface MemoryState {
  memories: Memory[];
  memoryItem: Memory;

  memoriesPage: number;
  memoriesLimit: number;
  memoriesTotal: number;
  memoriesTotalPages: number;

  memoryLoading: boolean;
  memoryError: string | null;
}

const initialState: MemoryState = {
  memories: [],
  memoryItem: {
    id: 0,
    user_id: 0,
    title: "no title",
    notes: "",
    created_at: "",
    updated_at: "",
    source_type: "",
    author: "",
    factor: "",
    process: "",
    evaluation_axis: "",
    information_amount: "",
    tags: "",
    read_status: "",
    read_date: "",
  },

  memoriesPage: 1,
  memoriesLimit: 20,
  memoriesTotal: 0,
  memoriesTotalPages: 0,

  memoryLoading: false,
  memoryError: null,
};

export const loadMemories = createAsyncThunk(
  "memory/loadMemories",
  async (_, { rejectWithValue }) => {
    const response = await fetchMemoriesClient();
    if ("error" in response) {
      return rejectWithValue(response.error);
    }
    return { memories: response.memories, meta: response.meta };
  },
);


export const getMemory = createAsyncThunk(
  "memory/getMemory",
  async (memoryId: number, { rejectWithValue }) => {
    const response = await getMemoryClient(memoryId);
    if ("error" in response) {
      return rejectWithValue(response.error);
    }
    return response.value;
  },
);

export const getMemoriesLimit = createAsyncThunk(
  "memory/getMemoriesLimit",
  async (payload: { page: number; limit: number }, { rejectWithValue }) => {
    const response = await getMemoriesLimitClient(payload.page, payload.limit);
    if ("error" in response) {
      return rejectWithValue(response.error);
    }
    return { memories: response.memories, meta: response.meta };
  },
);

export const createMemory = createAsyncThunk(
  "memory/createMemory",
  async (memoryData: AddMemory, { rejectWithValue }) => {
    const response = await addMemoryClient(memoryData);
    if ("error" in response) {
      return rejectWithValue(response.error);
    }
    return response.value;
  },
);

export const updateMemory = createAsyncThunk(
  "memory/updateMemory",
  async (memoryData: Memory, { rejectWithValue }) => {
    const response = await updateMemoryClient(memoryData);
    if ("error" in response) {
      return rejectWithValue(response.error);
    }
    return response.value;
  },
);

export const removeMemory = createAsyncThunk(
  "memory/removeMemory",
  async (id: number, { rejectWithValue }) => {
    const response = await deleteMemoryClient(id.toString());
    if ("error" in response) {
      return rejectWithValue(response.error);
    }
    return id;
  },
);

const memorySlice = createSlice({
  name: "memory",
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
      .addCase(
        loadMemories.fulfilled,
        (state, action: PayloadAction<LimitResponse<Memory, "memories">>) => {
          state.memoryLoading = false;
          state.memories = action.payload.memories;
        },
      )
      .addCase(loadMemories.rejected, (state, action) => {
        state.memoryLoading = false;
        state.memoryError = action.payload as string;
      })
      .addCase(getMemoriesLimit.pending, (state) => {
        state.memoryLoading = true;
        state.memoryError = null;
      })
      .addCase(
        getMemoriesLimit.fulfilled,
        (state, action: PayloadAction<LimitResponse<Memory, "memories">>) => {
          state.memoryLoading = false;
          state.memories = action.payload.memories;
          state.memoriesTotal = action.payload.meta.total;
          state.memoriesTotalPages = action.payload.meta.total_pages;
          state.memoriesLimit = action.payload.meta.limit;
        },
      )
      .addCase(getMemoriesLimit.rejected, (state, action) => {
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
      .addCase(
        createMemory.fulfilled,
        (state, action: PayloadAction<Memory>) => {
          state.memoryLoading = false;
          state.memories.push(action.payload);
        },
      )
      .addCase(createMemory.rejected, (state, action) => {
        state.memoryLoading = false;
        state.memoryError = action.payload as string;
      })
      .addCase(
        removeMemory.fulfilled,
        (state, action: PayloadAction<number>) => {
          state.memories = state.memories.filter(
            (memory) => memory.id !== action.payload,
          );
        },
      );
  },
});

export const { clearError } = memorySlice.actions;
export default memorySlice.reducer;
