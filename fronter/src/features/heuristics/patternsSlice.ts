import { createSlice, createAsyncThunk, PayloadAction } from '@reduxjs/toolkit';
import { fetchPatternsService, addPatternService, updatePatternService, deletePatternService } from '../../services/heuristicsApi';
import { AddHeuristicsPattern, HeuristicsPattern } from '../../model/heuristics';

interface PatternsState {
  patterns: HeuristicsPattern[];
  patternsLoading: boolean;
  patternsError: string | null;
}

const initialState: PatternsState = {
  patterns: [],
  patternsLoading: false,
  patternsError: null,
}

export const loadPatterns = createAsyncThunk(
  'patterns/loadPatterns',
  async (_, { rejectWithValue }) => {
    try {
      const response = await fetchPatternsService();
      if (response.error) {
        return rejectWithValue(response.error);
      }
      return response.patterns || response;
    } catch (error: any) {
      return rejectWithValue(error.message);
    }
  }
)

export const createPattern = createAsyncThunk(
  'patterns/createPattern',
  async (patternData: AddHeuristicsPattern, { rejectWithValue }) => {
    try {
      const response = await addPatternService(patternData);
      if (response.error) {
        return rejectWithValue(response.error);
      }
      return response;
    } catch (error: any) {
      return rejectWithValue(error.message);
    }
  }
)

export const updatePattern = createAsyncThunk(
  'patterns/updatePattern',
  async (patternData: HeuristicsPattern, { rejectWithValue }) => {
    try {
      const response = await updatePatternService(patternData);
      if (response.error) {
        return rejectWithValue(response.error);
      }
      return response;
    } catch (error: any) {
      return rejectWithValue(error.message);
    }
  }
)

export const removePattern = createAsyncThunk(
  'patterns/removePattern',
  async (id: number, { rejectWithValue }) => {
    try {
      const response = await deletePatternService(String(id));
      if (response.error) {
        return rejectWithValue(response.error);
      }
      return { id };
    } catch (error: any) {
      return rejectWithValue(error.message);
    }
  }
)

const patternsSlice = createSlice({
  name: 'patterns',
  initialState,
  reducers: {},
  extraReducers: (builder) => {
    builder
      .addCase(loadPatterns.pending, (state) => {
        state.patternsLoading = true;
        state.patternsError = null;
      })
      .addCase(loadPatterns.fulfilled, (state, action: PayloadAction<HeuristicsPattern[]>) => {
        state.patternsLoading = false;
        state.patterns = action.payload;
      })
      .addCase(loadPatterns.rejected, (state, action) => {
        state.patternsLoading = false;
        state.patternsError = action.payload as string;
      })
      .addCase(createPattern.fulfilled, (state, action: PayloadAction<HeuristicsPattern>) => {
        state.patterns.push(action.payload);
      })
      .addCase(updatePattern.fulfilled, (state, action: PayloadAction<HeuristicsPattern>) => {
        const idx = state.patterns.findIndex(p => p.id === action.payload.id);
        if (idx !== -1) state.patterns[idx] = action.payload;
      })
      .addCase(removePattern.fulfilled, (state, action: PayloadAction<{ id: number }>) => {
        state.patterns = state.patterns.filter(p => p.id !== action.payload.id);
      });
  },
});

export default patternsSlice.reducer;