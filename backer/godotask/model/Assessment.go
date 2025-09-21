package model

import "time"

type Assessment struct {
	ID                  int       `gorm:"primaryKey" json:"id"`
	TaskID              int       `json:"task_id"`
	UserID              int       `json:"user_id"`
	EffectivenessScore  int       `json:"effectiveness_score"`
	EffortScore         int       `json:"effort_score"`
	ImpactScore         int       `json:"impact_score"`
	QualitativeFeedback string    `json:"qualitative_feedback"`
	CreatedAt           time.Time `json:"created_at"`
	UpdatedAt           time.Time `json:"updated_at"`
}

func (Assessment) TableName() string {
  return "assessments"
}
