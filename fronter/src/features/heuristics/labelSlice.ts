import { createSlice, createAsyncThunk, PayloadAction } from '@reduxjs/toolkit';
import { 
  QuantificationLabel, 
  LabelDataset, 
  LabelRelation, 
  LabelStatistics,
  CreateLabelRequest,
  UpdateLabelRequest,
  VerifyLabelRequest,
  LabelSearchQuery
} from '../../model/quantificationLabel';

interface LabelState {
  // ラベルデータ
  labels: QuantificationLabel[];
  datasets: LabelDataset[];
  relations: LabelRelation[];
  
  // 現在の選択/編集
  currentLabel: QuantificationLabel | null;
  currentDataset: LabelDataset | null;
  
  // 検索・フィルタ
  searchQuery: LabelSearchQuery;
  searchResults: QuantificationLabel[];
  
  // 統計情報
  statistics: LabelStatistics | null;
  
  // ラベル作成/編集状態
  labelEditor: {
    mode: 'create' | 'edit' | 'view' | null;
    originalText: string;
    imagePreview: string | null;
    annotations: any[];
    concepts: string[];
    suggestedValues: Array<{
      value: number;
      unit: string;
      confidence: number;
      source: string;
    }>;
  };
  
  // UI状態
  loading: boolean;
  error: string | null;
  selectedLabels: string[];
  bulkOperation: {
    type: 'verify' | 'delete' | 'export' | null;
    progress: number;
    total: number;
  };
}

const initialState: LabelState = {
  labels: [],
  datasets: [],
  relations: [],
  currentLabel: null,
  currentDataset: null,
  searchQuery: {},
  searchResults: [],
  statistics: null,
  labelEditor: {
    mode: null,
    originalText: '',
    imagePreview: null,
    annotations: [],
    concepts: [],
    suggestedValues: [],
  },
  loading: false,
  error: null,
  selectedLabels: [],
  bulkOperation: {
    type: null,
    progress: 0,
    total: 0,
  },
};

// 非同期アクション
export const createLabel = createAsyncThunk(
  'label/create',
  async (request: CreateLabelRequest, { rejectWithValue }) => {
    try {
      const formData = new FormData();
      formData.append('text', request.text);
      formData.append('description', request.description);
      formData.append('value', request.value.toString());
      formData.append('unit', request.unit);
      formData.append('domain', request.domain);
      formData.append('category', request.category);
      
      if (request.imageFile) {
        formData.append('image', request.imageFile);
      }
      
      if (request.imageUrl) {
        formData.append('imageUrl', request.imageUrl);
      }
      
      if (request.concepts) {
        formData.append('concepts', JSON.stringify(request.concepts));
      }
      
      if (request.tags) {
        formData.append('tags', JSON.stringify(request.tags));
      }
      
      const response = await fetch('/api/heuristics/labels', {
        method: 'POST',
        body: formData,
        credentials: 'include',
      });
      
      const data = await response.json();
      if (!response.ok) throw new Error(data.error || 'ラベル作成失敗');
      
      return data;
    } catch (error: any) {
      return rejectWithValue(error.message);
    }
  }
);

export const updateLabel = createAsyncThunk(
  'label/update',
  async (request: UpdateLabelRequest, { rejectWithValue }) => {
    try {
      const response = await fetch(`/api/heuristics/labels/${request.id}`, {
        method: 'PUT',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(request),
        credentials: 'include',
      });
      
      const data = await response.json();
      if (!response.ok) throw new Error(data.error || 'ラベル更新失敗');
      
      return data;
    } catch (error: any) {
      return rejectWithValue(error.message);
    }
  }
);

export const verifyLabel = createAsyncThunk(
  'label/verify',
  async (request: VerifyLabelRequest, { rejectWithValue }) => {
    try {
      const response = await fetch(`/api/heuristics/labels/${request.labelId}/verify`, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(request),
        credentials: 'include',
      });
      
      const data = await response.json();
      if (!response.ok) throw new Error(data.error || 'ラベル検証失敗');
      
      return data;
    } catch (error: any) {
      return rejectWithValue(error.message);
    }
  }
);

