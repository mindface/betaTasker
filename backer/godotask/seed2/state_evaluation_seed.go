package seed

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/godotask/model"
	"github.com/google/uuid"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

// SeedStateEvaluations - 状態評価のシードデータ
func SeedStateEvaluations(db *gorm.DB) error {
	log.Println("Starting state evaluations seeding...")

	// Sample state evaluations
	evaluations := []model.StateEvaluation{
		{
			ID:         uuid.New().String(),
			UserID:     "user_001",
			TaskID:     1,
			Level:      1,
			WorkTarget: "[職務カテゴリ: 切削・初品確認 / 分類コード: MA-C-01] 対象工程 L1-1: 初品加工・基本寸法確認",
			CurrentState: datatypes.JSON(`{
				"accuracy": 0.75,
				"efficiency": 0.68,
				"consistency": 0.72,
				"innovation": 0.45
			}`),
			TargetState: datatypes.JSON(`{
				"accuracy": 0.85,
				"efficiency": 0.75,
				"consistency": 0.80,
				"innovation": 0.55
			}`),
			EvaluationScore: 67.5,
			Framework:      "robot_precision_framework",
			Tools: datatypes.JSON(`{
				"machine": "NC旋盤（Mazak QT-200）",
				"material": "AISI 4340鋼材",
				"tool_spec": "TNMG160408 (汎用) 標準コーティング"
			}`),
			ProcessData: datatypes.JSON(`{
				"cutting_conditions": "メーカー推奨値",
				"inspection_method": "目視確認",
				"quality_check": "基本寸法確認"
			}`),
			LearnedKnowledge: "基本的な切削条件とバリ発生の関係を理解。目視確認の重要性を認識。",
			Status:          "completed",
			CreatedAt:       time.Now().Add(-24 * time.Hour),
			UpdatedAt:       time.Now().Add(-22 * time.Hour),
		},
		{
			ID:         uuid.New().String(),
			UserID:     "user_001",
			TaskID:     2,
			Level:      2,
			WorkTarget: "[職務カテゴリ: 品質改善・条件最適化 / 分類コード: MA-Q-02] 対象工程 L2-1: 材料硬度変動への対応",
			CurrentState: datatypes.JSON(`{
				"accuracy": 0.82,
				"efficiency": 0.71,
				"consistency": 0.79,
				"innovation": 0.63
			}`),
			TargetState: datatypes.JSON(`{
				"accuracy": 0.90,
				"efficiency": 0.80,
				"consistency": 0.85,
				"innovation": 0.70
			}`),
			EvaluationScore: 73.8,
			Framework:      "force_control_framework",
			Tools: datatypes.JSON(`{
				"machine": "NC旋盤（Mazak QT-200）",
				"material": "AISI 4340鋼材",
				"tool_spec": "汎用チップ (ISO CNMG) Al2O3コーティング"
			}`),
			ProcessData: datatypes.JSON(`{
				"hardness_range": "HRC 28-32",
				"adjustment": "切削速度10%減速",
				"optimization": "工具寿命延長方向"
			}`),
			LearnedKnowledge: "工具寿命と粗さのトレードオフを初めて意識。材料特性への適応の重要性を理解。",
			Status:          "completed",
			CreatedAt:       time.Now().Add(-18 * time.Hour),
			UpdatedAt:       time.Now().Add(-16 * time.Hour),
		},
		{
			ID:         uuid.New().String(),
			UserID:     "user_002",
			TaskID:     3,
			Level:      3,
			WorkTarget: "[職務カテゴリ: 高精度加工・品質保証 / 分類コード: MA-H-03] 対象工程 L3-1: 真円度0.005mm以下の高精度加工",
			CurrentState: datatypes.JSON(`{
				"accuracy": 0.88,
				"efficiency": 0.85,
				"consistency": 0.86,
				"innovation": 0.78
			}`),
			TargetState: datatypes.JSON(`{
				"accuracy": 0.95,
				"efficiency": 0.90,
				"consistency": 0.92,
				"innovation": 0.85
			}`),
			EvaluationScore: 84.2,
			Framework:      "precision_framework",
			Tools: datatypes.JSON(`{
				"machine": "5軸マシニングセンタ（DMG MORI）",
				"material": "チタン合金 Ti-6Al-4V",
				"tool_spec": "CBN工具 高速切削対応"
			}`),
			ProcessData: datatypes.JSON(`{
				"target_roundness": "0.005mm以下",
				"surface_roughness": "Ra0.4以下",
				"compensation": "温度補正と適応制御"
			}`),
			LearnedKnowledge: "熱変形の予測モデル構築の重要性を認識。高精度加工には環境制御が不可欠。",
			Status:          "completed",
			CreatedAt:       time.Now().Add(-12 * time.Hour),
			UpdatedAt:       time.Now().Add(-10 * time.Hour),
		},
		{
			ID:         uuid.New().String(),
			UserID:     "user_003",
			TaskID:     4,
			Level:      4,
			WorkTarget: "[職務カテゴリ: 難削材加工・特殊技術 / 分類コード: MA-S-04] 対象工程 L4-1: インコネル718の高効率加工",
			CurrentState: datatypes.JSON(`{
				"accuracy": 0.91,
				"efficiency": 0.88,
				"consistency": 0.89,
				"innovation": 0.85
			}`),
			TargetState: datatypes.JSON(`{
				"accuracy": 0.96,
				"efficiency": 0.93,
				"consistency": 0.94,
				"innovation": 0.90
			}`),
			EvaluationScore: 88.2,
			Framework:      "advanced_machining_framework",
			Tools: datatypes.JSON(`{
				"machine": "5軸マシニングセンタ（DMG MORI）",
				"material": "インコネル718",
				"tool_spec": "セラミック工具 高圧クーラント"
			}`),
			ProcessData: datatypes.JSON(`{
				"challenge": "急激な工具摩耗と加工硬化",
				"strategy": "断続切削と工具経路最適化",
				"cooling": "高圧クーラント適用"
			}`),
			LearnedKnowledge: "難削材には従来の常識が通用しない。革新的アプローチが必要。",
			Status:          "running",
			CreatedAt:       time.Now().Add(-6 * time.Hour),
			UpdatedAt:       time.Now().Add(-1 * time.Hour),
		},
		{
			ID:         uuid.New().String(),
			UserID:     "user_003",
			TaskID:     5,
			Level:      5,
			WorkTarget: "[職務カテゴリ: 技術指導・知識伝承 / 分類コード: MA-T-05] 対象工程 L5-1: 若手技術者への体系的指導",
			CurrentState: datatypes.JSON(`{
				"accuracy": 0.93,
				"efficiency": 0.91,
				"consistency": 0.92,
				"innovation": 0.89
			}`),
			TargetState: datatypes.JSON(`{
				"accuracy": 0.98,
				"efficiency": 0.95,
				"consistency": 0.96,
				"innovation": 0.94
			}`),
			EvaluationScore: 91.2,
			Framework:      "knowledge_transfer_framework",
			Tools: datatypes.JSON(`{
				"equipment": "全設備",
				"materials": "全材料",
				"teaching_tools": "指導用教材・シミュレータ"
			}`),
			ProcessData: datatypes.JSON(`{
				"challenge": "暗黙知の形式知化",
				"solution": "動画教材とVR訓練システム",
				"goal": "次世代技術者育成"
			}`),
			LearnedKnowledge: "技術伝承にはデジタルツールの活用が不可欠。体系的な指導方法論が重要。",
			Status:          "pending",
			CreatedAt:       time.Now().Add(-2 * time.Hour),
			UpdatedAt:       time.Now().Add(-1 * time.Hour),
		},
	}

	// Insert data
	for _, evaluation := range evaluations {
		if err := db.Create(&evaluation).Error; err != nil {
			log.Printf("Error inserting state evaluation %s: %v", evaluation.ID, err)
			return err
		}
	}

	log.Printf("✓ Successfully seeded %d state evaluations", len(evaluations))
	return nil
}

