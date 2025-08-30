package seed

import (
	"encoding/json"
	"fmt"
	"time"
	"github.com/godotask/model"
	"gorm.io/gorm"
)

// SeedHeuristics - ヒューリスティクスデータのシード
func SeedHeuristics(db *gorm.DB) error {
	// 分析データのシード
	if err := seedHeuristicsAnalysis(db); err != nil {
		return fmt.Errorf("failed to seed heuristics analysis: %v", err)
	}

	// トラッキングデータのシード
	if err := seedHeuristicsTracking(db); err != nil {
		return fmt.Errorf("failed to seed heuristics tracking: %v", err)
	}

	// インサイトデータのシード
	if err := seedHeuristicsInsights(db); err != nil {
		return fmt.Errorf("failed to seed heuristics insights: %v", err)
	}

	// パターンデータのシード
	if err := seedHeuristicsPatterns(db); err != nil {
		return fmt.Errorf("failed to seed heuristics patterns: %v", err)
	}

	// モデルデータのシード
	if err := seedHeuristicsModels(db); err != nil {
		return fmt.Errorf("failed to seed heuristics models: %v", err)
	}

	return nil
}

func seedHeuristicsAnalysis(db *gorm.DB) error {
	analyses := []model.HeuristicsAnalysis{
		{
			UserID:       1,
			TaskID:       1,
			AnalysisType: "performance",
			Result:       toJSON(map[string]interface{}{
				"completion_rate": 0.85,
				"accuracy": 0.92,
				"speed": "fast",
				"errors": []string{},
			}),
			Score:     87.5,
			Status:    "completed",
			CreatedAt: time.Now().AddDate(0, 0, -7),
		},
		{
			UserID:       1,
			TaskID:       2,
			AnalysisType: "behavioral",
			Result:       toJSON(map[string]interface{}{
				"pattern": "consistent",
				"focus_time": 45,
				"break_frequency": 3,
				"productivity_peak": "morning",
			}),
			Score:     92.3,
			Status:    "completed",
			CreatedAt: time.Now().AddDate(0, 0, -5),
		},
		{
			UserID:       2,
			TaskID:       3,
			AnalysisType: "cognitive_load",
			Result:       toJSON(map[string]interface{}{
				"complexity": "high",
				"stress_level": 0.6,
				"attention_span": 35,
				"multitasking": false,
			}),
			Score:     78.9,
			Status:    "completed",
			CreatedAt: time.Now().AddDate(0, 0, -3),
		},
	}

	for _, analysis := range analyses {
		if err := db.Create(&analysis).Error; err != nil {
			return err
		}
	}
	return nil
}

func seedHeuristicsTracking(db *gorm.DB) error {
	trackings := []model.HeuristicsTracking{
		{
			UserID:    1,
			Action:    "task_started",
			Context:   toJSON(map[string]interface{}{
				"task_id": 1,
				"task_type": "development",
				"environment": "vscode",
			}),
			SessionID: "sess_001",
			Timestamp: time.Now().AddDate(0, 0, -7),
			Duration:  0,
		},
		{
			UserID:    1,
			Action:    "code_written",
			Context:   toJSON(map[string]interface{}{
				"lines": 45,
				"language": "go",
				"file": "controller.go",
			}),
			SessionID: "sess_001",
			Timestamp: time.Now().AddDate(0, 0, -7).Add(15 * time.Minute),
			Duration:  900000, // 15分
		},
		{
			UserID:    1,
			Action:    "test_run",
			Context:   toJSON(map[string]interface{}{
				"test_count": 12,
				"passed": 10,
				"failed": 2,
			}),
			SessionID: "sess_001",
			Timestamp: time.Now().AddDate(0, 0, -7).Add(30 * time.Minute),
			Duration:  120000, // 2分
		},
		{
			UserID:    2,
			Action:    "document_read",
			Context:   toJSON(map[string]interface{}{
				"document": "API_GUIDE.md",
				"section": "authentication",
				"scroll_depth": 0.75,
			}),
			SessionID: "sess_002",
			Timestamp: time.Now().AddDate(0, 0, -5),
			Duration:  300000, // 5分
		},
	}

	for _, tracking := range trackings {
		if err := db.Create(&tracking).Error; err != nil {
			return err
		}
	}
	return nil
}

