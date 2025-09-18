
export interface KnowledgePattern {
  id: string;
  type: "tacit" | "explicit" | "hybrid";
  domain: string;
  tacit_knowledge: string;
  explicit_form: string;
  conversion_path: string;    // JSONB → any （SECIモデルのパス）
  accuracy: number;
  coverage: number;
  consistency: number;
  abstract_level: string;
  created_at: string;     // time.Time → ISO8601 string
  updated_at: string;     // time.Time → ISO8601 string
  deleted_at?: string | null;     // gorm.DeletedAt → nullable timestamp
}

export interface AddKnowledgePattern {
  type: "tacit" | "explicit" | "hybrid";
  domain: string;
  tacit_knowledge: string;
  explicit_form: string;
  conversion_path: string;   // JSONB → any （SECIモデルのパス）
  accuracy: number;
  coverage: number;
  consistency: number;
  abstract_level: string;
}
