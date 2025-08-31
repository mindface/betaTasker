import { createSlice, createAsyncThunk, PayloadAction } from '@reduxjs/toolkit';
import * as heuristicsApi from '../../services/heuristicsApi';
import { heuristicsCache, CACHE_KEYS } from '../../services/heuristicsCache';
import {
  HeuristicsAnalysis,
  HeuristicsAnalysisRequest,
  HeuristicsTracking,
  HeuristicsTrackingData,
  HeuristicsInsight,
  HeuristicsPattern,
  HeuristicsModel,
  HeuristicsTrainRequest,
  AnalysisFilters,
  PatternFilters,
  InsightFilters,
  PaginationState,
  DEFAULT_PAGINATION
} from '../../model/heuristics';

// 状態の型定義
interface HeuristicsState {
  // 分析
  analysis: {
    items: HeuristicsAnalysis[];
    current: HeuristicsAnalysis | null;
    loading: boolean;
    error: string | null;
    filters: AnalysisFilters;
    pagination: PaginationState;
  };
  
  // トラッキング
  tracking: {
    data: HeuristicsTracking[];
    loading: boolean;
    error: string | null;
    realtime: boolean;
    sessionId: string | null;
  };
  
  // パターン
  patterns: {
    items: HeuristicsPattern[];
    loading: boolean;
    error: string | null;
    categories: string[];
    confidenceThreshold: number;
    metadata: {
      user_id: string;
      data_type: string;
      period: string;
    } | null;
  };
  
  // インサイト
  insights: {
    items: HeuristicsInsight[];
    current: HeuristicsInsight | null;
    loading: boolean;
    error: string | null;
    total: number;
    filters: InsightFilters;
    pagination: PaginationState;
  };
  
  // モデル
  models: {
    items: HeuristicsModel[];
    current: HeuristicsModel | null;
    loading: boolean;
    error: string | null;
  };
}

const initialState: HeuristicsState = {
  analysis: {
    items: [],
    current: null,
    loading: false,
    error: null,
    filters: {},
    pagination: DEFAULT_PAGINATION
  },
  
  tracking: {
    data: [],
    loading: false,
    error: null,
    realtime: false,
    sessionId: null
  },
  
  patterns: {
    items: [],
    loading: false,
    error: null,
    categories: [],
    confidenceThreshold: 0.7,
    metadata: null
  },
  
  insights: {
    items: [],
    current: null,
    loading: false,
    error: null,
    total: 0,
    filters: {},
    pagination: DEFAULT_PAGINATION
  },
  
  models: {
    items: [],
    current: null,
    loading: false,
    error: null
  }
};

// 分析関連のThunks
export const analyzeData = createAsyncThunk(
  'heuristics/analyze',
  async (request: HeuristicsAnalysisRequest, { rejectWithValue, dispatch }) => {
    try {
      // 分析開始の状態更新
      dispatch(setAnalysisStatus('pending'));
      
      const response = await heuristicsApi.analyzeData(request);
      
      // キャッシュに保存
      heuristicsCache.set(`${CACHE_KEYS.ANALYSIS}_${response.id}`, response);
      
      return response;
      
    } catch (error) {
      console.error('Analysis failed:', error);
      
      if (error instanceof Error) {
        return rejectWithValue({
          code: 'ANALYSIS_FAILED',
          message: error.message,
          details: error
        });
      }
      
      return rejectWithValue({
        code: 'SYSTEM_ERROR',
        message: '分析処理中にエラーが発生しました',
        details: error
      });
    }
  }
);

export const fetchAnalysisById = createAsyncThunk(
  'heuristics/fetchAnalysisById',
  async (id: string, { rejectWithValue }) => {
    try {
      // キャッシュから取得を試行
      const cached = heuristicsCache.get<HeuristicsAnalysis>(`${CACHE_KEYS.ANALYSIS}_${id}`);
      if (cached) {
        return cached;
      }
      
      const response = await heuristicsApi.getAnalysisById(id);
      
      // キャッシュに保存
      heuristicsCache.set(`${CACHE_KEYS.ANALYSIS}_${id}`, response);
      
      return response;
    } catch (error) {
      if (error instanceof Error) {
        return rejectWithValue(error.message);
      }
      return rejectWithValue('分析結果の取得に失敗しました');
    }
  }
);

