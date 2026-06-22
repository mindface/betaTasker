import { createSlice, createAsyncThunk, PayloadAction } from '@reduxjs/toolkit';
import { 
  analysisAdapter, 
  AnalysisEntityState, 
  initialAnalysisEntityState,
  analysisEntityHelpers
} from '../entities/analysis';
import * as heuristicsApi from '../../../services/heuristicsApi';
import { HeuristicsAnalysis, HeuristicsAnalysisRequest } from '../../../model/heuristics';

// Collection状態の定義
interface AnalysisCollectionState {
  ids: string[];
  pagination: {
    limit: number;
    offset: number;
    hasMore: boolean;
  };
  filters: {
    userId?: number;
    taskId?: number;
    analysisType?: string;
    status?: string;
    dateFrom?: string;
    dateTo?: string;
  };
  loading: boolean;
  error: string | null;
  lastFetch: string | null;
}

// UI状態の定義
interface AnalysisUIState {
  selectedAnalysisId: number | null;
  showAnalysisForm: boolean;
  showFetchForm: boolean;
  currentTab: 'all' | 'completed' | 'pending' | 'failed';
  sortBy: 'created_at' | 'score' | 'status';
  sortOrder: 'asc' | 'desc';
}

// 統合状態の定義
interface AnalysisState {
  entities: AnalysisEntityState;
  collections: {
    all: AnalysisCollectionState;
    byUser: Record<number, AnalysisCollectionState>;
    byTask: Record<number, AnalysisCollectionState>;
    recent: AnalysisCollectionState;
  };
  ui: AnalysisUIState;
  currentAnalysis: HeuristicsAnalysis | null;
}

const initialCollectionState: AnalysisCollectionState = {
  ids: [],
  pagination: { limit: 20, offset: 0, hasMore: true },
  filters: {},
  loading: false,
  error: null,
  lastFetch: null
};

const initialUIState: AnalysisUIState = {
  selectedAnalysisId: null,
  showAnalysisForm: false,
  showFetchForm: false,
  currentTab: 'all',
  sortBy: 'created_at',
  sortOrder: 'desc'
};

const initialState: AnalysisState = {
  entities: initialAnalysisEntityState,
  collections: {
    all: { ...initialCollectionState },
    byUser: {},
    byTask: {},
    recent: { ...initialCollectionState, pagination: { ...initialCollectionState.pagination, limit: 10 } }
  },
  ui: initialUIState,
  currentAnalysis: null
};

// Async Thunks
export const analyzeData = createAsyncThunk(
  'analysis/analyzeData',
  async (request: HeuristicsAnalysisRequest, { rejectWithValue }) => {
    const response = await heuristicsApi.analyzeData(request);
    if ('error' in response) {
      return rejectWithValue(response.error);
    }
    return response as HeuristicsAnalysis;
  }
);

export const fetchAnalysisById = createAsyncThunk(
  'analysis/fetchById',
  async (id: string, { rejectWithValue }) => {
    const response = await heuristicsApi.getAnalysisById(id);
    if ('error' in response) {
      return rejectWithValue(response.error);
    }
    return response as HeuristicsAnalysis;
  }
);

export const fetchAnalyses = createAsyncThunk(
  'analysis/fetchList',
  async (params: {
    filters?: AnalysisCollectionState['filters'];
    pagination?: { limit: number; offset: number };
  } = {}, { rejectWithValue }) => {
    // Note: 実際のAPI実装後に更新
    try {
      // 暫定的な実装
      const mockAnalyses: HeuristicsAnalysis[] = [];
      return {
        analyses: mockAnalyses,
        total: 0,
        hasMore: false
      };
    } catch (error: any) {
      return rejectWithValue(error.message);
    }
  }
);

export const fetchUserAnalyses = createAsyncThunk(
  'analysis/fetchByUser',
  async (params: {
    userId: number;
    pagination?: { limit: number; offset: number };
  }, { rejectWithValue }) => {
    try {
      // 実際のAPI実装後に更新
      const mockAnalyses: HeuristicsAnalysis[] = [];
      return {
        analyses: mockAnalyses,
        total: 0,
        hasMore: false,
        userId: params.userId
      };
    } catch (error: any) {
      return rejectWithValue(error.message);
    }
  }
);

