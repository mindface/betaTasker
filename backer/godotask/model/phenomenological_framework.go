package model

import (
	"time"
	"gorm.io/gorm"
)

// PhenomenologicalFramework - 現象学的フレームワーク
type PhenomenologicalFramework struct {
	ID            string         `gorm:"type:varchar(255);primaryKey" json:"id"`
	TaskID        int            `json:"task_id" gorm:"index"`
	Name          string         `json:"name" gorm:"type:varchar(255)"`
	Description   string         `json:"description" gorm:"type:text"`
	Goal          string         `json:"goal" gorm:"type:text"`  // G（Goal）
	Scope         string         `json:"scope" gorm:"type:text"` // A（範囲）
	Process       JSON           `json:"process" gorm:"type:jsonb"`
	Result        JSON           `json:"result" gorm:"type:jsonb"`
	Feedback      JSON           `json:"feedback" gorm:"type:jsonb"`
	LimitMin      float64        `json:"limit_min"`
	LimitMax      float64        `json:"limit_max"`
	GoalFunction  string         `json:"goal_function" gorm:"type:text"`
	AbstractLevel string         `json:"abstract_level" gorm:"type:varchar(50)"`
	Domain        string         `json:"domain" gorm:"type:varchar(100)"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`

	Task Task `json:"task" gorm:"foreignKey:TaskID"`
}
