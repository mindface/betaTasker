package model

import (
	"time"
	"gorm.io/gorm"

	"github.com/godotask/lib"
)

type JSON = lib.JSON

// QuantificationLabel - 定量化ラベルのメインモデル
type QuantificationLabel struct {
	ID  string    `gorm:"type:varchar(255);primaryKey" json:"id"`

	UserID int `json:"user_id" gorm:"index"`
	TaskID int  `json:"task_id" gorm:"index"`

	// 言語情報
	OriginalText    string `json:"original_text" gorm:"type:text"`
	NormalizedText  string `json:"normalized_text" gorm:"type:text"`
	Category        string `json:"category" gorm:"index"`
	// Context         *string `json:"context" gorm:"type:text"`
	Domain          string `json:"domain" gorm:"index"`

	// 画像情報
	// ImageURL         *string `json:"image_url" gorm:"type:text"`
	// ThumbnailURL     *string `json:"thumbnail_url" gorm:"type:text"`
	// ImageDescription *string `json:"image_description" gorm:"type:text"`
	// ImageMetadata    *JSON   `json:"image_metadata" gorm:"type:jsonb"`

	// 定量化情報
	// Value       *float64 `json:"value"`
	// Unit        *string  `json:"unit" gorm:"index"`
	// MinRange    *float64 `json:"min_range"`
	// MaxRange    *float64 `json:"max_range"`
	// TypicalValue *float64 `json:"typical_value"`
	// Precision   *int     `json:"precision"`
	// Confidence  *float64 `json:"confidence"`

	// 概念情報
	// AbstractLevel     *string `json:"abstract_level" gorm:"index"` // concrete, semi-abstract, abstract
	// CulturalContext   *string `json:"cultural_context"`
	// TemporalContext   *string `json:"temporal_context"`
	// SpatialContext    *string `json:"spatial_context"`
	// RelatedConcepts   *JSON   `json:"related_concepts" gorm:"type:jsonb"`
	// SemanticTags      *JSON   `json:"semantic_tags" gorm:"type:jsonb"`

	// 評価情報
	// Accuracy       *float64 `json:"accuracy"`
	// Consistency    *float64 `json:"consistency"`
	// Reproducibility *float64 `json:"reproducibility"`
	// Usability      *float64 `json:"usability"`
	// VerificationCount *int  `json:"verification_count"`
	// LastVerified   *time.Time `json:"last_verified"`

	// メタデータ
	// Source           *string `json:"source" gorm:"index"` // manual, automatic, hybrid
	// Validated        *bool   `json:"validated" gorm:"index;default:false"`
	// PublicVisibility *bool   `json:"public_visibility" gorm:"default:true"`
	// Tags             *JSON   `json:"tags" gorm:"type:jsonb"`
	// Notes            *string `json:"notes" gorm:"type:text"`

	// 履歴管理
	// Version   *int    `json:"version" gorm:"default:1"`
	// CreatedBy *string `json:"created_by"`
	// UpdatedBy *string `json:"updated_by"`

	// タイムスタンプ
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	// DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	// リレーション
	// Annotations []ImageAnnotation `json:"annotations" gorm:"foreignKey:LabelID"`
	// Revisions   []LabelRevision   `json:"revisions" gorm:"foreignKey:LabelID"`
}

// ImageAnnotation - 画像アノテーション
type ImageAnnotation struct {
	ID       string  `gorm:"type:varchar(255);primaryKey" json:"id"`
	LabelID  string  `json:"label_id" gorm:"index"`
	Type     string  `json:"type"` // region, point, measurement, text
	X        float64 `json:"x"`
	Y        float64 `json:"y"`
	Width    float64 `json:"width"`
	Height   float64 `json:"height"`
	Label    string  `json:"label"`
	Value    float64 `json:"value"`
	Unit     string  `json:"unit"`
	Confidence float64 `json:"confidence"`
	CreatedBy  string  `json:"created_by"`
	CreatedAt  time.Time `json:"created_at"`
}

