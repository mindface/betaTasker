import { createSlice, createAsyncThunk, PayloadAction } from '@reduxjs/toolkit';

interface QuantificationMetrics {
  reproducibility: number;
  shareability: number;
  standardization: number;
  confidence: number;
}

interface QuantificationData {
  id: string;
  userId: number;
  taskId: number;
  
  // 定量化レベル
  level: 1 | 2 | 3 | 4;
  levelDescription: string;
  
  // 生データと変換
  rawValue: string;
  quantifiedValue: number;
  unit: string;
  variance: number;
  
  // パターン情報
  patternType: '同一事象' | '類似事象' | '類推可能' | '組み合わせ';
  patternConfidence: number;
  
  // メトリクス
  metrics: QuantificationMetrics;
  
  // ドメインコンテキスト
  domain: string;
  relatedDomains: string[];
  transferability: number;
  
  // タイムスタンプ
  createdAt: string;
  updatedAt: string;
}

interface PatternEvolution {
  patternId: string;
  version: number;
  changes: Array<{
    date: string;
    modification: string;
    reason: string;
    impact: number;
  }>;
  currentEffectiveness: number;
}

interface QuantificationState {
  // データ
  quantificationData: QuantificationData[];
  patternEvolutions: PatternEvolution[];
  
  // 集計メトリクス
  aggregateMetrics: {
    averageReproducibility: number;
    averageShareability: number;
    level3OrHigherRate: number;
    totalDataPoints: number;
  };
  
  // 状態管理
  loading: boolean;
  error: string | null;
  lastSync: string | null;
}

const initialState: QuantificationState = {
  quantificationData: [],
  patternEvolutions: [],
  aggregateMetrics: {
    averageReproducibility: 0,
    averageShareability: 0,
    level3OrHigherRate: 0,
    totalDataPoints: 0,
  },
  loading: false,
  error: null,
  lastSync: null,
};

// 非同期アクション
export const collectQuantificationData = createAsyncThunk(
  'quantification/collectData',
  async (params: { userId: number; taskId: number; rawData: any }, { rejectWithValue }) => {
    try {
      const response = await fetch('/api/heuristics/quantification/collect', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(params),
        credentials: 'include',
      });
      
      const data = await response.json();
      if (!response.ok) throw new Error(data.error || '定量化データ収集失敗');
      
      return data;
    } catch (error: any) {
      return rejectWithValue(error.message);
    }
  }
);

export const analyzePattern = createAsyncThunk(
  'quantification/analyzePattern',
  async (params: { dataPoints: QuantificationData[] }, { rejectWithValue }) => {
    try {
      const response = await fetch('/api/heuristics/quantification/analyze', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(params),
        credentials: 'include',
      });
      
      const data = await response.json();
      if (!response.ok) throw new Error(data.error || 'パターン分析失敗');
      
      return data;
    } catch (error: any) {
      return rejectWithValue(error.message);
    }
  }
);

export const calculateMetrics = createAsyncThunk(
  'quantification/calculateMetrics',
  async (params: { userId?: number; period?: string }, { rejectWithValue }) => {
    try {
      const queryParams = new URLSearchParams();
      if (params.userId) queryParams.append('userId', params.userId.toString());
      if (params.period) queryParams.append('period', params.period);
      
      const response = await fetch(`/api/heuristics/quantification/metrics?${queryParams}`, {
        method: 'GET',
        credentials: 'include',
      });
      
      const data = await response.json();
      if (!response.ok) throw new Error(data.error || 'メトリクス計算失敗');
      
      return data;
    } catch (error: any) {
      return rejectWithValue(error.message);
    }
  }
);

const quantificationSlice = createSlice({
  name: 'quantification',
  initialState,
  reducers: {
    // ローカルデータ更新
    addLocalQuantificationData: (state, action: PayloadAction<QuantificationData>) => {
      state.quantificationData.push(action.payload);
      recalculateAggregateMetrics(state);
    },
    
    updateQuantificationLevel: (state, action: PayloadAction<{ id: string; level: 1 | 2 | 3 | 4 }>) => {
      const data = state.quantificationData.find(d => d.id === action.payload.id);
      if (data) {
        data.level = action.payload.level;
        data.levelDescription = getLevelDescription(action.payload.level);
        recalculateAggregateMetrics(state);
      }
    },
    
    recordPatternEvolution: (state, action: PayloadAction<PatternEvolution>) => {
      const existing = state.patternEvolutions.find(p => p.patternId === action.payload.patternId);
      if (existing) {
        existing.version = action.payload.version;
        existing.changes = action.payload.changes;
        existing.currentEffectiveness = action.payload.currentEffectiveness;
      } else {
        state.patternEvolutions.push(action.payload);
      }
    },
    
    clearQuantificationData: (state) => {
      state.quantificationData = [];
      state.aggregateMetrics = initialState.aggregateMetrics;
    },
  },
  
  extraReducers: (builder) => {
    builder
      // collectQuantificationData
      .addCase(collectQuantificationData.pending, (state) => {
        state.loading = true;
        state.error = null;
      })
      .addCase(collectQuantificationData.fulfilled, (state, action) => {
        state.loading = false;
        state.quantificationData.push(action.payload);
        state.lastSync = new Date().toISOString();
        recalculateAggregateMetrics(state);
      })
      .addCase(collectQuantificationData.rejected, (state, action) => {
        state.loading = false;
        state.error = action.payload as string;
      })
      
      // analyzePattern
      .addCase(analyzePattern.fulfilled, (state, action) => {
        if (action.payload.evolution) {
          const evolution = action.payload.evolution as PatternEvolution;
          const existing = state.patternEvolutions.find(p => p.patternId === evolution.patternId);
          if (existing) {
            Object.assign(existing, evolution);
          } else {
            state.patternEvolutions.push(evolution);
          }
        }
      })
      
      // calculateMetrics
      .addCase(calculateMetrics.fulfilled, (state, action) => {
        state.aggregateMetrics = action.payload.metrics;
        state.lastSync = new Date().toISOString();
      });
  },
});

// ヘルパー関数
function getLevelDescription(level: 1 | 2 | 3 | 4): string {
  const descriptions = {
    1: '感覚的定量化',
    2: '部分的定量化',
    3: '構造化定量化',
    4: '体系的定量化',
  };
  return descriptions[level];
}

function recalculateAggregateMetrics(state: QuantificationState): void {
  const data = state.quantificationData;
  if (data.length === 0) {
    state.aggregateMetrics = initialState.aggregateMetrics;
    return;
  }
  
  const totalReproducibility = data.reduce((sum, d) => sum + d.metrics.reproducibility, 0);
  const totalShareability = data.reduce((sum, d) => sum + d.metrics.shareability, 0);
  const level3OrHigher = data.filter(d => d.level >= 3).length;
  
  state.aggregateMetrics = {
    averageReproducibility: totalReproducibility / data.length,
    averageShareability: totalShareability / data.length,
    level3OrHigherRate: level3OrHigher / data.length,
    totalDataPoints: data.length,
  };
}

export const {
  addLocalQuantificationData,
  updateQuantificationLevel,
  recordPatternEvolution,
  clearQuantificationData,
} = quantificationSlice.actions;

export default quantificationSlice.reducer;