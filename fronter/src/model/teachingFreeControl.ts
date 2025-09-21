import { Task } from './task';

export interface TeachingFreeControl {
  id: string;
  task_id: number;
  robot_id: string;
  task_type: string;
  vision_system: any;       // JSONB → any
  force_control: any;       // JSONB → any
  ai_model: any;            // JSONB → any
  learning_data: any;       // JSONB → any
  success_rate: number;
  adaptation_time: number;
  error_recovery: any;      // JSONB → any
  performance_log: any;     // JSONB → any
  created_at: string;       // ISO date string
  updated_at: string;       // ISO date string
  deleted_at?: string | null; // nullable timestamp

  task?: Task;
}

export interface AddTeachingFreeControl {
  task_id: number;
  robot_id: string;
  task_type: string;
  vision_system: any;       // JSONB → any
  force_control: any;       // JSONB → any
  ai_model: any;            // JSONB → any
  learning_data: any;       // JSONB → any
  success_rate: number;
  adaptation_time: number;
  error_recovery: any;      // JSONB → any
  performance_log: any;     // JSONB → any
}
