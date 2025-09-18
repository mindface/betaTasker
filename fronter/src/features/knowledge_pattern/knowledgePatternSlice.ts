import { createSlice, createAsyncThunk, PayloadAction } from '@reduxjs/toolkit';
import { fetchKnowledgePatternsService, addKnowledgePatternService, updateKnowledgePatternService, deleteKnowledgePatternService } from '../../services/knowledgePattern';
import { KnowledgePattern, AddKnowledgePattern } from '../../model/knowledgePattern';

interface knowledgePatternState {
  knowledgePatterns: KnowledgePattern[];
  knowledgePatternsLoading: boolean;
  knowledgePatternsError: string | null;
}

const initialState: knowledgePatternState = {
  knowledgePatterns: [],
  knowledgePatternsLoading: false,
  knowledgePatternsError: null,
}

export const loadKnowledgePatterns = createAsyncThunk(
  'knowledgePatterns/loadKnowledgePatterns',
  async (_, { rejectWithValue }) => {
    try {
      const response = await fetchKnowledgePatternsService();
      if (response.error) {
        return rejectWithValue(response.error);
      }
      return response.languageOptimization || response;
    } catch (error: any) {
      return rejectWithValue(error.message);
    }
  }
)

export const createKnowledgePattern = createAsyncThunk(
  'knowledgePatterns/createKnowledgePattern',
  async (knowledgePatternData: AddKnowledgePattern, { rejectWithValue }) => {
    try {
      const response = await addKnowledgePatternService(knowledgePatternData);
      if (response.error) {
        return rejectWithValue(response.error);
      }
      return response;
    } catch (error: any) {
      return rejectWithValue(error.message);
    }
  }
)

export const updateKnowledgePattern = createAsyncThunk(
  'knowledgePatterns/updateKnowledgePattern',
  async (knowledgePatternData: KnowledgePattern, { rejectWithValue }) => {
    try {
      const response = await updateKnowledgePatternService(knowledgePatternData);
      if (response.error) {
        return rejectWithValue(response.error);
      }
      return response;
    } catch (error: any) {
      return rejectWithValue(error.message);
    }
  }
)

export const removeKnowledgePattern = createAsyncThunk(
  'knowledgePatterns/removeKnowledgePattern',
  async (id: string, { rejectWithValue }) => {
    try {
      const response = await deleteKnowledgePatternService(String(id));
      if (response.error) {
        return rejectWithValue(response.error);
      }
      return { id };
    } catch (error: any) {
      return rejectWithValue(error.message);
    }
  }
)

const knowledgePatternSlice = createSlice({
  name: 'knowledgePatterns',
  initialState,
  reducers: {},
  extraReducers: (builder) => {
    builder
      .addCase(loadKnowledgePatterns.pending, (state) => {
        state.knowledgePatternsLoading = true;
        state.knowledgePatternsError = null;
      })
      .addCase(loadKnowledgePatterns.fulfilled, (state, action: PayloadAction<KnowledgePattern[]>) => {
        state.knowledgePatternsLoading = false;
        state.knowledgePatterns = action.payload;
      })
      .addCase(loadKnowledgePatterns.rejected, (state, action) => {
        state.knowledgePatternsLoading = false;
        state.knowledgePatternsError = action.payload as string;
      })
      .addCase(createKnowledgePattern.fulfilled, (state, action: PayloadAction<KnowledgePattern>) => {
        state.knowledgePatterns.push(action.payload);
      })
      .addCase(updateKnowledgePattern.fulfilled, (state, action: PayloadAction<KnowledgePattern>) => {
        const idx = state.knowledgePatterns.findIndex(a => a.id === action.payload.id);
        if (idx !== -1) state.knowledgePatterns[idx] = action.payload;
      })
      .addCase(removeKnowledgePattern.fulfilled, (state, action: PayloadAction<{ id: string }>) => {
        state.knowledgePatterns = state.knowledgePatterns.filter(a => a.id !== action.payload.id);
      });
  },
});

export default knowledgePatternSlice.reducer;
