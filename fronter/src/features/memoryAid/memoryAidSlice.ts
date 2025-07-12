import { createSlice, createAsyncThunk, PayloadAction } from '@reduxjs/toolkit';
import { fetchMemoryAidsByCode } from '../../services/memoryAidApi';
import { MemoryContext } from '../../model/memoryAid';

interface MemoryAidState {
  contexts: MemoryContext[];
  loading: boolean;
  error: string | null;
}

const initialState: MemoryAidState = {
  contexts: [],
  loading: false,
  error: null,
};

export const loadMemoryAidsByCode = createAsyncThunk(
  'memoryAid/loadMemoryAidsByCode',
  async (code: string, { rejectWithValue }) => {
    try {
      const response = await fetchMemoryAidsByCode(code);
      if (response.error) {
        return rejectWithValue(response.error);
      }
      // APIレスポンスが {contexts: MemoryContext[]} 形式の場合
      return response.contexts;
    } catch (error: any) {
      return rejectWithValue(error.message);
    }
  }
);

const memoryAidSlice = createSlice({
  name: 'memoryAid',
  initialState,
  reducers: {},
  extraReducers: (builder) => {
    builder
      .addCase(loadMemoryAidsByCode.pending, (state) => {
        state.loading = true;
        state.error = null;
      })
      .addCase(loadMemoryAidsByCode.fulfilled, (state, action: PayloadAction<MemoryContext[]>) => {
        state.loading = false;
        state.contexts = action.payload;
      })
      .addCase(loadMemoryAidsByCode.rejected, (state, action) => {
        state.loading = false;
        state.error = action.payload as string;
      });
  },
});

export default memoryAidSlice.reducer;
