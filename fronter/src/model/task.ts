import {
  HeuristicsModel,
  HeuristicsAnalysis,
  HeuristicsPattern,
  HeuristicsTracking,
  HeuristicsInsight
} from "./heuristics";
import {
  LanguageOptimization
} from "./languageOptimization";

export interface Task {
  id: number;
  user_id: number;
  memory_id?: number | null;
  title: string;
  description: string;
  date?: string | null; // ISO8601形式
  status: string; // todo, in_progress, completed
  priority: number;
  created_at: string;
  updated_at: string;

  // リレーション
  heuristics_model?: HeuristicsModel;
  heuristics_analysis?: HeuristicsAnalysis[];
  heuristics_patterns?: HeuristicsPattern[];
  heuristics_trackings?: HeuristicsTracking[];
  heuristics_insights?: HeuristicsInsight[];
  language_optimizations?: LanguageOptimization[];
}

export interface AddTask {
  user_id: number;
  memory_id?: number | null;
  title: string;
  description: string;
  date?: string | null;
  status: string;
  priority: number;
}
