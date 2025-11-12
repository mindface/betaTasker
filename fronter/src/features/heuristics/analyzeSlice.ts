import { createSlice, createAsyncThunk } from '@reduxjs/toolkit';
import { analyzeHeuristics } from '../../services/heuristicsApi';
import { AnalyzeResponse } from '../../app/api/heuristics/analyze/route';

export const fetchAnalyze = createAsyncThunk(
  'analyze/fetchAnalyze',
  async (params: any) => await analyzeHeuristics(params)
);

const analyzeSlice = createSlice({
  name: 'analyze',
  initialState: {
    loading: false,
    data: null as AnalyzeResponse['data'] | null,
    error: null as string | null
  },
  reducers: {},
  extraReducers: builder => {
    builder
      .addCase(fetchAnalyze.pending, state => {
        state.loading = true;
        state.error = null;
      })
      .addCase(fetchAnalyze.fulfilled, (state, action) => {
        state.loading = false;
        state.data = action.payload.data;
        state.error = action.payload.error ?? null;
      })
      .addCase(fetchAnalyze.rejected, (state, action) => {
        state.loading = false;
        state.error = action.error.message ?? 'Unknown error';
      });
  }
});

export default analyzeSlice.reducer;