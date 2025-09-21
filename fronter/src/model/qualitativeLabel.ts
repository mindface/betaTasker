
export interface QualitativeLabel {
  id: string;
  task_id: number;
  user_id: number;
  Content: string;
  Category: string;
  created_at: string; // ISO date string
  updated_at: string; // ISO date string
}

export interface AddQualitativeLabel {
  task_id: number;
  user_id: number;
  Content: string;
  Category: string;
}
