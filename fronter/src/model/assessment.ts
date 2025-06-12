export interface Assessment {
  id: number;
  task_id: number;
  user_id: number;
  effectiveness_score: number;
  effort_score: number;
  impact_score: number;
  qualitative_feedback: string;
  created_at: string;
  updated_at: string;
}

export interface AddAssessment {
  task_id: number;
  user_id: number;
  effectiveness_score: number;
  effort_score: number;
  impact_score: number;
  qualitative_feedback: string;
}
