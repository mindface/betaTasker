package model

import "time"

type Assessment struct {
	ID                  int       `gorm:"primaryKey" json:"id"`
	TaskID              int       `json:"task_id"`
	UserID              int       `json:"user_id"`
	// 有効性スコア
	EffectivenessScore  int       `json:"effectiveness_score"`
	// 努力スコア
	EffortScore         int       `json:"effort_score"`
	// 印象値影響スコア
	ImpactScore         int       `json:"impact_score"`
	// 定性的フィードバック
	QualitativeFeedback string    `json:"qualitative_feedback"`
	CreatedAt           time.Time `json:"created_at"`
	UpdatedAt           time.Time `json:"updated_at"`
}

func (Assessment) TableName() string {
  return "assessments"
}
