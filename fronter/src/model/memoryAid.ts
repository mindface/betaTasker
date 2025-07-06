// /src/model/memoryAid.ts
export interface TechnicalFactor {
  id: number;
  context_id: number;
  tool_spec: string;
  eval_factors: string;
  measurement: string;
  concern: string;
  created_at: string;
}

export interface KnowledgeTransformation {
  id: number;
  context_id: number;
  transformation: string;
  countermeasure: string;
  model_feedback: string;
  learned_knowledge: string;
  created_at: string;
}

export interface MemoryContext {
  id: number;
  user_id: number;
  task_id: number;
  level: number;
  work_target: string;
  machine: string;
  material_spec: string;
  change_factor: string;
  goal: string;
  created_at: string;
  technical_factors: TechnicalFactor[];
  knowledge_transformations: KnowledgeTransformation[];
}
