import { createSlice, createAsyncThunk, PayloadAction } from '@reduxjs/toolkit';
import * as heuristicsApi from '../../client/heuristicsApi';
import {
  HeuristicsAnalysis,
  HeuristicsAnalysisRequest,
  HeuristicsTracking,
  HeuristicsTrackingData,
  HeuristicsInsight,
  HeuristicsPattern,
  HeuristicsModel,
  HeuristicsTrainRequest
} from '../../model/heuristics';

interface HeuristicsState {
  // 分析関連
  analyses: HeuristicsAnalysis[];
  currentAnalysis: HeuristicsAnalysis | null;
  analysisLoading: boolean;
  analysisError: string | null;
  
  // トラッキング関連
  trackingData: HeuristicsTracking[];
  trackingLoading: boolean;
  trackingError: string | null;
  
  // インサイト関連
  insights: HeuristicsInsight[];
  currentInsight: HeuristicsInsight | null;
  insightsTotal: number;
  insightsLoading: boolean;
  insightsError: string | null;
  
  // パターン関連
  patterns: HeuristicsPattern[];
  patternsLoading: boolean;
  patternsError: string | null;
  
  // モデル関連
  currentModel: HeuristicsModel | null;
  modelLoading: boolean;
  modelError: string | null;
}

const initialState: HeuristicsState = {
  analyses: [],
  currentAnalysis: null,
  analysisLoading: false,
  analysisError: null,
  
  trackingData: [],
  trackingLoading: false,
  trackingError: null,
  
  insights: [],
  currentInsight: null,
  insightsTotal: 0,
  insightsLoading: false,
  insightsError: null,
  
  patterns: [],
  patternsLoading: false,
  patternsError: null,
  
  currentModel: null,
  modelLoading: false,
  modelError: null,
};

// 分析関連のThunks
export const analyzeData = createAsyncThunk(
  'heuristics/analyze',
  async (request: HeuristicsAnalysisRequest, { rejectWithValue }) => {
    const response = await heuristicsApi.analyzeData(request);
    if ('error' in response) {
      return rejectWithValue(response.error);
    }
    return response.value as HeuristicsAnalysis[];
  }
);

export const fetchAnalysisById = createAsyncThunk(
  'heuristics/fetchAnalysisById',
  async (id: string, { rejectWithValue }) => {
    const response = await heuristicsApi.getAnalysisById(id);
    if ('error' in response) {
      return rejectWithValue(response.error);
    }
    return response as HeuristicsAnalysis;
  }
);

// トラッキング関連のThunks
export const trackUserBehavior = createAsyncThunk(
  'heuristics/track',
  async (trackData: HeuristicsTrackingData, { rejectWithValue }) => {
    const response = await heuristicsApi.trackBehavior(trackData);
    if ('error' in response) {
      return rejectWithValue(response.error);
    }
    if ('data' in response) {
      return response.data;
    }
    return response;
  }
);

export const fetchTrackingData = createAsyncThunk(
  'heuristics/fetchTrackingData',
  async (userId: string, { rejectWithValue }) => {
    const response = await heuristicsApi.getTrackingData();
    if ('error' in response) {
      return rejectWithValue(response.error);
    }
    return response.value as HeuristicsTracking[];
  }
);

// インサイト関連のThunks
export const loadInsights = createAsyncThunk(
  'heuristics/loadInsights',
  async (params: { limit?: number; offset?: number; user_id?: string } | undefined, { rejectWithValue }) => {
    const response = await heuristicsApi.fetchInsights(params);
    if ('error' in response) {
      return rejectWithValue(response.error);
    }
    return { insights: response.value, total: response.value.length } as { insights: HeuristicsInsight[]; total: number };
  }
);

export const fetchInsightById = createAsyncThunk(
  'heuristics/fetchInsightById',
  async (id: string, { rejectWithValue }) => {
    const response = await heuristicsApi.getInsightById(id);
    if ('error' in response) {
      return rejectWithValue(response.error);
    }
    return response.value as HeuristicsInsight;
  }
);

