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
    const response = await fetchQualitativeLabelsService();
    if ('error' in response) {
      return rejectWithValue(response.error);
    }
    return response.value;
  }
)

export const createQualitativeLabel = createAsyncThunk(
  'qualitativeLabel/createQualitativeLabel',
  async (qualitativeLabelData: AddQualitativeLabel, { rejectWithValue }) => {
    const response = await addQualitativeLabelService(qualitativeLabelData);
    if ('error' in response) {
      return rejectWithValue(response.error);
    }
    return response.value;
  }
)

export const updateQualitativeLabel = createAsyncThunk(
  'qualitativeLabel/updateQualitativeLabel',
  async (qualitativeLabelData: QualitativeLabel, { rejectWithValue }) => {
    const response = await updateQualitativeLabelService(qualitativeLabelData);
    if ('error' in response) {
      return rejectWithValue(response.error);
    }
    return response.value;
  }
)

export const removeQualitativeLabel = createAsyncThunk(
  'qualitativeLabel/removeQualitativeLabel',
  async (id: string, { rejectWithValue }) => {
    const response = await deleteQualitativeLabelService(String(id));
    if ('error' in response) {
      return rejectWithValue(response.error);
    }
    return { id };
  }
)

const qualitativeLabelSlice = createSlice({
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

export default qualitativeLabelSlice.reducer;
