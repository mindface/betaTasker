import { Task } from './task';

export interface LanguageOptimization {
  id: string;
  task_id: number;
  original_text: string;
  optimized_text: string;
  domain: string;
  abstraction_level: string;
  precision: number;
  clarity: number;
  completeness: number;
  context: string;          // JSONB → any
  transformation: string;   // JSONB → any
  evaluation_score: number;
  created_at: string;    // time.Time → ISO8601 string
  updated_at: string;    // time.Time → ISO8601 string
  deleted_at?: string | null; // gorm.DeletedAt → nullable timestamp

  task?: Task; // リレーション（必要なら Task インターフェースを別途定義）
}

export interface AddLanguageOptimization {
  task_id: number;
  original_text: string;
  optimized_text: string;
  domain: string;
  abstraction_level: string;
  precision: number;
  clarity: number;
  completeness: number;
  context: string;          // JSONB → any
  transformation: string;   // JSONB → any
  evaluation_score: number;
}