func seedHeuristicsInsights(db *gorm.DB) error {
	insights := []model.HeuristicsInsight{
		{
			UserID:      1,
			Type:        "productivity",
			Title:       "朝の生産性が高い",
			Description: "過去30日間のデータ分析により、午前9時から11時の間に最も高い生産性を示しています",
			Confidence:  0.89,
			Data:        toJSON(map[string]interface{}{
				"peak_hours": []int{9, 10, 11},
				"average_output": 1.45,
				"comparison": "+23% vs afternoon",
			}),
			IsActive:  true,
			CreatedAt: time.Now().AddDate(0, 0, -3),
		},
		{
			UserID:      1,
			Type:        "learning",
			Title:       "Go言語スキルが向上中",
			Description: "コード品質とテスト成功率が継続的に改善されています",
			Confidence:  0.92,
			Data:        toJSON(map[string]interface{}{
				"improvement_rate": 0.15,
				"error_reduction": 0.32,
				"test_coverage": 0.78,
			}),
			IsActive:  true,
			CreatedAt: time.Now().AddDate(0, 0, -2),
		},
		{
			UserID:      2,
			Type:        "workflow",
			Title:       "頻繁な休憩が効果的",
			Description: "25分の作業後に5分の休憩を取ることで、全体的な生産性が向上しています",
			Confidence:  0.76,
			Data:        toJSON(map[string]interface{}{
				"optimal_work_duration": 25,
				"optimal_break_duration": 5,
				"productivity_increase": 0.18,
			}),
			IsActive:  true,
			CreatedAt: time.Now().AddDate(0, 0, -1),
		},
	}

	for _, insight := range insights {
		if err := db.Create(&insight).Error; err != nil {
			return err
		}
	}
	return nil
}

func seedHeuristicsPatterns(db *gorm.DB) error {
	patterns := []model.HeuristicsPattern{
		{
			Name:     "Morning Productivity Peak",
			Category: "time_management",
			Pattern:  toJSON(map[string]interface{}{
				"time_range": "09:00-11:00",
				"days": []string{"Mon", "Tue", "Wed", "Thu", "Fri"},
				"task_types": []string{"coding", "design", "analysis"},
			}),
			Frequency: 85,
			Accuracy:  0.91,
			LastSeen:  time.Now().AddDate(0, 0, -1),
			CreatedAt: time.Now().AddDate(0, -1, 0),
		},
		{
			Name:     "Test-Driven Development",
			Category: "development",
			Pattern:  toJSON(map[string]interface{}{
				"sequence": []string{"write_test", "write_code", "refactor"},
				"average_cycle_time": 45,
				"success_rate": 0.88,
			}),
			Frequency: 72,
			Accuracy:  0.85,
			LastSeen:  time.Now().AddDate(0, 0, -2),
			CreatedAt: time.Now().AddDate(0, -1, -15),
		},
		{
			Name:     "Documentation After Implementation",
			Category: "workflow",
			Pattern:  toJSON(map[string]interface{}{
				"trigger": "feature_complete",
				"action": "write_documentation",
				"delay_minutes": 30,
			}),
			Frequency: 63,
			Accuracy:  0.79,
			LastSeen:  time.Now().AddDate(0, 0, -3),
			CreatedAt: time.Now().AddDate(0, -2, 0),
		},
	}

	for _, pattern := range patterns {
		if err := db.Create(&pattern).Error; err != nil {
			return err
		}
	}
	return nil
}

func seedHeuristicsModels(db *gorm.DB) error {
	models := []model.HeuristicsModel{
		{
			ModelType: "productivity_predictor",
			Version:   "1.2.0",
			Parameters: toJSON(map[string]interface{}{
				"algorithm": "random_forest",
				"features": 24,
				"trees": 100,
				"max_depth": 10,
			}),
			Performance: toJSON(map[string]interface{}{
				"accuracy": 0.87,
				"precision": 0.85,
				"recall": 0.89,
				"f1_score": 0.87,
			}),
			Status:    "ready",
			TrainedAt: time.Now().AddDate(0, 0, -14),
			CreatedAt: time.Now().AddDate(0, 0, -14),
		},
		{
			ModelType: "pattern_detector",
			Version:   "2.0.1",
			Parameters: toJSON(map[string]interface{}{
				"algorithm": "lstm",
				"sequence_length": 50,
				"hidden_units": 128,
				"layers": 3,
			}),
			Performance: toJSON(map[string]interface{}{
				"accuracy": 0.91,
				"loss": 0.12,
				"validation_accuracy": 0.89,
			}),
			Status:    "ready",
			TrainedAt: time.Now().AddDate(0, 0, -7),
			CreatedAt: time.Now().AddDate(0, 0, -7),
		},
		{
			ModelType: "cognitive_load_estimator",
			Version:   "1.0.0",
			Parameters: toJSON(map[string]interface{}{
				"algorithm": "gradient_boosting",
				"estimators": 150,
				"learning_rate": 0.1,
				"max_features": "sqrt",
			}),
			Performance: toJSON(map[string]interface{}{
				"mae": 0.15,
				"rmse": 0.22,
				"r2_score": 0.83,
			}),
			Status:    "training",
			TrainedAt: time.Now(),
			CreatedAt: time.Now().AddDate(0, 0, -1),
		},
	}

	for _, model := range models {
		if err := db.Create(&model).Error; err != nil {
			return err
		}
	}
	return nil
}

// Helper function to convert map to JSON string
func toJSON(data map[string]interface{}) string {
	jsonData, _ := json.Marshal(data)
	return string(jsonData)
}