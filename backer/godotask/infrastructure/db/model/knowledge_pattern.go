package model

import (
	"time"
	"gorm.io/gorm"
)

// KnowledgePattern - 知識パターン（暗黙知→形式知）
type KnowledgePattern struct {
	ID              string    `gorm:"type:varchar(255);primaryKey" json:"id"`
  TaskID          uint      `json:"task_id" gorm:"index"`
	Type            string    `json:"type" gorm:"type:varchar(50)"` // tacit, explicit, hybrid
	Domain          string    `json:"domain" gorm:"type:varchar(100)"`
	TacitKnowledge  string    `json:"tacit_knowledge" gorm:"type:text"`
	ExplicitForm    string    `json:"explicit_form" gorm:"type:text"`
	ConversionPath  JSON      `json:"conversion_path" gorm:"type:jsonb"` // SECIモデルのパス
	Accuracy        float64   `json:"accuracy"`
	Coverage        float64   `json:"coverage"`
	Consistency     float64   `json:"consistency"`
	AbstractLevel   string    `json:"abstract_level" gorm:"type:varchar(50)"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}