export const fetchAnalyses = createAsyncThunk(
  'heuristics/fetchAnalyses',
  async (params: { filters?: AnalysisFilters; pagination?: Partial<PaginationState> } | undefined, { rejectWithValue }) => {
    try {
      const response = await heuristicsApi.getAnalyses(params?.filters, params?.pagination);
      
      // キャッシュに保存
      heuristicsCache.set(CACHE_KEYS.ANALYSES, response);
      
      return response;
    } catch (error) {
      if (error instanceof Error) {
        return rejectWithValue(error.message);
      }
      return rejectWithValue('分析一覧の取得に失敗しました');
    }
  }
);

// トラッキング関連のThunks
export const trackUserBehavior = createAsyncThunk(
  'heuristics/track',
  async (trackData: HeuristicsTrackingData, { rejectWithValue }) => {
    try {
      const response = await heuristicsApi.trackBehavior(trackData);
      
      // キャッシュを無効化
      heuristicsCache.invalidate(CACHE_KEYS.TRACKING);
      
      return response;
    } catch (error) {
      if (error instanceof Error) {
        return rejectWithValue(error.message);
      }
      return rejectWithValue('行動追跡に失敗しました');
    }
  }
);

export const fetchTrackingData = createAsyncThunk(
  'heuristics/fetchTrackingData',
  async (userId: string, { rejectWithValue }) => {
    try {
      // キャッシュから取得を試行
      const cached = heuristicsCache.get<HeuristicsTracking[]>(`${CACHE_KEYS.TRACKING}_${userId}`);
      if (cached) {
        return cached;
      }
      
      const response = await heuristicsApi.getTrackingData(userId);
      
      // キャッシュに保存
      heuristicsCache.set(`${CACHE_KEYS.TRACKING}_${userId}`, response);
      
      return response;
    } catch (error) {
      if (error instanceof Error) {
        return rejectWithValue(error.message);
      }
      return rejectWithValue('トラッキングデータの取得に失敗しました');
    }
  }
);

// パターン関連のThunks
export const fetchPatterns = createAsyncThunk(
  'heuristics/fetchPatterns',
  async (params: {
    user_id?: string;
    data_type?: string;
    period?: string;
    filters?: PatternFilters;
  } | undefined, { rejectWithValue }) => {
    try {
      const response = await heuristicsApi.detectPatterns(params);
      
      // キャッシュに保存
      heuristicsCache.set(CACHE_KEYS.PATTERNS, response);
      
      return response;
    } catch (error) {
      if (error instanceof Error) {
        return rejectWithValue(error.message);
      }
      return rejectWithValue('パターン検出に失敗しました');
    }
  }
);

// インサイト関連のThunks
export const fetchInsights = createAsyncThunk(
  'heuristics/fetchInsights',
  async (params: {
    limit?: number;
    offset?: number;
    user_id?: string;
    filters?: InsightFilters;
  } | undefined, { rejectWithValue }) => {
    try {
      const response = await heuristicsApi.fetchInsights(params);
      
      // キャッシュに保存
      heuristicsCache.set(CACHE_KEYS.INSIGHTS, response);
      
      return response;
    } catch (error) {
      if (error instanceof Error) {
        return rejectWithValue(error.message);
      }
      return rejectWithValue('インサイトの取得に失敗しました');
    }
  }
);

export const fetchInsightById = createAsyncThunk(
  'heuristics/fetchInsightById',
  async (id: string, { rejectWithValue }) => {
    try {
      // キャッシュから取得を試行
      const cached = heuristicsCache.get<HeuristicsInsight>(`${CACHE_KEYS.INSIGHT}_${id}`);
      if (cached) {
        return cached;
      }
      
      const response = await heuristicsApi.getInsightById(id);
      
      // キャッシュに保存
      heuristicsCache.set(`${CACHE_KEYS.INSIGHT}_${id}`, response);
      
      return response;
    } catch (error) {
      if (error instanceof Error) {
        return rejectWithValue(error.message);
      }
      return rejectWithValue('インサイトの取得に失敗しました');
    }
  }
);

