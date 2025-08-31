import { createSlice, createAsyncThunk } from '@reduxjs/toolkit';
import { getHeuristicsInsights } from '../../services/heuristicsApi';
import { InsightsResponse } from '../../app/api/heuristics/insights/route';

export const fetchInsights = createAsyncThunk(
  'insights/fetchInsights',
  async () => await getHeuristicsInsights()
);

const insightsSlice = createSlice({
  name: 'insights',
  initialState: {
    loading: false,
    data: null as InsightsResponse['data'] | null,
    error: null as string | null
  },
  reducers: {},
  extraReducers: builder => {
    builder
      .addCase(fetchInsights.pending, state => {
        state.loading = true;
        state.error = null;
      })
      .addCase(fetchInsights.fulfilled, (state, action) => {
        state.loading = false;
        state.data = action.payload.data;
        state.error = action.payload.error ?? null;
      })
      .addCase(fetchInsights.rejected, (state, action) => {
        state.loading = false;
        state.error = action.error.message ?? 'Unknown error';
      });
  }
});

export default insightsSlice.reducer;