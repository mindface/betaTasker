package model

import (
	"time"
	"gorm.io/datatypes"
)

type StateEvaluation struct {
	ID                   string                 `gorm:"type:varchar(255);primaryKey" json:"id"`
	UserID              string                 `gorm:"type:varchar(255);not null" json:"user_id"`
	TaskID              int                    `gorm:"not null" json:"task_id"`
	Level               int                    `gorm:"not null" json:"level"`
	WorkTarget          string                 `gorm:"type:text" json:"work_target"`
	CurrentState        datatypes.JSON         `gorm:"type:jsonb" json:"current_state"`
	TargetState         datatypes.JSON         `gorm:"type:jsonb" json:"target_state"`
	EvaluationScore     float64                `gorm:"type:decimal(5,2)" json:"evaluation_score"`
	Framework           string                 `gorm:"type:varchar(255)" json:"framework"`
	Tools               datatypes.JSON         `gorm:"type:jsonb" json:"tools"`
	ProcessData         datatypes.JSON         `gorm:"type:jsonb" json:"process_data"`
	Results             datatypes.JSON         `gorm:"type:jsonb" json:"results"`
	LearnedKnowledge    string                 `gorm:"type:text" json:"learned_knowledge"`
	Status              string                 `gorm:"type:varchar(50);default:'pending'" json:"status"`
	CreatedAt           time.Time              `json:"created_at"`
	UpdatedAt           time.Time              `json:"updated_at"`
}

type ToolMatchingResult struct {
	ID                  string                 `gorm:"type:varchar(255);primaryKey" json:"id"`
	StateEvaluationID   string                 `gorm:"type:varchar(255);not null" json:"state_evaluation_id"`
	RobotID             string                 `gorm:"type:varchar(255)" json:"robot_id"`
	OptimizationModelID string                 `gorm:"type:varchar(255)" json:"optimization_model_id"`
	MatchingScore       float64                `gorm:"type:decimal(5,3)" json:"matching_score"`
	Recommendations     datatypes.JSON         `gorm:"type:jsonb" json:"recommendations"`
	Parameters          datatypes.JSON         `gorm:"type:jsonb" json:"parameters"`
	ExpectedPerformance datatypes.JSON         `gorm:"type:jsonb" json:"expected_performance"`
	CreatedAt           time.Time              `json:"created_at"`
}

type ProcessMonitoring struct {
	ID                  string                 `gorm:"type:varchar(255);primaryKey" json:"id"`
	StateEvaluationID   string                 `gorm:"type:varchar(255);not null" json:"state_evaluation_id"`
	ProcessType         string                 `gorm:"type:varchar(100)" json:"process_type"`
	MonitoringData      datatypes.JSON         `gorm:"type:jsonb" json:"monitoring_data"`
	Metrics             datatypes.JSON         `gorm:"type:jsonb" json:"metrics"`
	Anomalies           datatypes.JSON         `gorm:"type:jsonb" json:"anomalies"`
	Status              string                 `gorm:"type:varchar(50)" json:"status"`
	StartTime           time.Time              `json:"start_time"`
	EndTime             *time.Time             `json:"end_time"`
	CreatedAt           time.Time              `json:"created_at"`
	UpdatedAt           time.Time              `json:"updated_at"`
}

type LearningPattern struct {
	ID                string                 `gorm:"type:varchar(255);primaryKey" json:"id"`
	UserID            string                 `gorm:"type:varchar(255);not null" json:"user_id"`
	PatternType       string                 `gorm:"type:varchar(100)" json:"pattern_type"`
	Domain            string                 `gorm:"type:varchar(100)" json:"domain"`
	TacitKnowledge    string                 `gorm:"type:text" json:"tacit_knowledge"`
	ExplicitForm      string                 `gorm:"type:text" json:"explicit_form"`
	SECIStage         string                 `gorm:"type:varchar(50)" json:"seci_stage"`
	Method            string                 `gorm:"type:varchar(100)" json:"method"`
	Accuracy          float64                `gorm:"type:decimal(5,3)" json:"accuracy"`
	Coverage          float64                `gorm:"type:decimal(5,3)" json:"coverage"`
	Consistency       float64                `gorm:"type:decimal(5,3)" json:"consistency"`
	AbstractLevel     string                 `gorm:"type:varchar(10)" json:"abstract_level"`
	Validated         bool                   `gorm:"default:false" json:"validated"`
	CreatedAt         time.Time              `json:"created_at"`
	UpdatedAt         time.Time              `json:"updated_at"`
}

type EvaluationRequest struct {
	UserID          string                 `json:"user_id" binding:"required"`
	TaskID          int                    `json:"task_id" binding:"required"`
	Level           int                    `json:"level" binding:"required"`
	WorkTarget      string                 `json:"work_target" binding:"required"`
	CurrentState    map[string]interface{} `json:"current_state" binding:"required"`
	TargetState     map[string]interface{} `json:"target_state" binding:"required"`
	Framework       string                 `json:"framework"`
}

type ToolMatchingRequest struct {
	StateEvaluationID string                 `json:"state_evaluation_id" binding:"required"`
	Requirements      map[string]interface{} `json:"requirements"`
	Constraints       map[string]interface{} `json:"constraints"`
	Preferences       map[string]interface{} `json:"preferences"`
}

type ProcessMonitoringRequest struct {
	StateEvaluationID string                 `json:"state_evaluation_id" binding:"required"`
	ProcessType       string                 `json:"process_type" binding:"required"`
	InitialData       map[string]interface{} `json:"initial_data"`
}