// SeedToolMatchingResults - ツールマッチング結果のシードデータ
func SeedToolMatchingResults(db *gorm.DB) error {
	log.Println("Starting tool matching results seeding...")

	// Get some state evaluations to reference
	var evaluations []model.StateEvaluation
	if err := db.Limit(3).Find(&evaluations).Error; err != nil {
		return fmt.Errorf("failed to get state evaluations for tool matching: %v", err)
	}

	if len(evaluations) == 0 {
		log.Println("No state evaluations found, skipping tool matching results")
		return nil
	}

	toolResults := []model.ToolMatchingResult{
		{
			ID:                uuid.New().String(),
			StateEvaluationID: evaluations[0].ID,
			RobotID:          "teaching_free_arm_v1",
			OptimizationModelID: "trajectory_optimization",
			MatchingScore:    0.87,
			Recommendations: datatypes.JSON(`{
				"robot": {
					"model": "TF-ARM-001",
					"dof": 6,
					"payload_capacity": 5.0,
					"reach": 850.0,
					"precision": 0.02,
					"recommended_use": "高精度作業に最適、AI学習機能搭載"
				},
				"optimization": {
					"model_name": "軌道最適化",
					"type": "control_theory",
					"objective": "minimize(time) + minimize(energy) subject to collision_free",
					"expected_improvement": "25%の改善が期待できます"
				},
				"process": {
					"setup_steps": ["視覚システムのキャリブレーション", "力センサの初期設定"],
					"success_criteria": {"target_score": 81.0, "accuracy_threshold": 0.85}
				}
			}`),
			Parameters: datatypes.JSON(`{
				"robot": {
					"max_payload": 4.0,
					"working_speed": 600.0,
					"precision_mode": true,
					"safety_factor": 1.5
				},
				"optimization": {
					"max_velocity": 1000,
					"max_acceleration": 5000,
					"sampling_time": 0.001
				}
			}`),
			ExpectedPerformance: datatypes.JSON(`{
				"predicted_score": 81.0,
				"confidence_level": "High (80-90%)",
				"estimated_timeline": "3 weeks",
				"detailed_metrics": {
					"accuracy_improvement": "15.0%",
					"efficiency_improvement": "20.0%",
					"consistency_improvement": "12.0%"
				}
			}`),
			CreatedAt: time.Now().Add(-20 * time.Hour),
		},
		{
			ID:                uuid.New().String(),
			StateEvaluationID: evaluations[1].ID,
			RobotID:          "collaborative_robot_v2",
			OptimizationModelID: "energy_optimization",
			MatchingScore:    0.92,
			Recommendations: datatypes.JSON(`{
				"robot": {
					"model": "COBOT-2024",
					"dof": 7,
					"payload_capacity": 10.0,
					"recommended_use": "AI学習機能搭載、人間協働対応"
				},
				"optimization": {
					"model_name": "エネルギー最適化",
					"type": "ml_based",
					"expected_improvement": "30%の改善が期待できます"
				}
			}`),
			Parameters: datatypes.JSON(`{
				"robot": {
					"max_payload": 8.0,
					"working_speed": 1200.0,
					"safety_factor": 1.3
				},
				"optimization": {
					"learning_rate": 0.001,
					"batch_size": 32,
					"epochs": 1000
				}
			}`),
			ExpectedPerformance: datatypes.JSON(`{
				"predicted_score": 88.5,
				"confidence_level": "Very High (90%+)",
				"estimated_timeline": "2 weeks"
			}`),
			CreatedAt: time.Now().Add(-15 * time.Hour),
		},
		{
			ID:                uuid.New().String(),
			StateEvaluationID: evaluations[2].ID,
			RobotID:          "precision_arm_micro",
			OptimizationModelID: "force_optimization",
			MatchingScore:    0.95,
			Recommendations: datatypes.JSON(`{
				"robot": {
					"model": "MICRO-100",
					"dof": 6,
					"payload_capacity": 1.0,
					"precision": 0.005,
					"recommended_use": "超高精度作業専用"
				},
				"optimization": {
					"model_name": "力制御最適化",
					"type": "hybrid",
					"expected_improvement": "40%の改善が期待できます"
				}
			}`),
			Parameters: datatypes.JSON(`{
				"robot": {
					"max_payload": 0.8,
					"working_speed": 300.0,
					"precision_mode": true
				},
				"optimization": {
					"kp": 100,
					"ki": 10,
					"kd": 5,
					"sampling_rate": 1000
				}
			}`),
			ExpectedPerformance: datatypes.JSON(`{
				"predicted_score": 93.8,
				"confidence_level": "Very High (90%+)",
				"estimated_timeline": "1 weeks"
			}`),
			CreatedAt: time.Now().Add(-10 * time.Hour),
		},
	}

	// Insert data
	for _, result := range toolResults {
		if err := db.Create(&result).Error; err != nil {
			log.Printf("Error inserting tool matching result %s: %v", result.ID, err)
			return err
		}
	}

	log.Printf("✓ Successfully seeded %d tool matching results", len(toolResults))
	return nil
}