// モデル関連のThunks
export const trainModel = createAsyncThunk(
  'heuristics/trainModel',
  async (request: HeuristicsTrainRequest, { rejectWithValue }) => {
    try {
      const response = await heuristicsApi.trainModel(request);
      
      // キャッシュを無効化
      heuristicsCache.invalidate(CACHE_KEYS.MODELS);
      
      return response;
    } catch (error) {
      if (error instanceof Error) {
        return rejectWithValue(error.message);
      }
      return rejectWithValue('モデル学習に失敗しました');
    }
  }
);

// Slice定義
const heuristicsSlice = createSlice({
  name: 'heuristics',
  initialState,
  reducers: {
    // 分析関連
    setAnalysisStatus: (state, action: PayloadAction<'pending' | 'completed' | 'failed'>) => {
      state.analysis.current = state.analysis.current ? {
        ...state.analysis.current,
        status: action.payload
      } : null;
    },
    setAnalysisFilters: (state, action: PayloadAction<AnalysisFilters>) => {
      state.analysis.filters = action.payload;
    },
    clearAnalysisError: (state) => {
      state.analysis.error = null;
    },
    
    // トラッキング関連
    setTrackingRealtime: (state, action: PayloadAction<boolean>) => {
      state.tracking.realtime = action.payload;
    },
    setSessionId: (state, action: PayloadAction<string>) => {
      state.tracking.sessionId = action.payload;
    },
    clearTrackingError: (state) => {
      state.tracking.error = null;
    },
    
    // パターン関連
    setConfidenceThreshold: (state, action: PayloadAction<number>) => {
      state.patterns.confidenceThreshold = action.payload;
    },
    clearPatternsError: (state) => {
      state.patterns.error = null;
    },
    
    // インサイト関連
    setInsightFilters: (state, action: PayloadAction<InsightFilters>) => {
      state.insights.filters = action.payload;
    },
    clearInsightsError: (state) => {
      state.insights.error = null;
    },
    
    // 共通
    clearAllErrors: (state) => {
      state.analysis.error = null;
      state.tracking.error = null;
      state.patterns.error = null;
      state.insights.error = null;
      state.models.error = null;
    },
    resetState: () => initialState,
    // 新しく追加する関数
    addPattern: (state, action: PayloadAction<HeuristicsPattern>) => {
      state.patterns.items.unshift(action.payload);
      state.patterns.categories = Array.from(new Set([...state.patterns.categories, action.payload.category]));
    },
    addInsight: (state, action: PayloadAction<HeuristicsInsight>) => {
      state.insights.items.unshift(action.payload);
      state.insights.total += 1;
    },
    updateAnalysis: (state, action: PayloadAction<HeuristicsAnalysis>) => {
      const index = state.analysis.items.findIndex(item => item.id === action.payload.id);
      if (index !== -1) {
        state.analysis.items[index] = action.payload;
      }
      if (state.analysis.current?.id === action.payload.id) {
        state.analysis.current = action.payload;
      }
    }
  },
  extraReducers: (builder) => {
    // 分析関連
    builder
      .addCase(analyzeData.pending, (state) => {
        state.analysis.loading = true;
        state.analysis.error = null;
      })
      .addCase(analyzeData.fulfilled, (state, action) => {
        state.analysis.loading = false;
        state.analysis.current = action.payload;
        state.analysis.items.unshift(action.payload);
        state.analysis.pagination.total += 1;
      })
      .addCase(analyzeData.rejected, (state, action) => {
        state.analysis.loading = false;
        state.analysis.error = action.payload as string;
      })
      .addCase(fetchAnalysisById.pending, (state) => {
        state.analysis.loading = true;
        state.analysis.error = null;
      })
      .addCase(fetchAnalysisById.fulfilled, (state, action) => {
        state.analysis.loading = false;
        state.analysis.current = action.payload;
      })
      .addCase(fetchAnalysisById.rejected, (state, action) => {
        state.analysis.loading = false;
        state.analysis.error = action.payload as string;
      })
      .addCase(fetchAnalyses.pending, (state) => {
        state.analysis.loading = true;
        state.analysis.error = null;
      })
      .addCase(fetchAnalyses.fulfilled, (state, action) => {
        state.analysis.loading = false;
        state.analysis.items = action.payload.analyses;
        state.analysis.pagination = action.payload.pagination;
      })
      .addCase(fetchAnalyses.rejected, (state, action) => {
        state.analysis.loading = false;
        state.analysis.error = action.payload as string;
      });
    
    // トラッキング関連
    builder
      .addCase(trackUserBehavior.pending, (state) => {
        state.tracking.loading = true;
        state.tracking.error = null;
      })
      .addCase(trackUserBehavior.fulfilled, (state) => {
        state.tracking.loading = false;
      })
      .addCase(trackUserBehavior.rejected, (state, action) => {
        state.tracking.loading = false;
        state.tracking.error = action.payload as string;
      })
      .addCase(fetchTrackingData.pending, (state) => {
        state.tracking.loading = true;
        state.tracking.error = null;
      })
      .addCase(fetchTrackingData.fulfilled, (state, action) => {
        state.tracking.loading = false;
        state.tracking.data = action.payload;
      })
      .addCase(fetchTrackingData.rejected, (state, action) => {
        state.tracking.loading = false;
        state.tracking.error = action.payload as string;
      });
    
    // パターン関連
    builder
      .addCase(fetchPatterns.pending, (state) => {
        state.patterns.loading = true;
        state.patterns.error = null;
      })
      .addCase(fetchPatterns.fulfilled, (state, action) => {
        state.patterns.loading = false;
        state.patterns.items = action.payload.patterns;
        state.patterns.metadata = action.payload.metadata;
        
        // カテゴリの抽出
        const categories = Array.from(new Set(action.payload.patterns.map(p => p.category)));
        state.patterns.categories = categories;
      })
      .addCase(fetchPatterns.rejected, (state, action) => {
        state.patterns.loading = false;
        state.patterns.error = action.payload as string;
      });
    
    // インサイト関連
    builder
      .addCase(fetchInsights.pending, (state) => {
        state.insights.loading = true;
        state.insights.error = null;
      })
      .addCase(fetchInsights.fulfilled, (state, action) => {
        state.insights.loading = false;
        state.insights.items = action.payload.insights;
        state.insights.total = action.payload.total;
        state.insights.pagination = {
          page: Math.floor(action.payload.offset / action.payload.limit) + 1,
          limit: action.payload.limit,
          total: action.payload.total,
          totalPages: Math.ceil(action.payload.total / action.payload.limit)
        };
      })
      .addCase(fetchInsights.rejected, (state, action) => {
        state.insights.loading = false;
        state.insights.error = action.payload as string;
      })
      .addCase(fetchInsightById.pending, (state) => {
        state.insights.loading = true;
        state.insights.error = null;
      })
      .addCase(fetchInsightById.fulfilled, (state, action) => {
        state.insights.loading = false;
        state.insights.current = action.payload;
      })
      .addCase(fetchInsightById.rejected, (state, action) => {
        state.insights.loading = false;
        state.insights.error = action.payload as string;
      });
    
    // モデル関連
    builder
      .addCase(trainModel.pending, (state) => {
        state.models.loading = true;
        state.models.error = null;
      })
      .addCase(trainModel.fulfilled, (state, action) => {
        state.models.loading = false;
        state.models.current = action.payload;
        state.models.items.unshift(action.payload);
      })
      .addCase(trainModel.rejected, (state, action) => {
        state.models.loading = false;
        state.models.error = action.payload as string;
      });
  },
});

export const {
  setAnalysisStatus,
  setAnalysisFilters,
  clearAnalysisError,
  setTrackingRealtime,
  setSessionId,
  clearTrackingError,
  setConfidenceThreshold,
  clearPatternsError,
  setInsightFilters,
  clearInsightsError,
  clearAllErrors,
  resetState,
  // 新しく追加する関数
  addPattern,
  addInsight,
  updateAnalysis
} = heuristicsSlice.actions;

export default heuristicsSlice.reducer;