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