// SeedProcessMonitoring - プロセス監視のシードデータ
func SeedProcessMonitoring(db *gorm.DB) error {
	log.Println("Starting process monitoring seeding...")

	// Get some state evaluations to reference
	var evaluations []model.StateEvaluation
	if err := db.Limit(2).Find(&evaluations).Error; err != nil {
		return fmt.Errorf("failed to get state evaluations for process monitoring: %v", err)
	}

	if len(evaluations) == 0 {
		log.Println("No state evaluations found, skipping process monitoring")
		return nil
	}

	monitoringRecords := []model.ProcessMonitoring{
		{
			ID:                uuid.New().String(),
			StateEvaluationID: evaluations[0].ID,
			ProcessType:       "robot_assembly",
			MonitoringData: datatypes.JSON(`{
				"force_x": 5.2,
				"force_y": 3.1,
				"force_z": 15.8,
				"position_error": 0.015,
				"cycle_time": 27.3,
				"success_rate": 0.96,
				"temperature": 42.5,
				"parts_assembled": 145
			}`),
			Metrics: datatypes.JSON(`{
				"efficiency": 0.89,
				"quality": 0.96,
				"stability": 0.92,
				"overall": 0.92
			}`),
			Anomalies: datatypes.JSON(`[
				{
					"type": "force_variation",
					"severity": "medium",
					"description": "Force variation exceeded threshold",
					"value": 12.5,
					"threshold": 10.0,
					"timestamp": "2024-01-15T10:30:00Z"
				}
			]`),
			Status:    "stopped",
			StartTime: time.Now().Add(-4 * time.Hour),
			EndTime:   func() *time.Time { t := time.Now().Add(-1 * time.Hour); return &t }(),
			CreatedAt: time.Now().Add(-4 * time.Hour),
			UpdatedAt: time.Now().Add(-1 * time.Hour),
		},
		{
			ID:                uuid.New().String(),
			StateEvaluationID: evaluations[1].ID,
			ProcessType:       "robot_welding",
			MonitoringData: datatypes.JSON(`{
				"current": 185.2,
				"voltage": 26.8,
				"speed": 9.5,
				"penetration": 3.7,
				"quality_score": 0.94,
				"arc_stability": 0.91,
				"spatter_count": 2,
				"weld_length": 2850
			}`),
			Metrics: datatypes.JSON(`{
				"efficiency": 0.91,
				"quality": 0.94,
				"stability": 0.88,
				"overall": 0.91
			}`),
			Anomalies: datatypes.JSON(`[
				{
					"type": "current_variation",
					"severity": "low",
					"description": "Current slightly above optimal range",
					"value": 185.2,
					"threshold": 180.0,
					"timestamp": "2024-01-15T14:15:00Z"
				}
			]`),
			Status:    "running",
			StartTime: time.Now().Add(-2 * time.Hour),
			CreatedAt: time.Now().Add(-2 * time.Hour),
			UpdatedAt: time.Now().Add(-30 * time.Minute),
		},
	}

	// Insert data
	for _, monitoring := range monitoringRecords {
		if err := db.Create(&monitoring).Error; err != nil {
			log.Printf("Error inserting process monitoring %s: %v", monitoring.ID, err)
			return err
		}
	}

	log.Printf("✓ Successfully seeded %d process monitoring records", len(monitoringRecords))
	return nil
}