// パターン検出関連のThunks
export const loadPatterns = createAsyncThunk(
  'heuristics/loadPatterns',
  async (params: { user_id?: string; data_type?: string; period?: string } | undefined, { rejectWithValue }) => {
    const response = await heuristicsApi.detectPatterns(params);
    if ('error' in response) {
      return rejectWithValue(response.error);
    }
    return response.value;
  }
);

// モデルトレーニング関連のThunks
export const trainHeuristicsModel = createAsyncThunk(
  'heuristics/trainModel',
  async (request: HeuristicsTrainRequest, { rejectWithValue }) => {
    const response = await heuristicsApi.trainModel(request);
    if ('error' in response) {
      return rejectWithValue(response.error);
    }
    return response.value;
  }
);

const heuristicsSlice = createSlice({
  name: 'heuristics',
  initialState,
  reducers: {
    clearAnalysisError: (state) => {
      state.analysisError = null;
    },
    clearTrackingError: (state) => {
      state.trackingError = null;
    },
    clearInsightsError: (state) => {
      state.insightsError = null;
    },
    clearPatternsError: (state) => {
      state.patternsError = null;
    },
    clearModelError: (state) => {
      state.modelError = null;
    },
  },
  extraReducers: (builder) => {
    // 分析関連
    builder
      .addCase(analyzeData.pending, (state) => {
        state.analysisLoading = true;
        state.analysisError = null;
      })
      .addCase(analyzeData.fulfilled, (state, action) => {
        state.analysisLoading = false;
        state.analyses = action.payload;
        // state.analyses.unshift(action.payload);
      })
      .addCase(analyzeData.rejected, (state, action) => {
        state.analysisLoading = false;
        state.analysisError = action.payload as string;
      })
      .addCase(fetchAnalysisById.fulfilled, (state, action) => {
        state.currentAnalysis = action.payload;
      });

    // トラッキング関連
    builder
      .addCase(trackUserBehavior.pending, (state) => {
        state.trackingLoading = true;
        state.trackingError = null;
      })
      .addCase(trackUserBehavior.fulfilled, (state) => {
        state.trackingLoading = false;
      })
      .addCase(trackUserBehavior.rejected, (state, action) => {
        state.trackingLoading = false;
        state.trackingError = action.payload as string;
      })
      .addCase(fetchTrackingData.fulfilled, (state, action) => {
        state.trackingData = action.payload;
      });
    
    // インサイト関連
    builder
      .addCase(loadInsights.pending, (state) => {
        state.insightsLoading = true;
        state.insightsError = null;
      })
      .addCase(loadInsights.fulfilled, (state, action) => {
        state.insightsLoading = false;
        state.insights = action.payload.insights;
        state.insightsTotal = action.payload.total;
      })
      .addCase(loadInsights.rejected, (state, action) => {
        state.insightsLoading = false;
        state.insightsError = action.payload as string;
      })
      .addCase(fetchInsightById.fulfilled, (state, action) => {
        state.currentInsight = action.payload;
      });

    // パターン検出関連
    builder
      .addCase(loadPatterns.pending, (state) => {
        state.patternsLoading = true;
        state.patternsError = null;
      })
      .addCase(loadPatterns.fulfilled, (state, action) => {
        state.patternsLoading = false;
        state.patterns = action.payload;
      })
      .addCase(loadPatterns.rejected, (state, action) => {
        state.patternsLoading = false;
        state.patternsError = action.payload as string;
      });

    // モデルトレーニング関連
    builder
      .addCase(trainHeuristicsModel.pending, (state) => {
        state.modelLoading = true;
        state.modelError = null;
      })
      .addCase(trainHeuristicsModel.fulfilled, (state, action) => {
        state.modelLoading = false;
        state.currentModel = action.payload;
      })
      .addCase(trainHeuristicsModel.rejected, (state, action) => {
        state.modelLoading = false;
        state.modelError = action.payload as string;
      });
  },
});

export const {
  clearAnalysisError,
  clearTrackingError,
  clearInsightsError,
  clearPatternsError,
  clearModelError,
} = heuristicsSlice.actions;

export default heuristicsSlice.reducer;