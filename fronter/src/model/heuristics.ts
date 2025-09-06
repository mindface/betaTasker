// Heuristics関連の型定義

export interface HeuristicsAnalysis {
  id: number;
  user_id: number;
  task_id: number;
  analysis_type: string;
  result: any; // JSONデータ
  score: number;
  status: string;
  created_at: string;
  updated_at: string;
}

export interface HeuristicsTracking {
  id: number;
  user_id: number;
  action: string;
  context: any; // JSONデータ
  session_id: string;
  timestamp: string;
  duration: number; // ミリ秒
  created_at: string;
  updated_at: string;
}

export interface HeuristicsInsight {
  id: number;
  user_id: number;
  type: string;
  title: string;
  description: string;
  confidence: number;
  data: any; // JSONデータ
  is_active: boolean;
  created_at: string;
  updated_at: string;
}

export interface HeuristicsPattern {
  id: number;
  name: string;
  category: string;
  pattern: any; // JSONデータ
  frequency: number;
  accuracy: number;
  last_seen: string;
  created_at: string;
  updated_at: string;
}

export interface HeuristicsModel {
  id: number;
  model_type: string;
  version: string;
  parameters: any; // JSONデータ
  performance: any; // JSONデータ
  status: 'training' | 'ready' | 'deprecated';
  trained_at: string;
  created_at: string;
  updated_at: string;
}

// リクエスト用の型
export interface AddHeuristicsAnalysis {
  user_id: number;
  task_id: number;
  analysis_type: string;
  result: any;
  score: number;
  status: string;
}

export interface AddHeuristicsTracking {
  user_id: number;
  action: string;
  context: any;
  session_id: string;
  timestamp: string;
  duration: number;
}

export interface AddHeuristicsInsight {
  user_id: number;
  type: string;
  title: string;
  description: string;
  confidence: number;
  data: any;
  is_active: boolean;
}

export interface AddHeuristicsPattern {
  name: string;
  category: string;
  pattern: any;
  frequency: number;
  accuracy: number;
  last_seen: string;
}

export interface HeuristicsAnalysisRequest {
  user_id?: number;
  task_id?: number;
  analysis_type?: string;
  data?: Record<string, any>;
}

export interface HeuristicsTrackingData {
  user_id: number;
  action: string;
  context?: Record<string, any>;
  session_id?: string;
  duration?: number;
}

export interface HeuristicsTrainRequest {
  model_type: string;
  parameters?: Record<string, any>;
  data_source?: string;
  training_data?: any[];
}