// SeedLearningPatterns - 学習パターンのシードデータ
func SeedLearningPatterns(db *gorm.DB) error {
	log.Println("Starting learning patterns seeding...")

	// Read from CSV if exists
	patterns, err := readLearningPatternsFromCSV()
	if err != nil {
		log.Printf("Could not read CSV, using default data: %v", err)
		patterns = getDefaultLearningPatterns()
	}

	// Insert data
	for _, pattern := range patterns {
		if err := db.Create(&pattern).Error; err != nil {
			log.Printf("Error inserting learning pattern %s: %v", pattern.ID, err)
			return err
		}
	}

	log.Printf("✓ Successfully seeded %d learning patterns", len(patterns))
	return nil
}

func readLearningPatternsFromCSV() ([]model.LearningPattern, error) {
	file, err := os.Open("seed/data/knowledge_patterns.csv")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var patterns []model.LearningPattern
	for i, record := range records {
		if i == 0 { // Skip header
			continue
		}

		if len(record) < 12 {
			continue
		}

		accuracy, _ := strconv.ParseFloat(record[7], 64)
		coverage, _ := strconv.ParseFloat(record[8], 64)
		consistency, _ := strconv.ParseFloat(record[9], 64)
		validated, _ := strconv.ParseBool(record[11])

		pattern := model.LearningPattern{
			ID:             uuid.New().String(),
			UserID:         "user_001", // Default user
			PatternType:    record[1],
			Domain:         record[2],
			TacitKnowledge: record[3],
			ExplicitForm:   record[4],
			SECIStage:      record[5],
			Method:         record[6],
			Accuracy:       accuracy,
			Coverage:       coverage,
			Consistency:    consistency,
			AbstractLevel:  record[10],
			Validated:      validated,
			CreatedAt:      time.Now().Add(-24 * time.Hour),
			UpdatedAt:      time.Now().Add(-24 * time.Hour),
		}
		patterns = append(patterns, pattern)
	}

	return patterns, nil
}