export const searchLabels = createAsyncThunk(
  'label/search',
  async (query: LabelSearchQuery, { rejectWithValue }) => {
    try {
      const queryParams = new URLSearchParams();
      
      if (query.text) queryParams.append('text', query.text);
      if (query.domain) queryParams.append('domain', query.domain);
      if (query.category) queryParams.append('category', query.category);
      if (query.minConfidence) queryParams.append('minConfidence', query.minConfidence.toString());
      if (query.verified !== undefined) queryParams.append('verified', query.verified.toString());
      if (query.limit) queryParams.append('limit', query.limit.toString());
      if (query.offset) queryParams.append('offset', query.offset.toString());
      if (query.sortBy) queryParams.append('sortBy', query.sortBy);
      if (query.sortOrder) queryParams.append('sortOrder', query.sortOrder);
      
      if (query.valueRange) {
        queryParams.append('minValue', query.valueRange.min.toString());
        queryParams.append('maxValue', query.valueRange.max.toString());
        queryParams.append('unit', query.valueRange.unit);
      }
      
      if (query.concepts) {
        queryParams.append('concepts', JSON.stringify(query.concepts));
      }
      
      if (query.dateRange) {
        queryParams.append('from', query.dateRange.from);
        queryParams.append('to', query.dateRange.to);
      }
      
      const response = await fetch(`/api/heuristics/labels/search?${queryParams}`, {
        method: 'GET',
        credentials: 'include',
      });
      
      const data = await response.json();
      if (!response.ok) throw new Error(data.error || 'ラベル検索失敗');
      
      return data;
    } catch (error: any) {
      return rejectWithValue(error.message);
    }
  }
);

export const loadStatistics = createAsyncThunk(
  'label/loadStatistics',
  async (_, { rejectWithValue }) => {
    try {
      const response = await fetch('/api/heuristics/labels/statistics', {
        method: 'GET',
        credentials: 'include',
      });
      
      const data = await response.json();
      if (!response.ok) throw new Error(data.error || '統計取得失敗');
      
      return data;
    } catch (error: any) {
      return rejectWithValue(error.message);
    }
  }
);

export const suggestQuantification = createAsyncThunk(
  'label/suggest',
  async (params: { text: string; imageUrl?: string; domain?: string }, { rejectWithValue }) => {
    try {
      const response = await fetch('/api/heuristics/labels/suggest', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(params),
        credentials: 'include',
      });
      
      const data = await response.json();
      if (!response.ok) throw new Error(data.error || '定量化提案失敗');
      
      return data;
    } catch (error: any) {
      return rejectWithValue(error.message);
    }
  }
);

