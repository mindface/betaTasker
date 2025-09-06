package model

import "time"

// NullString - Nullable string type
type NullString struct {
	String string
	Valid  bool
}

// NullFloat64 - Nullable float64 type
type NullFloat64 struct {
	Float64 float64
	Valid   bool
}

// RobotSpecification - ロボット仕様
type RobotSpecification struct {
	ID                      string       `gorm:"type:varchar(255);primaryKey" json:"id"`
	ModelName              string       `gorm:"type:varchar(255);not null" json:"model_name"`
	DOF                    int          `gorm:"not null" json:"dof"`                     // Degrees of Freedom
	ReachMm                float64      `gorm:"not null" json:"reach_mm"`                // Reach in millimeters
	PayloadKg              float64      `gorm:"not null" json:"payload_kg"`              // Payload in kilograms
	RepeatAccuracyMm       float64      `gorm:"not null" json:"repeat_accuracy_mm"`      // Repeat accuracy in mm
	MaxSpeedMmS            float64      `gorm:"not null" json:"max_speed_mm_s"`          // Max speed in mm/s
	WorkEnvelopeShape      string       `gorm:"type:varchar(100)" json:"work_envelope_shape"`
	TeachingMethod         string       `gorm:"type:varchar(100)" json:"teaching_method"`
	ControlType            string       `gorm:"type:varchar(100)" json:"control_type"`
	VisionSystem           *NullString  `gorm:"type:varchar(255)" json:"vision_system"`
	ForceSensor            *NullString  `gorm:"type:varchar(255)" json:"force_sensor"`
	AICapability           *NullString  `gorm:"type:text" json:"ai_capability"`
	SafetyFeatures         *NullString  `gorm:"type:text" json:"safety_features"`
	MaintenanceIntervalHours int        `gorm:"default:1000" json:"maintenance_interval_hours"`
	CreatedAt              time.Time    `json:"created_at"`
	UpdatedAt              time.Time    `json:"updated_at"`
}

// OptimizationModel - 最適化モデル
type OptimizationModel struct {
	ID               string        `gorm:"type:varchar(255);primaryKey" json:"id"`
	Name             string        `gorm:"type:varchar(255);not null" json:"name"`
	Type             string        `gorm:"type:varchar(100);not null" json:"type"`
	ObjectiveFunction string       `gorm:"type:text" json:"objective_function"`
	Constraints      string        `gorm:"type:text" json:"constraints"`
	Parameters       *NullString   `gorm:"type:text" json:"parameters"`
	PerformanceMetric *NullString  `gorm:"type:text" json:"performance_metric"`
	IterationCount   *NullFloat64  `gorm:"type:decimal(10,2)" json:"iteration_count"`
	ConvergenceRate  *NullFloat64  `gorm:"type:decimal(5,4)" json:"convergence_rate"`
	Domain           string        `gorm:"type:varchar(100)" json:"domain"`
	Application      string        `gorm:"type:text" json:"application"`
	CreatedAt        time.Time     `json:"created_at"`
	UpdatedAt        time.Time     `json:"updated_at"`
}