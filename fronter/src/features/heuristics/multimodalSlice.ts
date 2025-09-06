import { createSlice, createAsyncThunk, PayloadAction } from '@reduxjs/toolkit';

// 型定義
interface TextFeatures {
  text: string;
  tokens: string[];
  semanticVector: number[];
  ambiguityScore: number;
}

interface ImageFeatures {
  imageUrl: string;
  objects: Array<{
    label: string;
    confidence: number;
    boundingBox: { x: number; y: number; width: number; height: number };
  }>;
  measurements: Array<{
    type: string;
    value: number;
    unit: string;
  }>;
  confidence: number;
}

interface QuantifiedValue {
  value: number;
  unit: string;
  range: [number, number];
  confidence: number;
}

interface MultimodalData {
  id: string;
  userId: number;
  taskId: number;
  
  // 入力データ
  linguistic: TextFeatures;
  visual?: ImageFeatures;
  
  // 関連付け
  association: {
    mappingType: 'direct' | 'inferred' | 'learned';
    correlationScore: number;
    contextRelevance: number;
    historicalAccuracy: number;
  };
  
  // 定量化結果
  quantification: QuantifiedValue;
  
  // メタデータ
  timestamp: string;
  verified: boolean;
  userFeedback?: 'correct' | 'too_high' | 'too_low' | 'incorrect';
}

interface VisualMetaphor {
  id: string;
  metaphor: string;
  referenceObject: string;
  dimensions: {
    width: number;
    height: number;
    depth?: number;
  };
  imageUrl?: string;
  variability: {
    min: number;
    max: number;
  };
}

interface UserCalibration {
  userId: number;
  referenceObject: string;
  measurements: Record<string, number>;
  imageUrl: string;
  timestamp: string;
  confidence: number;
}

interface MultimodalState {
  // データ
  multimodalData: MultimodalData[];
  visualMetaphors: VisualMetaphor[];
  userCalibrations: UserCalibration[];
  
  // 処理中の状態
  currentProcessing: {
    text?: string;
    imageUrl?: string;
    status: 'idle' | 'processing' | 'completed' | 'error';
    result?: QuantifiedValue;
  };
  
  // マッピング辞書
  directMappings: Record<string, QuantifiedValue>;
  
  // 統計
  statistics: {
    totalProcessed: number;
    accuracyRate: number;
    averageConfidence: number;
    userConfirmationRate: number;
  };
  
  // UI状態
  loading: boolean;
  error: string | null;
}

const initialState: MultimodalState = {
  multimodalData: [],
  visualMetaphors: [
    {
      id: 'palm',
      metaphor: '手のひらサイズ',
      referenceObject: 'adult_palm',
      dimensions: { width: 10, height: 18 },
      variability: { min: 0.8, max: 1.2 }
    },
    {
      id: 'card',
      metaphor: '名刺大',
      referenceObject: 'business_card',
      dimensions: { width: 9.1, height: 5.5 },
      variability: { min: 0.95, max: 1.05 }
    },
    {
      id: 'coin500',
      metaphor: '500円玉大',
      referenceObject: '500yen_coin',
      dimensions: { width: 2.65, height: 2.65 },
      variability: { min: 0.98, max: 1.02 }
    }
  ],
  userCalibrations: [],
  currentProcessing: {
    status: 'idle'
  },
  directMappings: {
    '小さじ1杯': { value: 5, unit: 'ml', range: [4.5, 5.5], confidence: 0.9 },
    '大さじ1杯': { value: 15, unit: 'ml', range: [14, 16], confidence: 0.9 },
    'ひとつまみ': { value: 0.5, unit: 'g', range: [0.3, 0.7], confidence: 0.7 },
    '少々': { value: 0.2, unit: 'g', range: [0.1, 0.3], confidence: 0.6 },
    'カップ1杯': { value: 200, unit: 'ml', range: [180, 220], confidence: 0.8 },
  },
  statistics: {
    totalProcessed: 0,
    accuracyRate: 0,
    averageConfidence: 0,
    userConfirmationRate: 0,
  },
  loading: false,
  error: null,
};

// 非同期アクション
export const processMultimodal = createAsyncThunk(
  'multimodal/process',
  async (params: { text: string; imageFile?: File; userId: number; taskId: number }, { rejectWithValue }) => {
    try {
      const formData = new FormData();
      formData.append('text', params.text);
      if (params.imageFile) {
        formData.append('image', params.imageFile);
      }
      formData.append('userId', params.userId.toString());
      formData.append('taskId', params.taskId.toString());
      
      const response = await fetch('/api/heuristics/multimodal/process', {
        method: 'POST',
        body: formData,
        credentials: 'include',
      });
      
      const data = await response.json();
      if (!response.ok) throw new Error(data.error || 'マルチモーダル処理失敗');
      
      return data;
    } catch (error: any) {
      return rejectWithValue(error.message);
    }
  }
);

