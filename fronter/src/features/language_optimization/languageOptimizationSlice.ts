import { createSlice, createAsyncThunk, PayloadAction } from '@reduxjs/toolkit';
import { fetchLanguageOptimizationsService, addLanguageOptimizationService, updateLanguageOptimizationService, deleteLanguageOptimizationService } from '../../services/languageOptimization';
import { LanguageOptimization, AddLanguageOptimization } from '../../model/languageOptimization';

interface languageOptimizationState {
  languageOptimization: LanguageOptimization[];
  languageOptimizationLoading: boolean;
  languageOptimizationError: string | null;
}

const initialState: languageOptimizationState = {
  languageOptimization: [],
  languageOptimizationLoading: false,
  languageOptimizationError: null,
}

export const loadLanguageOptimization = createAsyncThunk(
  'languageOptimization/loadLanguageOptimization',
  async (_, { rejectWithValue }) => {
    try {
      const response = await fetchLanguageOptimizationsService();
      if (response.error) {
        return rejectWithValue(response.error);
      }
      return response.languageOptimization || response;
    } catch (error: any) {
      return rejectWithValue(error.message);
    }
  }
)

export const createLanguageOptimization = createAsyncThunk(
  'languageOptimization/createLanguageOptimization',
  async (languageOptimizationData: AddLanguageOptimization, { rejectWithValue }) => {
    try {
      const response = await addLanguageOptimizationService(languageOptimizationData);
      if (response.error) {
        return rejectWithValue(response.error);
      }
      return response;
    } catch (error: any) {
      return rejectWithValue(error.message);
    }
  }
)

export const updateLanguageOptimization = createAsyncThunk(
  'languageOptimization/updateLanguageOptimization',
  async (languageOptimizationData: LanguageOptimization, { rejectWithValue }) => {
    try {
      const response = await updateLanguageOptimizationService(languageOptimizationData);
      if (response.error) {
        return rejectWithValue(response.error);
      }
      return response;
    } catch (error: any) {
      return rejectWithValue(error.message);
    }
  }
)

export const removeLanguageOptimization = createAsyncThunk(
  'languageOptimization/removeLanguageOptimization',
  async (id: string, { rejectWithValue }) => {
    try {
      const response = await deleteLanguageOptimizationService(String(id));
      if (response.error) {
        return rejectWithValue(response.error);
      }
      return { id };
    } catch (error: any) {
      return rejectWithValue(error.message);
    }
  }
)

const languageOptimizationSlice = createSlice({
  name: 'processOptimization',
  initialState,
  reducers: {},
  extraReducers: (builder) => {
    builder
      .addCase(loadLanguageOptimization.pending, (state) => {
        state.languageOptimizationLoading = true;
        state.languageOptimizationError = null;
      })
      .addCase(loadLanguageOptimization.fulfilled, (state, action: PayloadAction<LanguageOptimization[]>) => {
        state.languageOptimizationLoading = false;
        state.languageOptimization = action.payload;
      })
      .addCase(loadLanguageOptimization.rejected, (state, action) => {
        state.languageOptimizationLoading = false;
        state.languageOptimizationError = action.payload as string;
      })
      .addCase(createLanguageOptimization.fulfilled, (state, action: PayloadAction<LanguageOptimization>) => {
        state.languageOptimization.push(action.payload);
      })
      .addCase(updateLanguageOptimization.fulfilled, (state, action: PayloadAction<LanguageOptimization>) => {
        const idx = state.languageOptimization.findIndex(a => a.id === action.payload.id);
        if (idx !== -1) state.languageOptimization[idx] = action.payload;
      })
      .addCase(removeLanguageOptimization.fulfilled, (state, action: PayloadAction<{ id: string }>) => {
        state.languageOptimization = state.languageOptimization.filter(a => a.id !== action.payload.id);
      });
  },
});

export default languageOptimizationSlice.reducer;
