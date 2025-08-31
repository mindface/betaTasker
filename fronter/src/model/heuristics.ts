// Heuristics関連の型定義

// 分析タイプの定義
export type AnalysisType = 'performance' | 'behavior' | 'pattern' | 'cognitive' | 'efficiency';

// 分析ステータスの定義
export type AnalysisStatus = 'pending' | 'completed' | 'failed';

// モデルステータスの定義
export type ModelStatus = 'training' | 'ready' | 'deprecated';

// 分析結果の詳細型
export interface AnalysisResult {
  insights: string[];
  patterns: Pattern[];
  recommendations: Recommendation[];
  confidence: number;
  execution_time: number;
  data_points: number;
}

// パターンの詳細型
export interface Pattern {
  id: string;
  name: string;
  description: string;
  confidence: number;
  frequency: number;
  category: string;
}

// 推奨事項の型
export interface Recommendation {
  id: string;
  title: string;
  description: string;
  priority: 'low' | 'medium' | 'high';
  impact: string;
  effort: string;
}

// 分析メタデータ
export interface AnalysisMetadata {
  execution_time: number;
  data_points: number;
  algorithm_version: string;
  parameters: Record<string, any>;
  environment: string;
}

// 分析フィルター
export interface AnalysisFilters {
  user_id?: number;
  task_id?: number;
  analysis_type?: AnalysisType;
  status?: AnalysisStatus;
  date_range?: {
    start: string;
    end: string;
  };
}

// インサイトフィルター
export interface InsightFilters {
  user_id?: number;
  type?: string;
  confidence_min?: number;
  is_active?: boolean;
  date_range?: {
    start: string;
    end: string;
  };
}

// パターンフィルター
export interface PatternFilters {
  user_id?: number;
  category?: string;
  confidence_min?: number;
  frequency_min?: number;
}

// ページネーション状態
export interface PaginationState {
  page: number;
  limit: number;
  total: number;
  totalPages: number;
}

// バリデーション結果
export interface ValidationResult {
  isValid: boolean;
  errors: string[];
}

// 強化された分析インターフェース
export interface HeuristicsAnalysis {
  id: number;
  user_id: number;
  task_id?: number;
  analysis_type: AnalysisType;
  result: AnalysisResult;
  score: number;
  status: AnalysisStatus;
  metadata: AnalysisMetadata;
  created_at: string;
  updated_at: string;
}

// 強化されたトラッキングインターフェース
export interface HeuristicsTracking {
  id: number;
  user_id: number;
  action: string;
  context: Record<string, any>;
  session_id: string;
  timestamp: string;
  duration: number; // ミリ秒
  metadata: {
    user_agent: string;
    ip_address: string;
    device_type: string;
  };
  created_at: string;
  updated_at: string;
}

// 強化されたインサイトインターフェース
export interface HeuristicsInsight {
  id: number;
  user_id: number;
  type: string;
  title: string;
  description: string;
  confidence: number;
  data: Record<string, any>;
  is_active: boolean;
  tags: string[];
  priority: 'low' | 'medium' | 'high';
  created_at: string;
  updated_at: string;
}

// 強化されたパターンインターフェース
export interface HeuristicsPattern {
  id: number;
  name: string;
  category: string;
  pattern: Record<string, any>;
  frequency: number;
  accuracy: number;
  confidence: number;
  last_seen: string;
  metadata: {
    detection_method: string;
    training_data_size: number;
    validation_score: number;
  };
  created_at: string;
  updated_at: string;
}

// 強化されたモデルインターフェース
export interface HeuristicsModel {
  id: number;
  model_type: string;
  version: string;
  parameters: Record<string, any>;
  performance: {
    accuracy: number;
    precision: number;
    recall: number;
    f1_score: number;
  };
  status: ModelStatus;
  trained_at: string;
  metadata: {
    training_data_size: number;
    training_duration: number;
    algorithm: string;
    hyperparameters: Record<string, any>;
  };
  created_at: string;
  updated_at: string;
}

// リクエスト用の型
export interface HeuristicsAnalysisRequest {
  user_id: number;
  task_id?: number;
  analysis_type: AnalysisType;
  data?: Record<string, any>;
  parameters?: Record<string, any>;
}

export interface HeuristicsTrackingData {
  user_id: number;
  action: string;
  context?: Record<string, any>;
  session_id?: string;
  duration?: number;
  metadata?: {
    user_agent?: string;
    ip_address?: string;
    device_type?: string;
  };
}

export interface HeuristicsTrainRequest {
  model_type: string;
  parameters?: Record<string, any>;
  data_source?: string;
  training_data?: any[];
  hyperparameters?: Record<string, any>;
}

// 定数定義
export const VALID_ANALYSIS_TYPES: AnalysisType[] = [
  'performance',
  'behavior',
  'pattern',
  'cognitive',
  'efficiency'
];

export const VALID_MODEL_STATUSES: ModelStatus[] = [
  'training',
  'ready',
  'deprecated'
];

export const DEFAULT_PAGINATION: PaginationState = {
  page: 1,
  limit: 20,
  total: 0,
  totalPages: 0
};

// バリデーション関数
export const validateAnalysisRequest = (data: HeuristicsAnalysisRequest): ValidationResult => {
  const errors: string[] = [];
  
  if (!data.user_id || data.user_id <= 0) {
    errors.push('有効なユーザーIDが必要です');
  }
  
  if (!data.analysis_type || !VALID_ANALYSIS_TYPES.includes(data.analysis_type)) {
    errors.push('有効な分析タイプを指定してください');
  }
  
  if (data.parameters && typeof data.parameters !== 'object') {
    errors.push('パラメータはオブジェクト形式で指定してください');
  }
  
  return {
    isValid: errors.length === 0,
    errors
  };
};

export const validateTrackingData = (data: HeuristicsTrackingData): ValidationResult => {
  const errors: string[] = [];
  
  if (!data.user_id || data.user_id <= 0) {
    errors.push('有効なユーザーIDが必要です');
  }
  
  if (!data.action || data.action.trim().length === 0) {
    errors.push('アクション名が必要です');
  }
  
  if (data.action && data.action.length > 200) {
    errors.push('アクション名は200文字以内で入力してください');
  }
  
  if (data.duration && (data.duration < 0 || data.duration > 86400000)) {
    errors.push('継続時間は0〜24時間（86400000ミリ秒）の範囲で入力してください');
  }
  
  return {
    isValid: errors.length === 0,
    errors
  };
};

export const validateTrainRequest = (data: HeuristicsTrainRequest): ValidationResult => {
  const errors: string[] = [];
  
  if (!data.model_type || data.model_type.trim().length === 0) {
    errors.push('モデルタイプが必要です');
  }
  
  if (data.parameters && typeof data.parameters !== 'object') {
    errors.push('パラメータはオブジェクト形式で指定してください');
  }
  
  if (data.training_data && !Array.isArray(data.training_data)) {
    errors.push('トレーニングデータは配列形式で指定してください');
  }
  
  if (data.hyperparameters && typeof data.hyperparameters !== 'object') {
    errors.push('ハイパーパラメータはオブジェクト形式で指定してください');
  }
  
  return {
    isValid: errors.length === 0,
    errors
  };
};