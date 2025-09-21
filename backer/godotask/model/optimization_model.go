package model

import "time"

type OptimizationModel struct {
	ID                string    `gorm:"primaryKey" json:"id"`
	Name              string    `json:"name"`
	Type              string    `json:"type"`
	ObjectiveFunction string    `json:"objective_function"`
	Constraints       string    `json:"constraints"`
	Parameters        string    `json:"parameters"`
	PerformanceMetric string    `json:"performance_metric"`
	IterationCount    float64   `json:"iteration_count"`
	ConvergenceRate   float64   `json:"convergence_rate"`
	Domain            string    `json:"domain"`
	Application       string    `json:"application"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`

	// 補助情報（関連テーブルがある場合の例）
	// ParametersList []OptimizationParameter `gorm:"foreignKey:ModelID" json:"parameters_list"`
}
