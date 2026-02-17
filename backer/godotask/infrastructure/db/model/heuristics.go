package model

import (
	"time"
	"gorm.io/gorm"
)

// HeuristicsAnalysis - 分析結果を保存
type HeuristicsAnalysis struct {
	ID          int           `gorm:"primaryKey" json:"id"`
	UserID      int           `json:"user_id"`
	TaskID      int           `json:"task_id"`
	AnalysisType string        `json:"analysis_type"`
	Result      string         `json:"result" gorm:"type:jsonb"`

	// 追加したユーザー学習指標
	TimeSpentMinutes int         `json:"time_spent_minutes"` // タスクに使った時間
	DifficultyScore  float64     `json:"difficulty_score"`   // タスク難易度推定
	EfficiencyScore  float64     `json:"efficiency_score"`   // 効率スコア
	ErrorCount       int         `json:"error_count"`        // エラー回数
	Confidence       float64     `json:"confidence"`         // 分析信頼度 (0〜1)

	Score       float64        `json:"score"`
	Status      string         `json:"status"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`

	Patterns []HeuristicsPattern `gorm:"-"`
	Insights []HeuristicsInsight `gorm:"-"`
	Modelers []HeuristicsModeler `gorm:"-"`

}

// HeuristicsTracking - ユーザー行動追跡データ
type HeuristicsTracking struct {
	ID          int           `gorm:"primaryKey" json:"id"`
	UserID      int           `json:"user_id"`
	TaskID      int           `json:"task_id"`

	Action      string         `json:"action"`
	Context     string         `json:"context" gorm:"type:jsonb"`
	SessionID   string         `json:"session_id"`

	// 追加：集中指標
	FocusLevel    *float64    `json:"focus_level"` // 0.0〜1.0
	IsDistraction bool         `json:"is_distraction"`

	Timestamp   time.Time      `json:"timestamp"`
	Duration    int            `json:"duration"` // ミリ秒単位
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}

// HeuristicsInsight - 生成されたインサイト
type HeuristicsInsight struct {
	ID          int           `gorm:"primaryKey" json:"id"`
	UserID      int           `json:"user_id"`
	TaskID      int           `json:"task_id"`

	Type        string         `json:"type"`
	Title       string         `json:"title"`
	Description string         `json:"description"`
	Confidence  float64        `json:"confidence"`
	Data        string         `json:"data" gorm:"type:jsonb"`
	// 追加：どの分析から生成されたか
	SourceAnalysisID *int      `json:"source_analysis_id"`
	// 追加：AI からの具体的改善アクション
	Recommendation string        `json:"recommendation"`
	// 追加：改善した場合の期待値 (0〜1)
	ExpectedImpact float64       `json:"expected_impact"`
	IsActive    bool           `json:"is_active" gorm:"default:true"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}

// HeuristicsPattern - 検出されたパターン
type HeuristicsPattern struct {
	ID          int           `gorm:"primaryKey" json:"id"`
	Name        string         `json:"name"`
	UserID    int           `json:"user_id"`    // 個人ごとのパターン分析も可能に
	TaskID    int          `json:"task_id"`
	// タスク分類
	TaskType  string         `json:"task_type"`  // coding, planning, machining など
	// 追加：パターンの重要度（0〜1）
	ImpactScore float64      `json:"impact_score"`

	Category    string         `json:"category"`
	Pattern     string         `json:"pattern" gorm:"type:jsonb"`
	Frequency   int            `json:"frequency"`
	Accuracy    float64        `json:"accuracy"`
	LastSeen    time.Time      `json:"last_seen"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}

// HeuristicsModel - 学習モデル情報
type HeuristicsModeler struct {
	ID          int           `gorm:"primaryKey" json:"id"`
	UserID    	int           `json:"user_id"`
	TaskID    	int           `json:"task_id"`
	ModelType   string         `json:"model_type"`
	Version     string         `json:"version"`
	Parameters  string         `json:"parameters" gorm:"type:jsonb"`
	Performance string         `json:"performance" gorm:"type:jsonb"`
	Status      string         `json:"status"` // training, ready, deprecated
	TrainedAt   time.Time      `json:"trained_at"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}

// リクエスト/レスポンス用の構造体
type HeuristicsAnalysisRequest struct {
	UserID       int                   `json:"user_id"`
	TaskID       int                   `json:"task_id"`
	AnalysisType string                 `json:"analysis_type"`
	Data         map[string]interface{} `json:"data"`
}

type HeuristicsTrackingData struct {
	ID        int                   `json:"id"`
	UserID    int                   `json:"user_id"`
	Action    string                 `json:"action"`
	Context   map[string]interface{} `json:"context"`
	SessionID string                 `json:"session_id"`
	Duration  int                    `json:"duration"`
}

type HeuristicsTrainRequest struct {
	ModelType  string                 `json:"model_type"`
	Parameters map[string]interface{} `json:"parameters"`
	DataSource string                 `json:"data_source"`
}