export const calibrateUser = createAsyncThunk(
  'multimodal/calibrate',
  async (params: { userId: number; referenceObject: string; imageFile: File }, { rejectWithValue }) => {
    try {
      const formData = new FormData();
      formData.append('userId', params.userId.toString());
      formData.append('referenceObject', params.referenceObject);
      formData.append('image', params.imageFile);
      
      const response = await fetch('/api/heuristics/multimodal/calibrate', {
        method: 'POST',
        body: formData,
        credentials: 'include',
      });
      
      const data = await response.json();
      if (!response.ok) throw new Error(data.error || 'キャリブレーション失敗');
      
      return data;
    } catch (error: any) {
      return rejectWithValue(error.message);
    }
  }
);

export const verifyQuantification = createAsyncThunk(
  'multimodal/verify',
  async (params: { dataId: string; feedback: string }, { rejectWithValue }) => {
    try {
      const response = await fetch('/api/heuristics/multimodal/verify', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(params),
        credentials: 'include',
      });
      
      const data = await response.json();
      if (!response.ok) throw new Error(data.error || '検証失敗');
      
      return data;
    } catch (error: any) {
      return rejectWithValue(error.message);
    }
  }
);

const multimodalSlice = createSlice({
  name: 'multimodal',
  initialState,
  reducers: {
    // テキスト処理開始
    startTextProcessing: (state, action: PayloadAction<string>) => {
      state.currentProcessing = {
        text: action.payload,
        status: 'processing'
      };
    },
    
    // 画像追加
    addImage: (state, action: PayloadAction<string>) => {
      state.currentProcessing.imageUrl = action.payload;
    },
    
    // 直接マッピング追加
    addDirectMapping: (state, action: PayloadAction<{ pattern: string; value: QuantifiedValue }>) => {
      state.directMappings[action.payload.pattern] = action.payload.value;
    },
    
    // メタファー追加
    addVisualMetaphor: (state, action: PayloadAction<VisualMetaphor>) => {
      state.visualMetaphors.push(action.payload);
    },
    
    // 統計更新
    updateStatistics: (state) => {
      const data = state.multimodalData;
      if (data.length === 0) return;
      
      const verified = data.filter(d => d.verified);
      const totalConfidence = data.reduce((sum, d) => sum + d.quantification.confidence, 0);
      const confirmations = data.filter(d => d.userFeedback === 'correct');
      
      state.statistics = {
        totalProcessed: data.length,
        accuracyRate: verified.length > 0 ? verified.length / data.length : 0,
        averageConfidence: totalConfidence / data.length,
        userConfirmationRate: confirmations.length / data.length,
      };
    },
    
    // エラークリア
    clearError: (state) => {
      state.error = null;
    },
    
    // 処理リセット
    resetProcessing: (state) => {
      state.currentProcessing = { status: 'idle' };
    },
  },
  
  extraReducers: (builder) => {
    builder
      // processMultimodal
      .addCase(processMultimodal.pending, (state) => {
        state.loading = true;
        state.error = null;
        state.currentProcessing.status = 'processing';
      })
      .addCase(processMultimodal.fulfilled, (state, action) => {
        state.loading = false;
        
        const multimodalData: MultimodalData = {
          id: action.payload.id,
          userId: action.payload.userId,
          taskId: action.payload.taskId,
          linguistic: action.payload.linguistic,
          visual: action.payload.visual,
          association: action.payload.association,
          quantification: action.payload.quantification,
          timestamp: new Date().toISOString(),
          verified: false,
        };
        
        state.multimodalData.push(multimodalData);
        state.currentProcessing = {
          status: 'completed',
          result: action.payload.quantification,
        };
        
        // 統計を更新
        multimodalSlice.caseReducers.updateStatistics(state);
      })
      .addCase(processMultimodal.rejected, (state, action) => {
        state.loading = false;
        state.error = action.payload as string;
        state.currentProcessing.status = 'error';
      })
      
      // calibrateUser
      .addCase(calibrateUser.fulfilled, (state, action) => {
        const calibration: UserCalibration = {
          userId: action.payload.userId,
          referenceObject: action.payload.referenceObject,
          measurements: action.payload.measurements,
          imageUrl: action.payload.imageUrl,
          timestamp: new Date().toISOString(),
          confidence: action.payload.confidence,
        };
        
        // 既存のキャリブレーションを更新または追加
        const index = state.userCalibrations.findIndex(
          c => c.userId === calibration.userId && c.referenceObject === calibration.referenceObject
        );
        
        if (index >= 0) {
          state.userCalibrations[index] = calibration;
        } else {
          state.userCalibrations.push(calibration);
        }
      })
      
      // verifyQuantification
      .addCase(verifyQuantification.fulfilled, (state, action) => {
        const data = state.multimodalData.find(d => d.id === action.payload.dataId);
        if (data) {
          data.verified = true;
          data.userFeedback = action.payload.feedback;
          
          // 統計を更新
          multimodalSlice.caseReducers.updateStatistics(state);
        }
      });
  },
});

export const {
  startTextProcessing,
  addImage,
  addDirectMapping,
  addVisualMetaphor,
  updateStatistics,
  clearError,
  resetProcessing,
} = multimodalSlice.actions;

export default multimodalSlice.reducer;