// LabelRevision - ラベル改訂履歴
type LabelRevision struct {
	ID        string    `gorm:"type:varchar(255);primaryKey" json:"id"`
	LabelID   string    `json:"label_id" gorm:"index"`
	Version   int       `json:"version"`
	Changes   JSON      `json:"changes" gorm:"type:jsonb"`
	Comment   string    `json:"comment" gorm:"type:text"`
	UserID    string    `json:"user_id"`
	Timestamp time.Time `json:"timestamp"`
}

// LabelDataset - ラベルデータセット
type LabelDataset struct {
	ID          string `gorm:"type:varchar(255);primaryKey" json:"id"`
	Name        string `json:"name"`
	Description string `json:"description" gorm:"type:text"`
	Domain      string `json:"domain" gorm:"index"`
	
	// 統計情報
	TotalLabels    int     `json:"total_labels"`
	VerifiedLabels int     `json:"verified_labels"`
	AverageAccuracy float64 `json:"average_accuracy"`
	
	// 品質メトリクス
	Completeness float64 `json:"completeness"`
	Consistency  float64 `json:"consistency"`
	Diversity    float64 `json:"diversity"`
	Balance      float64 `json:"balance"`

	// メタデータ
	Version   string `json:"version"`
	License   string `json:"license"`
	Citation  string `json:"citation"`
	CreatedBy string `json:"created_by"`

	// タイムスタンプ
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	// リレーション
	Labels []QuantificationLabel `json:"labels" gorm:"many2many:dataset_labels;"`
}

// LabelRelation - ラベル間の関係
type LabelRelation struct {
	ID           string  `gorm:"type:varchar(255);primaryKey" json:"id"`
	SourceID     string  `json:"source_id" gorm:"index"`
	TargetID     string  `json:"target_id" gorm:"index"`
	RelationType string  `json:"relation_type"` // synonym, hypernym, hyponym, etc.
	Strength     float64 `json:"strength"`
	Bidirectional bool   `json:"bidirectional"`
	Context      string  `json:"context"`
	Confidence   float64 `json:"confidence"`
	CreatedAt    time.Time `json:"created_at"`
}

