import { createSlice, createAsyncThunk, PayloadAction } from "@reduxjs/toolkit";
import { LearningData } from "../../model/learning";
import { fetchLearningData } from "../../client/learnApi";

interface LearningDataState {
  learningData: LearningData | null;
  learningLoading: boolean;
  learningError: string | null;
}

const initialState: LearningDataState = {
  learningData: null,
  learningLoading: false,
  learningError: null,
};

export const loadLearningData = createAsyncThunk(
  "learningData/loadLearningData",
  async (_, { rejectWithValue }) => {
    const response = await fetchLearningData();
    if ("error" in response) {
      return rejectWithValue(response.error);
    }
    return response.value;
  },
);

const learningDataSlice = createSlice({
  name: "learningData",
  initialState,
  reducers: {},
  extraReducers: (builder) => {
    builder
      .addCase(loadLearningData.pending, (state) => {
        state.learningLoading = true;
        state.learningError = null;
      })
      .addCase(
        loadLearningData.fulfilled,
        (state, action: PayloadAction<LearningData | { error: string }>) => {
          state.learningLoading = false;
          if ("error" in action.payload) {
            state.learningError = action.payload.error;
            state.learningData = null;
          } else {
            state.learningData = action.payload;
            state.learningError = null;
          }
        },
      )
      .addCase(loadLearningData.rejected, (state, action) => {
        state.learningLoading = false;
        state.learningError = action.payload as string;
      });
  },
});

export default learningDataSlice.reducer;
