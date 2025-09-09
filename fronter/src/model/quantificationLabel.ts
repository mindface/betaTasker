// 定量化ラベルのデータモデル

// ラベル付けされた定量化データ
export interface QuantificationLabel {
  id: string;

  // 言語情報
  linguistic: {
    originalText: string;              // 元のテキスト（例: "コップ半分"）
    normalizedText: string;            // 正規化テキスト
    category: string;                  // カテゴリ（量、サイズ、程度など）
    context: string;                   // 使用文脈
    domain: string;                    // ドメイン（料理、建築、デザインなど）
  };
  
  // 画像情報
  visual: {
    imageUrl: string;                  // 画像URL
    thumbnailUrl: string;              // サムネイルURL
    imageDescription: string;          // 画像の説明テキスト
    annotations: ImageAnnotation[];    // 画像アノテーション
    metadata: {
      width: number;
      height: number;
      format: string;
      capturedAt: string;
      device?: string;
    };
  };
  
  // 定量化情報
  quantification: {
    value: number;                     // 数値
    unit: string;                      // 単位
    range: {
      min: number;
      max: number;
      typical: number;                 // 典型値
      distribution: 'normal' | 'uniform' | 'skewed';
    };
    precision: number;                 // 精度（小数点以下の桁数）
    confidence: number;                // 信頼度
  };
  
  // 概念情報
  concept: {
    abstractLevel: 'concrete' | 'semi-abstract' | 'abstract';  // 抽象度
    culturalContext?: string;         // 文化的文脈
    temporalContext?: string;         // 時間的文脈
    spatialContext?: string;          // 空間的文脈
    relatedConcepts: string[];        // 関連概念
    semanticTags: string[];           // 意味タグ
  };
  
  // 評価情報
  evaluation: {
    accuracy: number;                  // 精度評価
    consistency: number;               // 一貫性評価
    reproducibility: number;           // 再現性評価
    usability: number;                 // 使いやすさ評価
    verificationCount: number;         // 検証回数
    lastVerified: string;             // 最終検証日時
  };
  
  // 履歴情報
  history: {
    createdAt: string;
    createdBy: string;
    updatedAt: string;
    updatedBy: string;
    version: number;
    revisions: LabelRevision[];
  };
  
  // メタデータ
  metadata: {
    source: 'manual' | 'automatic' | 'hybrid';  // ラベル付けソース
    validated: boolean;                          // 検証済みフラグ
    publicVisibility: boolean;                   // 公開フラグ
    tags: string[];                              // 汎用タグ
    notes?: string;                              // 備考
  };
}

// 画像アノテーション
export interface ImageAnnotation {
  id: string;
  type: 'region' | 'point' | 'measurement' | 'text';
  coordinates: {
    x: number;
    y: number;
    width?: number;
    height?: number;
  };
  label: string;
  value?: number;
  unit?: string;
  confidence: number;
  createdBy: string;
  createdAt: string;
}

// ラベル改訂履歴
export interface LabelRevision {
  revisionId: string;
  timestamp: string;
  userId: string;
  changes: {
    field: string;
    oldValue: any;
    newValue: any;
    reason?: string;
  }[];
  comment?: string;
}

// ラベルデータセット
export interface LabelDataset {
  id: string;
  name: string;
  description: string;
  domain: string;
  labels: QuantificationLabel[];
  
  statistics: {
    totalLabels: number;
    verifiedLabels: number;
    averageAccuracy: number;
    coverageByConcept: Record<string, number>;
    coverageByDomain: Record<string, number>;
  };
  
  quality: {
    completeness: number;              // データの完全性
    consistency: number;               // 一貫性
    diversity: number;                 // 多様性
    balance: number;                   // バランス
  };
  
  metadata: {
    createdAt: string;
    createdBy: string;
    lastUpdated: string;
    version: string;
    license?: string;
    citation?: string;
  };
}

// ラベル関係
export interface LabelRelation {
  id: string;
  sourceLabel: string;                // ソースラベルID
  targetLabel: string;                // ターゲットラベルID
  relationType: 'synonym' | 'hypernym' | 'hyponym' | 'meronym' | 'holonym' | 'similar' | 'opposite';
  strength: number;                    // 関係の強さ (0-1)
  bidirectional: boolean;             // 双方向関係かどうか
  context?: string;                   // 関係が成立する文脈
  confidence: number;
}

// ラベル検索クエリ
export interface LabelSearchQuery {
  text?: string;                      // テキスト検索
  domain?: string;                    // ドメインフィルタ
  category?: string;                  // カテゴリフィルタ
  valueRange?: {
    min: number;
    max: number;
    unit: string;
  };
  concepts?: string[];                // 概念フィルタ
  minConfidence?: number;             // 最小信頼度
  verified?: boolean;                 // 検証済みのみ
  dateRange?: {
    from: string;
    to: string;
  };
  limit?: number;
  offset?: number;
  sortBy?: 'relevance' | 'confidence' | 'date' | 'usage';
  sortOrder?: 'asc' | 'desc';
}

// ラベル作成リクエスト
export interface CreateLabelRequest {
  text: string;
  imageFile?: File;
  imageUrl?: string;
  description: string;
  value: number;
  unit: string;
  domain: string;
  category: string;
  concepts?: string[];
  tags?: string[];
}

// ラベル更新リクエスト
export interface UpdateLabelRequest {
  id: string;
  updates: Partial<{
    normalizedText: string;
    imageDescription: string;
    value: number;
    unit: string;
    confidence: number;
    concepts: string[];
    tags: string[];
    notes: string;
  }>;
  reason: string;                     // 更新理由
}

// ラベル検証リクエスト
export interface VerifyLabelRequest {
  labelId: string;
  verification: {
    accurate: boolean;
    consistency: boolean;
    reproducible: boolean;
    usable: boolean;
    feedback?: string;
    suggestedValue?: number;
    suggestedUnit?: string;
  };
  verifierId: string;
}

// ラベル統計
export interface LabelStatistics {
  totalLabels: number;
  labelsByDomain: Record<string, number>;
  labelsByCategory: Record<string, number>;
  labelsByConcept: Record<string, number>;
  
  averageMetrics: {
    confidence: number;
    accuracy: number;
    consistency: number;
    reproducibility: number;
  };
  
  distribution: {
    valueDistribution: Array<{
      range: string;
      count: number;
      percentage: number;
    }>;
    unitDistribution: Record<string, number>;
    conceptDistribution: Record<string, number>;
  };
  
  temporal: {
    createdLast24h: number;
    createdLast7d: number;
    createdLast30d: number;
    verifiedLast24h: number;
    verifiedLast7d: number;
    verifiedLast30d: number;
  };
  
  quality: {
    highConfidence: number;           // 信頼度80%以上
    mediumConfidence: number;         // 信頼度50-80%
    lowConfidence: number;            // 信頼度50%未満
    verified: number;
    unverified: number;
  };
}