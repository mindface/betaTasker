import { createSlice, createAsyncThunk, PayloadAction } from '@reduxjs/toolkit';
import { fetchInsightsService, getInsightsForUserService, addInsightService, updateInsightService, deleteInsightService } from '../../services/heuristicsApi';
import { AddHeuristicsInsight, HeuristicsInsight } from '../../model/heuristics';

interface InsightsState {
  insights: HeuristicsInsight[];
  insightsLoading: boolean;
  insightsError: string | null;
}

const initialState: InsightsState = {
  insights: [],
  insightsLoading: false,
  insightsError: null,
}

export const loadInsights = createAsyncThunk(
  'insights/loadInsights',
  async (_, { rejectWithValue }) => {
    try {
      const response = await fetchInsightsService();
      if (response.error) {
        return rejectWithValue(response.error);
      }
      return response.insights || response;
    } catch (error: any) {
      return rejectWithValue(error.message);
    }
  }
)

export const createInsight = createAsyncThunk(
  'insights/createInsight',
  async (insightData: AddHeuristicsInsight, { rejectWithValue }) => {
    try {
      const response = await addInsightService(insightData);
      if (response.error) {
        return rejectWithValue(response.error);
      }
      return response;
    } catch (error: any) {
      return rejectWithValue(error.message);
    }
  }
)

export const getInsightsForUser = createAsyncThunk(
  'insights/getInsightsForUser',
  async (userId: number, { rejectWithValue }) => {
    try {
      const response = await getInsightsForUserService(userId);
      if (response.error) {
        return rejectWithValue(response.error);
      }
      console.log("payload", response)

      return response.insights || response;
    } catch (error: any) {
      return rejectWithValue(error.message);
    }
  }
)

export const updateInsight = createAsyncThunk(
  'insights/updateInsight',
  async (insightData: HeuristicsInsight, { rejectWithValue }) => {
    try {
      const response = await updateInsightService(insightData);
      if (response.error) {
        return rejectWithValue(response.error);
      }
      return response;
    } catch (error: any) {
      return rejectWithValue(error.message);
    }
  }
)

export const removeInsight = createAsyncThunk(
  'insights/removeInsight',
  async (id: number, { rejectWithValue }) => {
    try {
      const response = await deleteInsightService(String(id));
      if (response.error) {
        return rejectWithValue(response.error);
      }
      return { id };
    } catch (error: any) {
      return rejectWithValue(error.message);
    }
  }
)

const insightsSlice = createSlice({
  name: 'insights',
  initialState,
  reducers: {},
  extraReducers: (builder) => {
    builder
      .addCase(loadInsights.pending, (state) => {
        state.insightsLoading = true;
        state.insightsError = null;
      })
      .addCase(loadInsights.fulfilled, (state, action: PayloadAction<HeuristicsInsight[]>) => {
        state.insightsLoading = false;
        state.insights = action.payload;
      })
      .addCase(loadInsights.rejected, (state, action) => {
        state.insightsLoading = false;
        state.insightsError = action.payload as string;
      })
      .addCase(createInsight.fulfilled, (state, action: PayloadAction<HeuristicsInsight>) => {
        state.insights.push(action.payload);
      })
      .addCase(updateInsight.fulfilled, (state, action: PayloadAction<HeuristicsInsight>) => {
        const idx = state.insights.findIndex(i => i.id === action.payload.id);
        if (idx !== -1) state.insights[idx] = action.payload;
      })
      .addCase(removeInsight.fulfilled, (state, action: PayloadAction<{ id: number }>) => {
        state.insights = state.insights.filter(i => i.id !== action.payload.id);
      });
  },
});

export default insightsSlice.reducer;