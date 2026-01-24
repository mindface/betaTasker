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

type TechnicalFactor struct {
	ID         int       `gorm:"primaryKey" json:"id"`
	ContextID  int       `json:"context_id"`
	ToolSpec   string    `json:"tool_spec"`
	EvalFactors string   `json:"eval_factors"`
	MeasurementMethod string   `gorm:"column:measurement_method" json:"measurement_method"`
	Concern    string    `json:"concern"`
	CreatedAt  time.Time `json:"created_at"`
}

type KnowledgeTransformation struct {
	ID         int       `gorm:"primaryKey" json:"id"`
	ContextID  int       `json:"context_id"`
	Transformation string `json:"transformation"`
	Countermeasure  string `json:"countermeasure"`
	ModelFeedback   string `json:"model_feedback"`
	LearnedKnowledge string `json:"learned_knowledge"`
	CreatedAt  time.Time `json:"created_at"`
}

func (MemoryContext) TableName() string {
	return "memory_contexts"
}
func (TechnicalFactor) TableName() string {
	return "technical_factors"
}
func (KnowledgeTransformation) TableName() string {
	return "knowledge_transformations"
}
