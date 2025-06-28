
export interface Memory {
  id: number;
  user_id: number;
  source_type: string;
  title: string;
  author: string;
  notes: string;
  factor: string;
  process: string;
  evaluation_axis: string;
  information_amount: string;
  tags: string;
  read_status: string;
  read_date: string | null;
  created_at: string;
  updated_at: string;
}

export interface AddMemory {
  title: string;
  user_id: number;
  author: string;
  source_type: string;
  notes: string;
  factor: string;
  process: string;
  evaluation_axis: string;
  information_amount: string;
  tags: string;
  read_status: string;
  read_date: string;
}
