package model

import (
	"time"
)

type Task struct {
	ID          int        `gorm:"primaryKey" json:"id"`
	UserID      int        `json:"user_id"`
	MemoryID    int       `json:"memory_id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Date        *time.Time `json:"date"`
	Status      string     `json:"status"`
	Priority    int        `json:"priority"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`

	User       User        `json:"user" gorm:"foreignKey:UserID"`
	QualitativeLabels []QualitativeLabel `json:"qualitative_labels" gorm:"foreignKey:TaskID"`
	QuantificationLabels []QuantificationLabel `json:"quantification_labels" gorm:"foreignKey:TaskID"`
	MultimodalData   []MultimodalData   `json:"multimodal_data" gorm:"foreignKey:TaskID"`
	HeuristicsAnalysis []HeuristicsAnalysis `json:"heuristics_analysis" gorm:"foreignKey:TaskID"`
	HeuristicsTracking []HeuristicsTracking `json:"heuristics_trackings" gorm:"foreignKey:TaskID"`
	HeuristicsModel []HeuristicsModel `json:"heuristics_models" gorm:"foreignKey:TaskID"`
	HeuristicsInsight []HeuristicsInsight `json:"heuristics_insights" gorm:"foreignKey:TaskID"`
	HeuristicsPattern []HeuristicsPattern `json:"heuristics_patterns" gorm:"foreignKey:TaskID"`
  KnowledgePatterns    []KnowledgePattern    `json:"knowledge_patterns" gorm:"foreignKey:TaskID"`
  LanguageOptimization   []LanguageOptimization    `json:"language_optimizations" gorm:"foreignKey:TaskID"`
}

func (Task) TableName() string {
    return "tasks"
}