func getDefaultLearningPatterns() []model.LearningPattern {
	return []model.LearningPattern{
		{
			ID:             uuid.New().String(),
			UserID:         "user_001",
			PatternType:    "assembly_skill_pattern",
			Domain:         "robot_assembly",
			TacitKnowledge: "熟練工の『しっくりくる』感覚",
			ExplicitForm:   "力覚センサ値: Fx<0.5N Fy<0.5N Tz<0.1Nm",
			SECIStage:      "共同化→表出化→連結化→内面化",
			Method:         "力覚データ記録→パターン分析→閾値設定",
			Accuracy:       0.85,
			Coverage:       0.75,
			Consistency:    0.90,
			AbstractLevel:  "L1",
			Validated:      true,
			CreatedAt:      time.Now().Add(-48 * time.Hour),
			UpdatedAt:      time.Now().Add(-48 * time.Hour),
		},
		{
			ID:             uuid.New().String(),
			UserID:         "user_002",
			PatternType:    "vision_recognition_pattern",
			Domain:         "robot_vision",
			TacitKnowledge: "対象物の見分け方",
			ExplicitForm:   "深層学習モデル: ResNet-50 mAP=0.92",
			SECIStage:      "表出化→連結化",
			Method:         "画像収集→アノテーション→学習→検証",
			Accuracy:       0.92,
			Coverage:       0.88,
			Consistency:    0.95,
			AbstractLevel:  "L0",
			Validated:      true,
			CreatedAt:      time.Now().Add(-36 * time.Hour),
			UpdatedAt:      time.Now().Add(-36 * time.Hour),
		},
		{
			ID:             uuid.New().String(),
			UserID:         "user_003",
			PatternType:    "quality_inspection_pattern",
			Domain:         "robot_inspection",
			TacitKnowledge: "検査員の違和感察知",
			ExplicitForm:   "異常検知アルゴリズム: Isolation Forest threshold=0.05",
			SECIStage:      "共同化→表出化→連結化",
			Method:         "異常サンプル収集→特徴量抽出→モデル学習",
			Accuracy:       0.94,
			Coverage:       0.89,
			Consistency:    0.96,
			AbstractLevel:  "L1",
			Validated:      true,
			CreatedAt:      time.Now().Add(-24 * time.Hour),
			UpdatedAt:      time.Now().Add(-24 * time.Hour),
		},
	}
}