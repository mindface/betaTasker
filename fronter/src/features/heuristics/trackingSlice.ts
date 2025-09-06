import { createSlice, createAsyncThunk, PayloadAction } from '@reduxjs/toolkit';
import { fetchTrackingService, getTrackingForUserService, addTrackingService, updateTrackingService, deleteTrackingService } from '../../services/heuristicsApi';
import { AddHeuristicsTracking, HeuristicsTracking } from '../../model/heuristics';

interface TrackingState {
  tracking: HeuristicsTracking[];
  trackingLoading: boolean;
  trackingError: string | null;
}

const initialState: TrackingState = {
  tracking: [],
  trackingLoading: false,
  trackingError: null,
}

export const loadTracking = createAsyncThunk(
  'tracking/loadTracking',
  async (_, { rejectWithValue }) => {
    try {
      const response = await fetchTrackingService();
      if (response.error) {
        return rejectWithValue(response.error);
      }
      return response.tracking || response;
    } catch (error: any) {
      return rejectWithValue(error.message);
    }
  }
)

export const createTracking = createAsyncThunk(
  'tracking/createTracking',
  async (trackingData: AddHeuristicsTracking, { rejectWithValue }) => {
    try {
      const response = await addTrackingService(trackingData);
      if (response.error) {
        return rejectWithValue(response.error);
      }
      return response;
    } catch (error: any) {
      return rejectWithValue(error.message);
    }
  }
)

export const getTrackingForUser = createAsyncThunk(
  'tracking/getTrackingForUser',
  async (userId: number, { rejectWithValue }) => {
    try {
      const response = await getTrackingForUserService(userId);
      if (response.error) {
        return rejectWithValue(response.error);
      }
      console.log("payload", response)

      return response.tracking || response;
    } catch (error: any) {
      return rejectWithValue(error.message);
    }
  }
)

export const updateTracking = createAsyncThunk(
  'tracking/updateTracking',
  async (trackingData: HeuristicsTracking, { rejectWithValue }) => {
    try {
      const response = await updateTrackingService(trackingData);
      if (response.error) {
        return rejectWithValue(response.error);
      }
      return response;
    } catch (error: any) {
      return rejectWithValue(error.message);
    }
  }
)

export const removeTracking = createAsyncThunk(
  'tracking/removeTracking',
  async (id: number, { rejectWithValue }) => {
    try {
      const response = await deleteTrackingService(String(id));
      if (response.error) {
        return rejectWithValue(response.error);
      }
      return { id };
    } catch (error: any) {
      return rejectWithValue(error.message);
    }
  }
)

const trackingSlice = createSlice({
  name: 'tracking',
  initialState,
  reducers: {},
  extraReducers: (builder) => {
    builder
      .addCase(loadTracking.pending, (state) => {
        state.trackingLoading = true;
        state.trackingError = null;
      })
      .addCase(loadTracking.fulfilled, (state, action: PayloadAction<HeuristicsTracking[]>) => {
        state.trackingLoading = false;
        state.tracking = action.payload;
      })
      .addCase(loadTracking.rejected, (state, action) => {
        state.trackingLoading = false;
        state.trackingError = action.payload as string;
      })
      .addCase(createTracking.fulfilled, (state, action: PayloadAction<HeuristicsTracking>) => {
        state.tracking.push(action.payload);
      })
      .addCase(updateTracking.fulfilled, (state, action: PayloadAction<HeuristicsTracking>) => {
        const idx = state.tracking.findIndex(t => t.id === action.payload.id);
        if (idx !== -1) state.tracking[idx] = action.payload;
      })
      .addCase(removeTracking.fulfilled, (state, action: PayloadAction<{ id: number }>) => {
        state.tracking = state.tracking.filter(t => t.id !== action.payload.id);
      });
  },
});

export default trackingSlice.reducer;