export interface ProcessOptimization {
  id: string;
  process_id: string;
  optimization_type: string; // "speed" | "accuracy" | "energy" | "cost" として列挙型化も可能
  initial_state: any; // JSONB → any
  optimized_state: any; // JSONB → any
  improvement: number; // 改善率（%）
  method: string;
  iterations: number;
  convergence_time: number; // 収束時間（秒）
  validated_by: string;
  validation_date: string; // ISO8601文字列 (time.Time → string)
  created_at: string;
  updated_at: string;
  deleted_at?: string | null; // nullable timestamp
}

export interface AddProcessOptimization {
  process_id: string;
  optimization_type: string; // "speed" | "accuracy" | "energy" | "cost" として列挙型化も可能
  initial_state: any; // JSONB → any
  optimized_state: any; // JSONB → any
  improvement: number; // 改善率（%）
  method: string;
  iterations: number;
  convergence_time: number; // 収束時間（秒）
  validated_by: string;
  validation_date: string; // ISO8601文字列 (time.Time → string)
}
