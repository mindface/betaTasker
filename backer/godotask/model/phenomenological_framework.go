package model

import (
	"time"
	"gorm.io/gorm"
)

// PhenomenologicalFramework - 現象学的フレームワーク
type PhenomenologicalFramework struct {
	ID              string    `gorm:"type:varchar(255);primaryKey" json:"id"`
	Name            string    `json:"name" gorm:"type:varchar(255)"`
	Description     string    `json:"description" gorm:"type:text"`
	Goal            string    `json:"goal" gorm:"type:text"` // G（Goal）
	Scope           string    `json:"scope" gorm:"type:text"` // A（範囲）
	Process         JSON      `json:"process" gorm:"type:jsonb"` // Pa（プロセス）
	Result          JSON      `json:"result" gorm:"type:jsonb"`
	Feedback        JSON      `json:"feedback" gorm:"type:jsonb"`
	LimitMin        float64   `json:"limit_min"`
	LimitMax        float64   `json:"limit_max"`
	GoalFunction    string    `json:"goal_function" gorm:"type:text"` // goalFn()
	AbstractLevel   string    `json:"abstract_level" gorm:"type:varchar(50)"` // L0-L3
	Domain          string    `json:"domain" gorm:"type:varchar(100)"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

// KnowledgePattern - 知識パターン（暗黙知→形式知）
type KnowledgePattern struct {
	ID              string    `gorm:"type:varchar(255);primaryKey" json:"id"`
	Type            string    `json:"type" gorm:"type:varchar(50)"` // tacit, explicit, hybrid
	Domain          string    `json:"domain" gorm:"type:varchar(100)"`
	TacitKnowledge  string    `json:"tacit_knowledge" gorm:"type:text"`
	ExplicitForm    string    `json:"explicit_form" gorm:"type:text"`
	ConversionPath  JSON      `json:"conversion_path" gorm:"type:jsonb"` // SECIモデルのパス
	Accuracy        float64   `json:"accuracy"`
	Coverage        float64   `json:"coverage"`
	Consistency     float64   `json:"consistency"`
	AbstractLevel   string    `json:"abstract_level" gorm:"type:varchar(50)"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

// NOTE: OptimizationModel is already defined in robot_specification.go
// This definition has been removed to avoid duplicate declaration

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

// ProcessOptimization - プロセス最適化記録
type ProcessOptimization struct {
	ID              string    `gorm:"type:varchar(255);primaryKey" json:"id"`
	ProcessID       string    `json:"process_id" gorm:"index"`
	OptimizationType string   `json:"optimization_type"` // speed, accuracy, energy, cost
	InitialState    JSON      `json:"initial_state" gorm:"type:jsonb"`
	OptimizedState  JSON      `json:"optimized_state" gorm:"type:jsonb"`
	Improvement     float64   `json:"improvement"` // 改善率（%）
	Method          string    `json:"method" gorm:"type:text"`
	Iterations      int       `json:"iterations"`
	ConvergenceTime float64   `json:"convergence_time"` // 収束時間（秒）
	ValidatedBy     string    `json:"validated_by"`
	ValidationDate  time.Time `json:"validation_date"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

// TeachingFreeControl - ティーチング不要制御
type TeachingFreeControl struct {
	ID              string    `gorm:"type:varchar(255);primaryKey" json:"id"`
	RobotID         string    `json:"robot_id" gorm:"index"`
	TaskType        string    `json:"task_type"` // pick_place, assembly, welding, etc
	VisionSystem    JSON      `json:"vision_system" gorm:"type:jsonb"`
	ForceControl    JSON      `json:"force_control" gorm:"type:jsonb"`
	AIModel         JSON      `json:"ai_model" gorm:"type:jsonb"`
	LearningData    JSON      `json:"learning_data" gorm:"type:jsonb"`
	SuccessRate     float64   `json:"success_rate"`
	AdaptationTime  float64   `json:"adaptation_time"` // 新タスクへの適応時間
	ErrorRecovery   JSON      `json:"error_recovery" gorm:"type:jsonb"`
	PerformanceLog  JSON      `json:"performance_log" gorm:"type:jsonb"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

// LanguageOptimization - 言語最適化モデル
type LanguageOptimization struct {
	ID               string    `gorm:"type:varchar(255);primaryKey" json:"id"`
	OriginalText     string    `json:"original_text" gorm:"type:text"`
	OptimizedText    string    `json:"optimized_text" gorm:"type:text"`
	Domain           string    `json:"domain" gorm:"type:varchar(100)"`
	AbstractionLevel string    `json:"abstraction_level"` // L0-L3
	Precision        float64   `json:"precision"`
	Clarity          float64   `json:"clarity"`
	Completeness     float64   `json:"completeness"`
	Context          JSON      `json:"context" gorm:"type:jsonb"`
	Transformation   JSON      `json:"transformation" gorm:"type:jsonb"`
	EvaluationScore  float64   `json:"evaluation_score"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
	DeletedAt        gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}