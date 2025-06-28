import { createSlice, createAsyncThunk, PayloadAction } from '@reduxjs/toolkit';
import { fetchAssessmentsService, addAssessmentService, updateAssessmentService, deleteAssessmentService } from '../../services/assessmentApi';
import { AddAssessment, Assessment } from '../../model/assessment';

interface AssessmentState {
  assessments: Assessment[];
  assessmentLoading: boolean;
  assessmentError: string | null;
}

const initialState: AssessmentState = {
  assessments: [],
  assessmentLoading: false,
  assessmentError: null,
}

export const loadAssessments = createAsyncThunk(
  'assessment/loadAssessments',
  async (_, { rejectWithValue }) => {
    try {
      const response = await fetchAssessmentsService();
      if (response.error) {
        return rejectWithValue(response.error);
      }
      return response.assessments || response;
    } catch (error: any) {
      return rejectWithValue(error.message);
    }
  }
)

export const createAssessment = createAsyncThunk(
  'assessment/createAssessment',
  async (assessmentData: AddAssessment, { rejectWithValue }) => {
    try {
      const response = await addAssessmentService(assessmentData);
      if (response.error) {
        return rejectWithValue(response.error);
      }
      return response;
    } catch (error: any) {
      return rejectWithValue(error.message);
    }
  }
)

export const updateAssessment = createAsyncThunk(
  'assessment/updateAssessment',
  async (assessmentData: Assessment, { rejectWithValue }) => {
    try {
      const response = await updateAssessmentService(assessmentData);
      if (response.error) {
        return rejectWithValue(response.error);
      }
      return response;
    } catch (error: any) {
      return rejectWithValue(error.message);
    }
  }
)

export const removeAssessment = createAsyncThunk(
  'assessment/removeAssessment',
  async (id: number, { rejectWithValue }) => {
    try {
      const response = await deleteAssessmentService(String(id));
      if (response.error) {
        return rejectWithValue(response.error);
      }
      return { id };
    } catch (error: any) {
      return rejectWithValue(error.message);
    }
  }
)

const assessmentSlice = createSlice({
  name: 'assessment',
  initialState,
  reducers: {},
  extraReducers: (builder) => {
    builder
      .addCase(loadAssessments.pending, (state) => {
        state.assessmentLoading = true;
        state.assessmentError = null;
      })
      .addCase(loadAssessments.fulfilled, (state, action: PayloadAction<Assessment[]>) => {
        state.assessmentLoading = false;
        state.assessments = action.payload;
      })
      .addCase(loadAssessments.rejected, (state, action) => {
        state.assessmentLoading = false;
        state.assessmentError = action.payload as string;
      })
      .addCase(createAssessment.fulfilled, (state, action: PayloadAction<Assessment>) => {
        state.assessments.push(action.payload);
      })
      .addCase(updateAssessment.fulfilled, (state, action: PayloadAction<Assessment>) => {
        const idx = state.assessments.findIndex(a => a.id === action.payload.id);
        if (idx !== -1) state.assessments[idx] = action.payload;
      })
      .addCase(removeAssessment.fulfilled, (state, action: PayloadAction<{ id: number }>) => {
        state.assessments = state.assessments.filter(a => a.id !== action.payload.id);
      });
  },
});

export default assessmentSlice.reducer;
