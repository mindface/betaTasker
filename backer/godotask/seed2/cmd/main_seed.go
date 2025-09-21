package main

import (
	"fmt"
	"log"
	"github.com/godotask/model"
	"github.com/godotask/seed"
	// "github.com/seedcopy/memory_context_seed"
	// "github.com/seedcopy/task_seed"
	// "github.com/godotask/seedcopy/book_task_seed"
	// "github.com/godotask/seedcopy/heuristics_seed"
	// "github.com/godotask/seedcopy/phenomenological_framework_seed"
	// "github.com/godotask/seedcopy/state_evaluation_seed"
	// "github.com/godotask/seedcopy/tool_matching_result_seed"
	// "github.com/godotask/seedcopy/process_monitoring_seed"
	// "github.com/godotask/seedcopy/learning_pattern_seed"
	// "github.com/seedcopy/quantification_label_seed"
)

// RunAllSeeds - 全てのシードデータを実行
func main() {
	db := model.DB

	log.Println("Starting database seeding...")

	// Users のシード
	// log.Println("Seeding users...")
	// if err := seed.SeedUsers(db); err != nil {
	// 	return fmt.Errorf("failed to seed users: %v", err)
	// }
	// log.Println("✓ Users seeded successfully")

	// Tasks のシード
	log.Println("Seeding tasks...")
	if err := seed.SeedTasks(db); err != nil {
		return fmt.Errorf("failed to seed tasks: %v", err)
	}
	log.Println("✓ Tasks seeded successfully")

	// Memory Contextsのシード（シンプル版）
	log.Println("Seeding memory contexts...")
	if err := seed.SeedMemoryContexts(db); err != nil {
		return fmt.Errorf("failed to seed memory contexts: %v", err)
	}
	log.Println("✓ Memory contexts seeded successfully")

	// Books and Tasksのシード（seedModel.goの関数を使用）
	// log.Println("Seeding books and tasks...")
	// if err := seed.SeedBooksAndTasks(db); err != nil {
	// 	return fmt.Errorf("failed to seed books and tasks: %v", err)
	// }
	// log.Println("✓ Books and tasks seeded successfully")

	// ヒューリスティクスデータのシード
	log.Println("Seeding heuristics data...")
	if err := seed.SeedHeuristics(db); err != nil {
		return fmt.Errorf("failed to seed heuristics: %v", err)
	}
	log.Println("✓ Heuristics data seeded successfully")

	// 現象学的フレームワークデータのシード
	log.Println("Seeding phenomenological framework data...")
	if err := seed.SeedPhenomenologicalData(db); err != nil {
		return fmt.Errorf("failed to seed phenomenological data: %v", err)
	}
	log.Println("✓ Phenomenological framework data seeded successfully")

	// CSVファイルからのデータシード
	// log.Println("Seeding data from CSV files...")
	// if err := SeedFromCSVFiles(db); err != nil {
	// 	return fmt.Errorf("failed to seed CSV data: %v", err)
	// }
	log.Println("✓ CSV data seeded successfully")

	// 状態評価システムのシード
	log.Println("Seeding state evaluation data...")
	if err := seed.SeedStateEvaluations(db); err != nil {
		return fmt.Errorf("failed to seed state evaluations: %v", err)
	}
	log.Println("✓ State evaluation data seeded successfully")

	// ツールマッチング結果のシード
	log.Println("Seeding tool matching results...")
	if err := seed.SeedToolMatchingResults(db); err != nil {
		return fmt.Errorf("failed to seed tool matching results: %v", err)
	}
	log.Println("✓ Tool matching results seeded successfully")

	// プロセス監視のシード
	log.Println("Seeding process monitoring data...")
	if err := seed.SeedProcessMonitoring(db); err != nil {
		return fmt.Errorf("failed to seed process monitoring: %v", err)
	}
	log.Println("✓ Process monitoring data seeded successfully")

	// 学習パターンのシード
	log.Println("Seeding learning patterns...")
	if err := seed.SeedLearningPatterns(db); err != nil {
		return fmt.Errorf("failed to seed learning patterns: %v", err)
	}
	log.Println("✓ Learning patterns seeded successfully")

	// 定量化ラベルのシード
	log.Println("Seeding quantification labels...")
	if err := seed.SeedQuantificationLabels(db); err != nil {
		return fmt.Errorf("failed to seed quantification labels: %v", err)
	}
	log.Println("✓ Quantification labels seeded successfully")

	log.Println("Database seeding completed successfully!")
	return nil
}

// CleanAndSeed - データベースをクリーンアップしてからシード
func CleanAndSeed() error {
	db := model.DB

	log.Println("Cleaning database tables...")

	// テーブルのクリーンアップ（外部キー制約を考慮した逆順で削除）
	tables := []interface{}{
		&model.LearningPattern{},
		&model.ProcessMonitoring{},
		&model.ToolMatchingResult{},
		&model.StateEvaluation{},
		&model.QuantificationLabel{},
		&model.PhenomenologicalFramework{},
		&model.OptimizationModel{},
		// &model.RobotSpecification{},
		&model.HeuristicsModel{},
		&model.HeuristicsPattern{},
		&model.HeuristicsInsight{},
		&model.HeuristicsTracking{},
		&model.HeuristicsAnalysis{},
		&model.MemoryContext{},
		&model.Task{},
		&model.User{},
	}

	for _, table := range tables {
		if err := db.Exec("TRUNCATE TABLE ? RESTART IDENTITY CASCADE", table).Error; err != nil {
			// TRUNCATEが失敗した場合は、DELETEを使用
			if err := db.Delete(table, "1 = 1").Error; err != nil {
				log.Printf("Warning: Failed to clean table: %v", err)
			}
		}
	}

	log.Println("✓ Database cleaned")

	// シードデータの実行
	// return RunAllSeeds()
}