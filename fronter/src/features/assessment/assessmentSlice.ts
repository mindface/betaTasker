import { createSlice, createAsyncThunk, PayloadAction } from '@reduxjs/toolkit';
import { fetchAssessmentsService, getAssessmentsForTaskUserService,  addAssessmentService, updateAssessmentService, deleteAssessmentService } from '../../services/assessmentApi';
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
    const response = await fetchAssessmentsService();
    if ('error' in response) {
      return rejectWithValue(response.error);
    }
    return response.value;
  }
)

export const createAssessment = createAsyncThunk(
  'assessment/createAssessment',
  async (assessmentData: AddAssessment, { rejectWithValue }) => {
    const response = await addAssessmentService(assessmentData);
    if ("error" in response) {
      return rejectWithValue({
        message: response.error.message,
        name: response.error.name,
      });
    }
    return response;
  }
)

export const getAssessmentsForTaskUser = createAsyncThunk(
  'assessment/getAssessmentsForTaskUser',
  async (payload: { userId: number,taskId: number }, { rejectWithValue }) => {
    const response = await getAssessmentsForTaskUserService(payload.userId, payload.taskId);
    if ("error" in response) {
      return rejectWithValue({
        message: response.error.message,
        name: response.error.name,
      });
    }
    return response;
  }
)

export const updateAssessment = createAsyncThunk(
  'assessment/updateAssessment',
  async (assessmentData: Assessment, { rejectWithValue }) => {
    const response = await updateAssessmentService(assessmentData);
    if ("error" in response) {
      return rejectWithValue({
        message: response.error.message,
        name: response.error.name,
      });
    }
    return response;
  }
)

export const removeAssessment = createAsyncThunk(
  'assessment/removeAssessment',
  async (id: number, { rejectWithValue }) => {
    const response = await deleteAssessmentService(String(id));
    if ("error" in response) {
      return rejectWithValue({
        message: response.error.message,
        name: response.error.name,
      });
    }
    return { id };
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