// VisualMetaphor - 視覚的メタファー
type VisualMetaphor struct {
	ID              string  `gorm:"type:varchar(255);primaryKey" json:"id"`
	Metaphor        string  `json:"metaphor"`
	ReferenceObject string  `json:"reference_object"`
	Width           float64 `json:"width"`
	Height          float64 `json:"height"`
	Depth           float64 `json:"depth"`
	ImageURL        string  `json:"image_url"`
	MinVariability  float64 `json:"min_variability"`
	MaxVariability  float64 `json:"max_variability"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

// UserCalibration - ユーザーキャリブレーション
type UserCalibration struct {
	ID              string    `gorm:"type:varchar(255);primaryKey" json:"id"`
	UserID          uint      `json:"user_id" gorm:"index"`
	ReferenceObject string    `json:"reference_object"`
	Measurements    JSON      `json:"measurements" gorm:"type:jsonb"`
	ImageURL        string    `json:"image_url"`
	Confidence      float64   `json:"confidence"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

// MultimodalData - マルチモーダル処理データ
type MultimodalData struct {
	ID       string `gorm:"type:varchar(255);primaryKey" json:"id"`
	UserID   uint   `json:"user_id" gorm:"index"`
	TaskID   uint   `json:"task_id" gorm:"index"`

	// 言語特徴
	Text            string  `json:"text"`
	Tokens          JSON    `json:"tokens" gorm:"type:jsonb"`
	SemanticVector  JSON    `json:"semantic_vector" gorm:"type:jsonb"`
	AmbiguityScore  float64 `json:"ambiguity_score"`

	// 画像特徴
	ImageURL    string  `json:"image_url"`
	Objects     JSON    `json:"objects" gorm:"type:jsonb"`
	Measurements JSON   `json:"measurements" gorm:"type:jsonb"`
	ImageConfidence float64 `json:"image_confidence"`
	
	// 関連付け
	MappingType        string  `json:"mapping_type"` // direct, inferred, learned
	CorrelationScore   float64 `json:"correlation_score"`
	ContextRelevance   float64 `json:"context_relevance"`
	HistoricalAccuracy float64 `json:"historical_accuracy"`
	
	// 定量化結果
	Value      float64 `json:"value"`
	Unit       string  `json:"unit"`
	MinRange   float64 `json:"min_range"`
	MaxRange   float64 `json:"max_range"`
	Confidence float64 `json:"confidence"`
	
	// メタデータ
	Verified     bool   `json:"verified"`
	UserFeedback string `json:"user_feedback"` // correct, too_high, too_low, incorrect
	
	// タイムスタンプ
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// リクエスト/レスポンス構造体

// CreateLabelRequest - ラベル作成リクエスト
type CreateLabelRequest struct {
	Text        string   `json:"text" binding:"required"`
	Description string   `json:"description" binding:"required"`
	Value       float64  `json:"value" binding:"required"`
	Unit        string   `json:"unit" binding:"required"`
	Domain      string   `json:"domain" binding:"required"`
	Category    string   `json:"category" binding:"required"`
	ImageURL    string   `json:"image_url"`
	Concepts    []string `json:"concepts"`
	Tags        []string `json:"tags"`
}

// UpdateLabelRequest - ラベル更新リクエスト
type UpdateLabelRequest struct {
	ID      string                 `json:"id" binding:"required"`
	Updates map[string]interface{} `json:"updates" binding:"required"`
	Reason  string                 `json:"reason" binding:"required"`
}

// VerifyLabelRequest - ラベル検証リクエスト
type VerifyLabelRequest struct {
	LabelID      string                 `json:"label_id" binding:"required"`
	Verification map[string]interface{} `json:"verification" binding:"required"`
	VerifierID   string                 `json:"verifier_id" binding:"required"`
}

// LabelSearchQuery - ラベル検索クエリ
type LabelSearchQuery struct {
	Text         string   `json:"text"`
	Domain       string   `json:"domain"`
	Category     string   `json:"category"`
	MinValue     float64  `json:"min_value"`
	MaxValue     float64  `json:"max_value"`
	Unit         string   `json:"unit"`
	Concepts     []string `json:"concepts"`
	MinConfidence float64 `json:"min_confidence"`
	Verified     *bool    `json:"verified"`
	From         string   `json:"from"`
	To           string   `json:"to"`
	Limit        int      `json:"limit"`
	Offset       int      `json:"offset"`
	SortBy       string   `json:"sort_by"`
	SortOrder    string   `json:"sort_order"`
}

// LabelStatistics - ラベル統計
type LabelStatistics struct {
	TotalLabels         int                    `json:"total_labels"`
	LabelsByDomain      map[string]int         `json:"labels_by_domain"`
	LabelsByCategory    map[string]int         `json:"labels_by_category"`
	LabelsByConcept     map[string]int         `json:"labels_by_concept"`
	AverageMetrics      map[string]float64     `json:"average_metrics"`
	ValueDistribution   []map[string]interface{} `json:"value_distribution"`
	UnitDistribution    map[string]int         `json:"unit_distribution"`
	ConceptDistribution map[string]int         `json:"concept_distribution"`
	Temporal            map[string]int         `json:"temporal"`
	Quality             map[string]int         `json:"quality"`
}

// SuggestionRequest - 定量化提案リクエスト
type SuggestionRequest struct {
	Text     string `json:"text" binding:"required"`
	ImageURL string `json:"image_url"`
	Domain   string `json:"domain"`
}

// SuggestionResponse - 定量化提案レスポンス
type SuggestionResponse struct {
	Suggestions []struct {
		Value      float64 `json:"value"`
		Unit       string  `json:"unit"`
		Confidence float64 `json:"confidence"`
		Source     string  `json:"source"`
	} `json:"suggestions"`
}
