package model

import "time"

type MemoryContext struct {
	ID          int       `gorm:"primaryKey" json:"id"`
	UserID      int       `json:"user_id"`
	TaskID      int       `json:"task_id"`
	Level       int       `json:"level"`
	WorkTarget  string    `json:"work_target"`
	Machine     string    `json:"machine"`
	MaterialSpec string   `json:"material_spec"`
	ChangeFactor string   `json:"change_factor"`
	Goal        string    `json:"goal"`
	CreatedAt   time.Time `json:"created_at"`
	// 補助情報
	TechnicalFactors      []TechnicalFactor         `gorm:"foreignKey:ContextID" json:"technical_factors"`
	KnowledgeTransformations []KnowledgeTransformation `gorm:"foreignKey:ContextID" json:"knowledge_transformations"`
}

func (MemoryContext) TableName() string {
	return "memory_contexts"
}