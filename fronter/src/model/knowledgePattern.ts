
export interface ConversionPath {
  characteristics: string[];
  coefficient: string;
  level: number;
  outcomes: string[];
  pattern_type: string;
  task_type: string;
  triggers: string[];
}

export interface KnowledgePattern {
  id: string;
  TaskId: number;
  type: "tacit" | "explicit" | "hybrid";
  domain: string;
  tacit_knowledge: string;
  explicit_form: string;
  conversion_path: ConversionPath;
  accuracy: number;
  coverage: number;
  consistency: number;
  abstract_level: string;
  created_at: string; // time.Time → ISO8601 string
  updated_at: string; // time.Time → ISO8601 string
  deleted_at?: string | null; // gorm.DeletedAt → nullable timestamp
}

export interface AddKnowledgePattern {
  type: "tacit" | "explicit" | "hybrid";
  domain: string;
  tacit_knowledge: string;
  explicit_form: string;
  conversion_path: string; // JSONB → any （SECIモデルのパス）
  accuracy: number;
  coverage: number;
  consistency: number;
  abstract_level: string;
}
