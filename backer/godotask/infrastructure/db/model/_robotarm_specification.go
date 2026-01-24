package model

import (
	"time"
	"gorm.io/gorm"
)

// RobotArmSpecification - ロボットアーム仕様
type RobotArmSpecification struct {
	ID                  string    `gorm:"type:varchar(255);primaryKey" json:"id"`
	ModelName           string    `json:"model_name" gorm:"type:varchar(255)"`
	DOF                 int       `json:"dof"` // 自由度
	Reach               float64   `json:"reach"` // リーチ（mm）
	Payload             float64   `json:"payload"` // 可搬重量（kg）
	RepeatAccuracy      float64   `json:"repeat_accuracy"` // 繰り返し精度（mm）
	MaxSpeed            float64   `json:"max_speed"` // 最大速度（mm/s）
	WorkEnvelope        JSON      `json:"work_envelope" gorm:"type:jsonb"`
	JointLimits         JSON      `json:"joint_limits" gorm:"type:jsonb"`
	TeachingMethod      string    `json:"teaching_method" gorm:"type:varchar(100)"` // manual, vision, force, ai
	ControlSystem       JSON      `json:"control_system" gorm:"type:jsonb"`
	SafetyFeatures      JSON      `json:"safety_features" gorm:"type:jsonb"`
	MaintenanceSchedule JSON      `json:"maintenance_schedule" gorm:"type:jsonb"`
	CreatedAt           time.Time `json:"created_at"`
	UpdatedAt           time.Time `json:"updated_at"`
	DeletedAt           gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}
