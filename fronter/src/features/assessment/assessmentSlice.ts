import { createSlice, createAsyncThunk, PayloadAction } from "@reduxjs/toolkit";
import {
  fetchAssessmentsClient,
  getAssessmentsForTaskUserClient,
  getAssessmentsLimitClient,
  addAssessmentClient,
  updateAssessmentClient,
  deleteAssessmentClient,
} from "../../client/assessmentApi";
import { AddAssessment, Assessment } from "../../model/assessment";
import { LimitResponse } from "@/model/respose";

interface AssessmentState {
  assessments: Assessment[];

  assessmentsPage: number;
  assessmentsLimit: number;
  assessmentsTotal: number;
  assessmentsTotalPages: number;

  assessmentLoading: boolean;
  assessmentError: string | null;
}

const initialState: AssessmentState = {
  assessments: [],

  assessmentsPage: 1,
  assessmentsLimit: 20,
  assessmentsTotal: 0,
  assessmentsTotalPages: 0,

  assessmentLoading: false,
  assessmentError: null,
};

export const loadAssessments = createAsyncThunk(
  "assessment/loadAssessments",
  async (_, { rejectWithValue }) => {
    const response = await fetchAssessmentsClient();
    if ("error" in response) {
      return rejectWithValue(response.error);
    }
    return { assessments: response.assessments, meta: response.meta };
  },
);

export const createAssessment = createAsyncThunk(
  "assessment/createAssessment",
  async (assessmentData: AddAssessment, { rejectWithValue }) => {
    const response = await addAssessmentClient(assessmentData);
    if ("error" in response) {
      return rejectWithValue({
        message: response.error.message,
        name: response.error.name,
      });
    }
    return response;
  },
);

export const getAssessmentsForTaskUser = createAsyncThunk(
  "assessment/getAssessmentsForTaskUser",
  async (payload: { userId: number; taskId: number }, { rejectWithValue }) => {
    const response = await getAssessmentsForTaskUserClient(
      payload.userId,
      payload.taskId
    );
    if ("error" in response) {
      return rejectWithValue({
        message: response.error.message,
        name: response.error.name,
      });
    }
    return response;
  },
);

export const getAssessmentsLimit = createAsyncThunk(
  "assessment/getAssessmentsLimit",
  async (payload: { page: number; limit: number }, { rejectWithValue }) => {
    const response = await getAssessmentsLimitClient(
      payload.page,
      payload.limit,
    );
    if ("error" in response) {
      return rejectWithValue({
        message: response.error.message,
        name: response.error.name,
      });
    }
    return { assessments: response.assessments, meta: response.meta };
  },
);

export const updateAssessment = createAsyncThunk(
  "assessment/updateAssessment",
  async (assessmentData: Assessment, { rejectWithValue }) => {
    const response = await updateAssessmentClient(assessmentData);
    if ("error" in response) {
      return rejectWithValue({
        message: response.error.message,
        name: response.error.name,
      });
    }
    return response;
  },
);

export const removeAssessment = createAsyncThunk(
  "assessment/removeAssessment",
  async (id: number, { rejectWithValue }) => {
    const response = await deleteAssessmentClient(String(id));
    if ("error" in response) {
      return rejectWithValue({
        message: response.error.message,
        name: response.error.name,
      });
    }
    return { id };
  },
);

const assessmentSlice = createSlice({
  name: "assessment",
  initialState,
  reducers: {},
  extraReducers: (builder) => {
    builder
      .addCase(loadAssessments.pending, (state) => {
        state.assessmentLoading = true;
        state.assessmentError = null;
      })
      .addCase(
        loadAssessments.fulfilled,
        (state, action: PayloadAction<LimitResponse<Assessment,"assessments">>) => {
          state.assessmentLoading = false;
          state.assessments = action.payload.assessments;
        },
      )
      .addCase(loadAssessments.rejected, (state, action) => {
        state.assessmentLoading = false;
        state.assessmentError = action.payload as string;
      })
      .addCase(getAssessmentsLimit.pending, (state) => {
        state.assessmentLoading = true;
        state.assessmentError = null;
      })
      .addCase(
        getAssessmentsLimit.fulfilled,
        (state, action: PayloadAction<LimitResponse<Assessment,"assessments">>) => {
          state.assessmentLoading = false;

          state.assessments = action.payload.assessments;
          state.assessmentsPage = action.payload.meta.page;
          state.assessmentsLimit = action.payload.meta.per_page;
          state.assessmentsTotal = action.payload.meta.total;
          state.assessmentsTotalPages = action.payload.meta.total_pages;
        },
      )
      .addCase(getAssessmentsLimit.rejected, (state, action) => {
        state.assessmentLoading = false;
        state.assessmentError = action.payload as string;
      })
      .addCase(
        createAssessment.fulfilled,
        (state, action: PayloadAction<Assessment>) => {
          state.assessments.push(action.payload);
        },
      )
      .addCase(
        updateAssessment.fulfilled,
        (state, action: PayloadAction<Assessment>) => {
          const idx = state.assessments.findIndex(
            (a) => a.id === action.payload.id,
          );
          if (idx !== -1) state.assessments[idx] = action.payload;
        },
      )
      .addCase(
        removeAssessment.fulfilled,
        (state, action: PayloadAction<{ id: number }>) => {
          state.assessments = state.assessments.filter(
            (a) => a.id !== action.payload.id,
          );
        },
      );
  },
});

export default assessmentSlice.reducer;
