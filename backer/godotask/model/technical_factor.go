package model

import "time"

type TechnicalFactor struct {
	ID         int       `gorm:"primaryKey" json:"id"`
	ContextID  int       `json:"context_id"`
	ToolSpec   string    `json:"tool_spec"`
	EvalFactors string   `json:"eval_factors"`
	MeasurementMethod string   `gorm:"column:measurement_method" json:"measurement_method"`
	Concern    string    `json:"concern"`
	CreatedAt  time.Time `json:"created_at"`
}
func (TechnicalFactor) TableName() string {
	return "technical_factors"
}
