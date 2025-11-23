package model

import "time"

type KnowledgeTransformation struct {
	ID         int       `gorm:"primaryKey" json:"id"`
	ContextID  int       `json:"context_id"`
	Transformation string `json:"transformation"`
	Countermeasure  string `json:"countermeasure"`
	ModelFeedback   string `json:"model_feedback"`
	LearnedKnowledge string `json:"learned_knowledge"`
	CreatedAt  time.Time `json:"created_at"`
}

func (KnowledgeTransformation) TableName() string {
	return "knowledge_transformations"
}