// Slice定義
const analysisSlice = createSlice({
  name: 'analysis',
  initialState,
  reducers: {
    // UI状態の管理
    setSelectedAnalysis: (state, action: PayloadAction<number | null>) => {
      state.ui.selectedAnalysisId = action.payload;
    },
    
    setShowAnalysisForm: (state, action: PayloadAction<boolean>) => {
      state.ui.showAnalysisForm = action.payload;
    },
    
    setShowFetchForm: (state, action: PayloadAction<boolean>) => {
      state.ui.showFetchForm = action.payload;
    },
    
    setCurrentTab: (state, action: PayloadAction<AnalysisUIState['currentTab']>) => {
      state.ui.currentTab = action.payload;
    },
    
    setSorting: (state, action: PayloadAction<{
      sortBy: AnalysisUIState['sortBy'];
      sortOrder: AnalysisUIState['sortOrder'];
    }>) => {
      state.ui.sortBy = action.payload.sortBy;
      state.ui.sortOrder = action.payload.sortOrder;
    },
    
    // フィルタ管理
    setFilters: (state, action: PayloadAction<AnalysisCollectionState['filters']>) => {
      state.collections.all.filters = { ...state.collections.all.filters, ...action.payload };
    },
    
    clearFilters: (state) => {
      state.collections.all.filters = {};
    },
    
    // ページネーション管理
    setPagination: (state, action: PayloadAction<{ limit?: number; offset?: number }>) => {
      if (action.payload.limit !== undefined) {
        state.collections.all.pagination.limit = action.payload.limit;
      }
      if (action.payload.offset !== undefined) {
        state.collections.all.pagination.offset = action.payload.offset;
      }
    },
    
    resetPagination: (state) => {
      state.collections.all.pagination = { limit: 20, offset: 0, hasMore: true };
    },
    
    // エラー管理
    clearError: (state) => {
      state.collections.all.error = null;
    },
    
    // エンティティ直接操作
    addAnalysis: (state, action: PayloadAction<HeuristicsAnalysis>) => {
      analysisEntityHelpers.addOne(state.entities, action.payload);
      state.collections.all.ids.unshift(String(action.payload.id));
    },
    
    updateAnalysis: (state, action: PayloadAction<HeuristicsAnalysis>) => {
      analysisEntityHelpers.updateOne(state.entities, {
        id: action.payload.id,
        changes: action.payload
      });
    },
    
    removeAnalysis: (state, action: PayloadAction<number>) => {
      analysisEntityHelpers.removeOne(state.entities, action.payload);
      state.collections.all.ids = state.collections.all.ids.filter(id => id !== String(action.payload));
    },
    
    // キャッシュ管理
    markCacheExpired: (state, action: PayloadAction<number[]>) => {
      analysisEntityHelpers.markAsExpired(state.entities, action.payload);
    },
    
    clearExpiredCache: (state) => {
      analysisEntityHelpers.removeExpired(state.entities);
    },
    
    // 状態リセット
    reset: (state) => {
      return initialState;
    }
  },
  
  extraReducers: (builder) => {
    // analyzeData
    builder
      .addCase(analyzeData.pending, (state) => {
        state.collections.all.loading = true;
        state.collections.all.error = null;
      })
      .addCase(analyzeData.fulfilled, (state, action) => {
        state.collections.all.loading = false;
        state.currentAnalysis = action.payload;
        analysisEntityHelpers.addOne(state.entities, action.payload);
        state.collections.all.ids.unshift(String(action.payload.id));
      })
      .addCase(analyzeData.rejected, (state, action) => {
        state.collections.all.loading = false;
        state.collections.all.error = action.payload as string;
      });

    // fetchAnalysisById
    builder
      .addCase(fetchAnalysisById.pending, (state) => {
        state.collections.all.loading = true;
        state.collections.all.error = null;
      })
      .addCase(fetchAnalysisById.fulfilled, (state, action) => {
        state.collections.all.loading = false;
        state.currentAnalysis = action.payload;
        analysisEntityHelpers.upsertOne(state.entities, action.payload);
      })
      .addCase(fetchAnalysisById.rejected, (state, action) => {
        state.collections.all.loading = false;
        state.collections.all.error = action.payload as string;
      });

    // fetchAnalyses
    builder
      .addCase(fetchAnalyses.pending, (state) => {
        state.collections.all.loading = true;
        state.collections.all.error = null;
      })
      .addCase(fetchAnalyses.fulfilled, (state, action) => {
        state.collections.all.loading = false;
        state.collections.all.lastFetch = new Date().toISOString();
        
        if (action.payload.analyses.length > 0) {
          analysisEntityHelpers.upsertMany(state.entities, action.payload.analyses);
          
          if (state.collections.all.pagination.offset === 0) {
            // 新しいフェッチの場合、IDリストをリセット
            state.collections.all.ids = action.payload.analyses.map(a => String(a.id));
          } else {
            // ページング追加の場合、IDを追加
            const newIds = action.payload.analyses.map(a => String(a.id));
            state.collections.all.ids.push(...newIds);
          }
        }
        
        state.collections.all.pagination.hasMore = action.payload.hasMore;
      })
      .addCase(fetchAnalyses.rejected, (state, action) => {
        state.collections.all.loading = false;
        state.collections.all.error = action.payload as string;
      });

    // fetchUserAnalyses
    builder
      .addCase(fetchUserAnalyses.fulfilled, (state, action) => {
        const { userId, analyses, total, hasMore } = action.payload;
        
        if (!state.collections.byUser[userId]) {
          state.collections.byUser[userId] = { ...initialCollectionState };
        }
        
        state.collections.byUser[userId].loading = false;
        state.collections.byUser[userId].lastFetch = new Date().toISOString();
        
        if (analyses.length > 0) {
          analysisEntityHelpers.upsertMany(state.entities, analyses);
          state.collections.byUser[userId].ids = analyses.map(a => String(a.id));
        }
        
        state.collections.byUser[userId].pagination.hasMore = hasMore;
      });
  }
});

export const {
  setSelectedAnalysis,
  setShowAnalysisForm,
  setShowFetchForm,
  setCurrentTab,
  setSorting,
  setFilters,
  clearFilters,
  setPagination,
  resetPagination,
  clearError,
  addAnalysis,
  updateAnalysis,
  removeAnalysis,
  markCacheExpired,
  clearExpiredCache,
  reset
} = analysisSlice.actions;

export default analysisSlice.reducer;