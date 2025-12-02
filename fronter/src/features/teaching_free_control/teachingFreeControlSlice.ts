import { createSlice, createAsyncThunk, PayloadAction } from '@reduxjs/toolkit';
import { fetchTeachingFreeControlService, addTeachingFreeControlService, updateTeachingFreeControlService, deleteTeachingFreeControlService } from '../../client/teachingFreeControl';
import { TeachingFreeControl, AddTeachingFreeControl } from '../../model/teachingFreeControl';

interface teachingFreeControlState {
  teachingFreeControl: TeachingFreeControl[];
  teachingFreeControlLoading: boolean;
  teachingFreeControlError: string | null;
}

const initialState: teachingFreeControlState = {
  teachingFreeControl: [],
  teachingFreeControlLoading: false,
  teachingFreeControlError: null,
}

export const loadTeachingFreeControl = createAsyncThunk(
  'teachingFreeControl/loadTeachingFreeControl',
  async (_, { rejectWithValue }) => {
    const response = await fetchTeachingFreeControlService();
    if ('error' in response) {
      return rejectWithValue(response.error);
    }
    return response.value;
  }
)

export const createTeachingFreeControl = createAsyncThunk(
  'teachingFreeControl/createTeachingFreeControl',
  async (teachingFreeControlData: AddTeachingFreeControl, { rejectWithValue }) => {
    const response = await addTeachingFreeControlService(teachingFreeControlData);
    if ('error' in response) {
      return rejectWithValue(response.error);
    }
    return response.value;
  }
)

export const updateTeachingFreeControl = createAsyncThunk(
  'teachingFreeControl/updateTeachingFreeControl',
  async (teachingFreeControlData: TeachingFreeControl, { rejectWithValue }) => {
    const response = await updateTeachingFreeControlService(teachingFreeControlData);
    if ('error' in response) {
      return rejectWithValue(response.error);
    }
    return response.value;
  }
)

export const removeTeachingFreeControl = createAsyncThunk(
  'teachingFreeControl/removeTeachingFreeControl',
  async (id: string, { rejectWithValue }) => {
    const response = await deleteTeachingFreeControlService(String(id));
    if ('error' in response) {
      return rejectWithValue(response.error);
    }
    return { id };
  }
)

const TeachingFreeControlSlice = createSlice({
  name: 'teachingFreeControl',
  initialState,
  reducers: {},
  extraReducers: (builder) => {
    builder
      .addCase(loadTeachingFreeControl.pending, (state) => {
        state.teachingFreeControlLoading = true;
        state.teachingFreeControlError = null;
      })
      .addCase(loadTeachingFreeControl.fulfilled, (state, action: PayloadAction<TeachingFreeControl[]>) => {
        state.teachingFreeControlLoading = false;
        state.teachingFreeControl = action.payload;
      })
      .addCase(loadTeachingFreeControl.rejected, (state, action) => {
        state.teachingFreeControlLoading = false;
        state.teachingFreeControlError = action.payload as string;
      })
      .addCase(createTeachingFreeControl.fulfilled, (state, action: PayloadAction<TeachingFreeControl>) => {
        state.teachingFreeControl.push(action.payload);
      })
      .addCase(updateTeachingFreeControl.fulfilled, (state, action: PayloadAction<TeachingFreeControl>) => {
        const idx = state.teachingFreeControl.findIndex(a => a.id === action.payload.id);
        if (idx !== -1) state.teachingFreeControl[idx] = action.payload;
      })
      .addCase(removeTeachingFreeControl.fulfilled, (state, action: PayloadAction<{ id: string }>) => {
        state.teachingFreeControl = state.teachingFreeControl.filter(a => a.id !== action.payload.id);
      });
  },
});

export default TeachingFreeControlSlice.reducer;
