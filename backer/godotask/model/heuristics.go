package model

import (
	"time"
	"gorm.io/gorm"
)

// HeuristicsAnalysis - 分析結果を保存
type HeuristicsAnalysis struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	UserID      uint           `json:"user_id"`
	TaskID      uint           `json:"task_id"`
	AnalysisType string        `json:"analysis_type"`
	Result      string         `json:"result" gorm:"type:jsonb"`
	Score       float64        `json:"score"`
	Status      string         `json:"status"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}

// HeuristicsTracking - ユーザー行動追跡データ
type HeuristicsTracking struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	UserID      uint           `json:"user_id"`
	Action      string         `json:"action"`
	Context     string         `json:"context" gorm:"type:jsonb"`
	SessionID   string         `json:"session_id"`
	Timestamp   time.Time      `json:"timestamp"`
	Duration    int            `json:"duration"` // ミリ秒単位
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}

// HeuristicsInsight - 生成されたインサイト
type HeuristicsInsight struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	UserID      uint           `json:"user_id"`
	Type        string         `json:"type"`
	Title       string         `json:"title"`
	Description string         `json:"description"`
	Confidence  float64        `json:"confidence"`
	Data        string         `json:"data" gorm:"type:jsonb"`
	IsActive    bool           `json:"is_active" gorm:"default:true"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}

// HeuristicsPattern - 検出されたパターン
type HeuristicsPattern struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	Name        string         `json:"name"`
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
type HeuristicsModel struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
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
	UserID       uint                   `json:"user_id"`
	TaskID       uint                   `json:"task_id"`
	AnalysisType string                 `json:"analysis_type"`
	Data         map[string]interface{} `json:"data"`
}

type HeuristicsTrackingData struct {
	ID        uint                   `json:"id"`
	UserID    uint                   `json:"user_id"`
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