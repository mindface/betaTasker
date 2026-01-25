package model

import (
	"time"
	"gorm.io/gorm"
)

// ProcessOptimization - プロセス最適化記録
type ProcessOptimization struct {
	ID              string    `gorm:"type:varchar(255);primaryKey" json:"id"`
	TaskID        int            `json:"task_id" gorm:"index"`
	ProcessID       string    `json:"process_id" gorm:"index"`
	OptimizationType string   `json:"optimization_type"` // speed, accuracy, energy, cost
	InitialState    JSON      `json:"initial_state" gorm:"type:jsonb"`
	OptimizedState  JSON      `json:"optimized_state" gorm:"type:jsonb"`
	Improvement     float64   `json:"improvement"` // 改善率（%）
	Method          string    `json:"method" gorm:"type:text"`
	Iterations      int       `json:"iterations"`
	ConvergenceTime float64   `json:"convergence_time"` // 収束時間（秒）
	ValidatedBy     string    `json:"validated_by"`
	ValidationDate  time.Time `json:"validation_date"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}
