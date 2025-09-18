import { createSlice, createAsyncThunk, PayloadAction } from '@reduxjs/toolkit';
import { fetchQualitativeLabelsService, addQualitativeLabelService, updateQualitativeLabelService, deleteQualitativeLabelService } from '../../services/qualitativeLabel';
import { QualitativeLabel, AddQualitativeLabel } from '../../model/qualitativeLabel';

interface qualitativeLabelState {
  qualitativeLabels: QualitativeLabel[];
  qualitativeLabelLoading: boolean;
  qualitativeLabelError: string | null;
}

const initialState: qualitativeLabelState = {
  qualitativeLabels: [],
  qualitativeLabelLoading: false,
  qualitativeLabelError: null,
}

export const loadQualitativeLabels = createAsyncThunk(
  'qualitativeLabel/loadQualitativeLabels',
  async (_, { rejectWithValue }) => {
    try {
      const response = await fetchQualitativeLabelsService();
      if (response.error) {
        return rejectWithValue(response.error);
      }
      return response.languageOptimization || response;
    } catch (error: any) {
      return rejectWithValue(error.message);
    }
  }
)

export const createQualitativeLabel = createAsyncThunk(
  'qualitativeLabel/createQualitativeLabel',
  async (qualitativeLabelData: AddQualitativeLabel, { rejectWithValue }) => {
    try {
      const response = await addQualitativeLabelService(qualitativeLabelData);
      if (response.error) {
        return rejectWithValue(response.error);
      }
      return response;
    } catch (error: any) {
      return rejectWithValue(error.message);
    }
  }
)

export const updateQualitativeLabel = createAsyncThunk(
  'qualitativeLabel/updateQualitativeLabel',
  async (qualitativeLabelData: QualitativeLabel, { rejectWithValue }) => {
    try {
      const response = await updateQualitativeLabelService(qualitativeLabelData);
      if (response.error) {
        return rejectWithValue(response.error);
      }
      return response;
    } catch (error: any) {
      return rejectWithValue(error.message);
    }
  }
)

export const removeQualitativeLabel = createAsyncThunk(
  'qualitativeLabel/removeQualitativeLabel',
  async (id: string, { rejectWithValue }) => {
    try {
      const response = await deleteQualitativeLabelService(String(id));
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
  name: 'qualitativeLabel',
  initialState,
  reducers: {},
  extraReducers: (builder) => {
    builder
      .addCase(loadQualitativeLabels.pending, (state) => {
        state.qualitativeLabelLoading = true;
        state.qualitativeLabelError = null;
      })
      .addCase(loadQualitativeLabels.fulfilled, (state, action: PayloadAction<QualitativeLabel[]>) => {
        state.qualitativeLabelLoading = false;
        state.qualitativeLabels = action.payload;
      })
      .addCase(loadQualitativeLabels.rejected, (state, action) => {
        state.qualitativeLabelLoading = false;
        state.qualitativeLabelError = action.payload as string;
      })
      .addCase(createQualitativeLabel.fulfilled, (state, action: PayloadAction<QualitativeLabel>) => {
        state.qualitativeLabels.push(action.payload);
      })
      .addCase(updateQualitativeLabel.fulfilled, (state, action: PayloadAction<QualitativeLabel>) => {
        const idx = state.qualitativeLabels.findIndex(a => a.id === action.payload.id);
        if (idx !== -1) state.qualitativeLabels[idx] = action.payload;
      })
      .addCase(removeQualitativeLabel.fulfilled, (state, action: PayloadAction<{ id: string }>) => {
        state.qualitativeLabels = state.qualitativeLabels.filter(a => a.id !== action.payload.id);
      });
  },
});

export default languageOptimizationSlice.reducer;
