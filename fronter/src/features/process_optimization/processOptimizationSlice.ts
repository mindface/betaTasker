import { createSlice, createAsyncThunk, PayloadAction } from '@reduxjs/toolkit';
import { fetchProcessOptimizationsService, addProcessOptimizationService, updateProcessOptimizationService, deleteProcessOptimizationService } from '../../services/processOptimizationApi';
import { ProcessOptimization, AddProcessOptimization } from '../../model/processOptimization';

interface processOptimizationState {
  processOptimization: ProcessOptimization[];
  processOptimizationLoading: boolean;
  processOptimizationError: string | null;
}

const initialState: processOptimizationState = {
  processOptimization: [],
  processOptimizationLoading: false,
  processOptimizationError: null,
}

export const loadProcessOptimization = createAsyncThunk(
  'processOptimization/loadProcessOptimization',
  async (_, { rejectWithValue }) => {
    try {
      const response = await fetchProcessOptimizationsService();
      if (response.error) {
        return rejectWithValue(response.error);
      }
      return response.processOptimization || response;
    } catch (error: any) {
      return rejectWithValue(error.message);
    }
  }
)

export const createProcessOptimization = createAsyncThunk(
  'processOptimization/createProcessOptimization',
  async (processOptimizationData: AddProcessOptimization, { rejectWithValue }) => {
    try {
      const response = await addProcessOptimizationService(processOptimizationData);
      if (response.error) {
        return rejectWithValue(response.error);
      }
      return response;
    } catch (error: any) {
      return rejectWithValue(error.message);
    }
  }
)

export const updateProcessOptimization = createAsyncThunk(
  'processOptimization/updateProcessOptimization',
  async (processOptimizationData: ProcessOptimization, { rejectWithValue }) => {
    try {
      const response = await updateProcessOptimizationService(processOptimizationData);
      if (response.error) {
        return rejectWithValue(response.error);
      }
      return response;
    } catch (error: any) {
      return rejectWithValue(error.message);
    }
  }
)

export const removeProcessOptimization = createAsyncThunk(
  'processOptimization/removeProcessOptimization',
  async (id: string, { rejectWithValue }) => {
    try {
      const response = await deleteProcessOptimizationService(String(id));
      if (response.error) {
        return rejectWithValue(response.error);
      }
      return { id };
    } catch (error: any) {
      return rejectWithValue(error.message);
    }
  }
)

const processOptimizationSlice = createSlice({
  name: 'processOptimization',
  initialState,
  reducers: {},
  extraReducers: (builder) => {
    builder
      .addCase(loadProcessOptimization.pending, (state) => {
        state.processOptimizationLoading = true;
        state.processOptimizationError = null;
      })
      .addCase(loadProcessOptimization.fulfilled, (state, action: PayloadAction<ProcessOptimization[]>) => {
        state.processOptimizationLoading = false;
        state.processOptimization = action.payload;
      })
      .addCase(loadProcessOptimization.rejected, (state, action) => {
        state.processOptimizationLoading = false;
        state.processOptimizationError = action.payload as string;
      })
      .addCase(createProcessOptimization.fulfilled, (state, action: PayloadAction<ProcessOptimization>) => {
        state.processOptimization.push(action.payload);
      })
      .addCase(updateProcessOptimization.fulfilled, (state, action: PayloadAction<ProcessOptimization>) => {
        const idx = state.processOptimization.findIndex(a => a.id === action.payload.id);
        if (idx !== -1) state.processOptimization[idx] = action.payload;
      })
      .addCase(removeProcessOptimization.fulfilled, (state, action: PayloadAction<{ id: string }>) => {
        state.processOptimization = state.processOptimization.filter(a => a.id !== action.payload.id);
      });
  },
});

export default processOptimizationSlice.reducer;
