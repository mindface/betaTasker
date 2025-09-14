package model

import (
	"time"
	"gorm.io/gorm"
)

// LanguageOptimization - 言語最適化モデル
type LanguageOptimization struct {
	ID               string         `gorm:"type:varchar(255);primaryKey" json:"id"`
	TaskID           int            `json:"task_id" gorm:"index"`
	OriginalText     string         `json:"original_text" gorm:"type:text"`
	OptimizedText    string         `json:"optimized_text" gorm:"type:text"`
	Domain           string         `json:"domain" gorm:"type:varchar(100)"`
	AbstractionLevel string         `json:"abstraction_level"`
	Precision        float64        `json:"precision"`
	Clarity          float64        `json:"clarity"`
	Completeness     float64        `json:"completeness"`
	Context          JSON           `json:"context" gorm:"type:jsonb"`
	Transformation   JSON           `json:"transformation" gorm:"type:jsonb"`
	EvaluationScore  float64        `json:"evaluation_score"`
	CreatedAt        time.Time      `json:"created_at"`
	UpdatedAt        time.Time      `json:"updated_at"`
	DeletedAt        gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`

	Task Task `json:"task" gorm:"foreignKey:TaskID"`
}
