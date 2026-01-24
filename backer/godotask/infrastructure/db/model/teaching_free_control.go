package model

import (
	"time"
	"gorm.io/gorm"
)

// TeachingFreeControl - ティーチング不要制御
type TeachingFreeControl struct {
	ID             string         `gorm:"type:varchar(255);primaryKey" json:"id"`
	TaskID         int            `json:"task_id" gorm:"index"`
	RobotID        string         `json:"robot_id" gorm:"index"`
	TaskType       string         `json:"task_type"`
	VisionSystem   JSON           `json:"vision_system" gorm:"type:jsonb"`
	ForceControl   JSON           `json:"force_control" gorm:"type:jsonb"`
	AIModel        JSON           `json:"ai_model" gorm:"type:jsonb"`
	LearningData   JSON           `json:"learning_data" gorm:"type:jsonb"`
	SuccessRate    float64        `json:"success_rate"`
	AdaptationTime float64        `json:"adaptation_time"`
	ErrorRecovery  JSON           `json:"error_recovery" gorm:"type:jsonb"`
	PerformanceLog JSON           `json:"performance_log" gorm:"type:jsonb"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	DeletedAt      gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`

	Task Task `json:"task" gorm:"foreignKey:TaskID"`
}