const labelSlice = createSlice({
  name: 'label',
  initialState,
  reducers: {
    // ラベル選択・編集
    setCurrentLabel: (state, action: PayloadAction<QuantificationLabel | null>) => {
      state.currentLabel = action.payload;
    },
    
    selectLabel: (state, action: PayloadAction<string>) => {
      const id = action.payload;
      if (state.selectedLabels.includes(id)) {
        state.selectedLabels = state.selectedLabels.filter(labelId => labelId !== id);
      } else {
        state.selectedLabels.push(id);
      }
    },
    
    selectAllLabels: (state) => {
      state.selectedLabels = state.searchResults.map(label => label.id);
    },
    
    clearSelection: (state) => {
      state.selectedLabels = [];
    },
    
    // ラベルエディタ
    startLabelCreation: (state, action: PayloadAction<{ text: string; imageUrl?: string }>) => {
      state.labelEditor = {
        mode: 'create',
        originalText: action.payload.text,
        imagePreview: action.payload.imageUrl || null,
        annotations: [],
        concepts: [],
        suggestedValues: [],
      };
    },
    
    startLabelEditing: (state, action: PayloadAction<QuantificationLabel>) => {
      state.currentLabel = action.payload;
      state.labelEditor = {
        mode: 'edit',
        originalText: action.payload.linguistic.originalText,
        imagePreview: action.payload.visual.imageUrl,
        annotations: [...action.payload.visual.annotations],
        concepts: [...action.payload.concept.relatedConcepts],
        suggestedValues: [],
      };
    },
    
    addAnnotation: (state, action: PayloadAction<any>) => {
      state.labelEditor.annotations.push(action.payload);
    },
    
    updateAnnotation: (state, action: PayloadAction<{ index: number; annotation: any }>) => {
      state.labelEditor.annotations[action.payload.index] = action.payload.annotation;
    },
    
    removeAnnotation: (state, action: PayloadAction<number>) => {
      state.labelEditor.annotations.splice(action.payload, 1);
    },
    
    addConcept: (state, action: PayloadAction<string>) => {
      if (!state.labelEditor.concepts.includes(action.payload)) {
        state.labelEditor.concepts.push(action.payload);
      }
    },
    
    removeConcept: (state, action: PayloadAction<string>) => {
      state.labelEditor.concepts = state.labelEditor.concepts.filter(
        concept => concept !== action.payload
      );
    },
    
    closeLabelEditor: (state) => {
      state.labelEditor = {
        mode: null,
        originalText: '',
        imagePreview: null,
        annotations: [],
        concepts: [],
        suggestedValues: [],
      };
      state.currentLabel = null;
    },
    
    // 検索クエリ更新
    updateSearchQuery: (state, action: PayloadAction<Partial<LabelSearchQuery>>) => {
      state.searchQuery = { ...state.searchQuery, ...action.payload };
    },
    
    clearSearchQuery: (state) => {
      state.searchQuery = {};
    },
    
    // バルク操作
    startBulkOperation: (state, action: PayloadAction<{ type: any; total: number }>) => {
      state.bulkOperation = {
        type: action.payload.type,
        progress: 0,
        total: action.payload.total,
      };
    },
    
    updateBulkProgress: (state, action: PayloadAction<number>) => {
      state.bulkOperation.progress = action.payload;
    },
    
    completeBulkOperation: (state) => {
      state.bulkOperation = { type: null, progress: 0, total: 0 };
      state.selectedLabels = [];
    },
    
    // エラー処理
    clearError: (state) => {
      state.error = null;
    },
  },
  
  extraReducers: (builder) => {
    builder
      // createLabel
      .addCase(createLabel.pending, (state) => {
        state.loading = true;
        state.error = null;
      })
      .addCase(createLabel.fulfilled, (state, action) => {
        state.loading = false;
        state.labels.push(action.payload);
        state.searchResults.push(action.payload);
        state.labelEditor.mode = null;
      })
      .addCase(createLabel.rejected, (state, action) => {
        state.loading = false;
        state.error = action.payload as string;
      })
      
      // updateLabel
      .addCase(updateLabel.fulfilled, (state, action) => {
        const index = state.labels.findIndex(label => label.id === action.payload.id);
        if (index !== -1) {
          state.labels[index] = action.payload;
        }
        
        const searchIndex = state.searchResults.findIndex(label => label.id === action.payload.id);
        if (searchIndex !== -1) {
          state.searchResults[searchIndex] = action.payload;
        }
        
        state.currentLabel = action.payload;
        state.labelEditor.mode = null;
      })
      
      // verifyLabel
      .addCase(verifyLabel.fulfilled, (state, action) => {
        const label = state.labels.find(l => l.id === action.payload.labelId);
        if (label) {
          label.evaluation = action.payload.evaluation;
          label.metadata.validated = true;
        }
      })
      
      // searchLabels
      .addCase(searchLabels.fulfilled, (state, action) => {
        state.searchResults = action.payload.results;
        state.searchQuery = action.payload.query;
      })
      
      // loadStatistics
      .addCase(loadStatistics.fulfilled, (state, action) => {
        state.statistics = action.payload;
      })
      
      // suggestQuantification
      .addCase(suggestQuantification.fulfilled, (state, action) => {
        state.labelEditor.suggestedValues = action.payload.suggestions;
      });
  },
});

export const {
  setCurrentLabel,
  selectLabel,
  selectAllLabels,
  clearSelection,
  startLabelCreation,
  startLabelEditing,
  addAnnotation,
  updateAnnotation,
  removeAnnotation,
  addConcept,
  removeConcept,
  closeLabelEditor,
  updateSearchQuery,
  clearSearchQuery,
  startBulkOperation,
  updateBulkProgress,
  completeBulkOperation,
  clearError,
} = labelSlice.actions;

export default labelSlice.reducer;