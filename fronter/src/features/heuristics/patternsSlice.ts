import { createSlice, createAsyncThunk } from '@reduxjs/toolkit';
import { trackHeuristics } from '../../services/heuristicsApi';
import { TrackResponse } from '../../app/api/heuristics/track/route';

export const fetchTrack = createAsyncThunk(
  'track/fetchTrack',
  async (params: any) => await trackHeuristics(params)
);

const trackSlice = createSlice({
  name: 'track',
  initialState: {
    loading: false,
    data: null as TrackResponse['data'] | null,
    error: null as string | null
  },
  reducers: {},
  extraReducers: builder => {
    builder
      .addCase(fetchTrack.pending, state => {
        state.loading = true;
        state.error = null;
      })
      .addCase(fetchTrack.fulfilled, (state, action) => {
        state.loading = false;
        state.data = action.payload.data;
        state.error = action.payload.error ?? null;
      })
      .addCase(fetchTrack.rejected, (state, action) => {
        state.loading = false;
        state.error = action.error.message ?? 'Unknown error';
      });
  }
});

export default trackSlice.reducer;