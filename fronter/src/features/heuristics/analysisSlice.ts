import { createSlice, createAsyncThunk, PayloadAction } from '@reduxjs/toolkit';
import { fetchAnalysesService, getAnalysisForTaskUserService, addAnalysisService, updateAnalysisService, deleteAnalysisService } from '../../services/heuristicsApi';
import { AddHeuristicsAnalysis, HeuristicsAnalysis } from '../../model/heuristics';

interface AnalysisState {
  analyses: HeuristicsAnalysis[];
  analysisLoading: boolean;
  analysisError: string | null;
}

const initialState: AnalysisState = {
  analyses: [],
  analysisLoading: false,
  analysisError: null,
}

export const loadAnalyses = createAsyncThunk(
  'analysis/loadAnalyses',
  async (_, { rejectWithValue }) => {
    try {
      const response = await fetchAnalysesService();
      if (response.error) {
        return rejectWithValue(response.error);
      }
      return response.analyses || response;
    } catch (error: any) {
      return rejectWithValue(error.message);
    }
  }
)

export const createAnalysis = createAsyncThunk(
  'analysis/createAnalysis',
  async (analysisData: AddHeuristicsAnalysis, { rejectWithValue }) => {
    try {
      const response = await addAnalysisService(analysisData);
      if (response.error) {
        return rejectWithValue(response.error);
      }
      return response;
    } catch (error: any) {
      return rejectWithValue(error.message);
    }
  }
)

export const getAnalysisForTaskUser = createAsyncThunk(
  'analysis/getAnalysisForTaskUser',
  async (payload: { userId: number, taskId: number }, { rejectWithValue }) => {
    try {
      const response = await getAnalysisForTaskUserService(payload.userId, payload.taskId);
      if (response.error) {
        return rejectWithValue(response.error);
      }
      console.log("payload", response)

      return response.analyses || response;
    } catch (error: any) {
      return rejectWithValue(error.message);
    }
  }
)

export const updateAnalysis = createAsyncThunk(
  'analysis/updateAnalysis',
  async (analysisData: HeuristicsAnalysis, { rejectWithValue }) => {
    try {
      const response = await updateAnalysisService(analysisData);
      if (response.error) {
        return rejectWithValue(response.error);
      }
      return response;
    } catch (error: any) {
      return rejectWithValue(error.message);
    }
  }
)

export const removeAnalysis = createAsyncThunk(
  'analysis/removeAnalysis',
  async (id: number, { rejectWithValue }) => {
    try {
      const response = await deleteAnalysisService(String(id));
      if (response.error) {
        return rejectWithValue(response.error);
      }
      return { id };
    } catch (error: any) {
      return rejectWithValue(error.message);
    }
  }
)

const analysisSlice = createSlice({
  name: 'analysis',
  initialState,
  reducers: {},
  extraReducers: (builder) => {
    builder
      .addCase(loadAnalyses.pending, (state) => {
        state.analysisLoading = true;
        state.analysisError = null;
      })
      .addCase(loadAnalyses.fulfilled, (state, action: PayloadAction<HeuristicsAnalysis[]>) => {
        state.analysisLoading = false;
        state.analyses = action.payload;
      })
      .addCase(loadAnalyses.rejected, (state, action) => {
        state.analysisLoading = false;
        state.analysisError = action.payload as string;
      })
      .addCase(createAnalysis.fulfilled, (state, action: PayloadAction<HeuristicsAnalysis>) => {
        state.analyses.push(action.payload);
      })
      .addCase(updateAnalysis.fulfilled, (state, action: PayloadAction<HeuristicsAnalysis>) => {
        const idx = state.analyses.findIndex(a => a.id === action.payload.id);
        if (idx !== -1) state.analyses[idx] = action.payload;
      })
      .addCase(removeAnalysis.fulfilled, (state, action: PayloadAction<{ id: number }>) => {
        state.analyses = state.analyses.filter(a => a.id !== action.payload.id);
      });
  },
});

export default analysisSlice.reducer;