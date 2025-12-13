package main

import (
	"fmt"
	"log"
	"github.com/godotask/model"
	"github.com/godotask/seed"
	// "github.com/seed2/memory_context_seed"
)

// RunAllSeeds - 全てのシードデータを実行
func main() {
	model.InitDB()
	db := model.DB

	CleanDatabase()

	log.Println("Starting database seeding...")

	// Users のシード
	// log.Println("Seeding users...")
	// if err := seed.SeedUsers(db); err != nil {
	// 	return fmt.Errorf("failed to seed users: %v", err)
	// }
	// log.Println("✓ Users seeded successfully")

	// Memory Contextsのシード（シンプル版）
	log.Println("Seeding memory contexts...")
	if err := seed.SeedMemoryContexts(db); err != nil {
		fmt.Errorf("failed to seed memory contexts: %v", err)
	}
	log.Println("✓ Memory contexts seeded successfully")

	// Books and Tasksのシード（seedModel.goの関数を使用）
	// log.Println("Seeding books and tasks...")
	// if err := seed.SeedBooksAndTasks(db); err != nil {
	// 	fmt.Errorf("failed to seed books and tasks: %v", err)
	// }
	// log.Println("✓ Books and tasks seeded successfully")

	// Tasks のシード
	log.Println("Seeding tasks...")
	if err := seed.SeedTaskModelsFromCSV(db); err != nil {
		fmt.Errorf("failed to seed tasks: %v", err)
	}
	log.Println("✓ Tasks seeded successfully")

	// Memories のシード
	log.Println("Seeding memories...")
	if err := seed.SeedMemoriesModelsFromCSV(db); err != nil {
		fmt.Errorf("failed to seed memories: %v", err)
	}
	log.Println("✓ memories seeded successfully")

	// Assessments のシード
	log.Println("Seeding assessment...")
	if err := seed.SeedAssessmentsModelsFromCSV(db); err != nil {
		fmt.Errorf("failed to seed assessments: %v", err)
	}
	log.Println("✓ assessments seeded successfully")

	// Books のシード
	log.Println("Seeding book...")
	if err := seed.SeedBookModelsFromCSV(db); err != nil {
		fmt.Errorf("failed to seed books: %v", err)
	}
	log.Println("✓ books seeded successfully")

	// OptimizationModels のシード
	log.Println("Seeding OptimizationModels...")
	if err := seed.SeedOptimizationModelsFromCSV(db); err != nil {
		fmt.Errorf("failed to seed OptimizationModels: %v", err)
	}
	log.Println("✓ OptimizationModels seeded successfully")

	// QualitativeLabels のシード
	log.Println("Seeding Qualitative Labels ...")
	if err := seed.SeedQualitativeLabelsFromCSV(db); err != nil {
		fmt.Errorf("failed to seed Qualitative Labels: %v", err)
	}
	log.Println("✓ Qualitative Labels seeded successfully")


	// PhenomenologicalFrameworks のシード
	log.Println("Seeding PhenomenologicalFrameworks...")
	if err := seed.SeedPhenomenologicalFrameworksFromCSV(db); err != nil {
		fmt.Errorf("failed to seed OptimizationModels: %v", err)
	}
	log.Println("✓ PhenomenologicalFrameworks seeded successfully")

	// QuantificationLabelsFromCSV のシード
	log.Println("Seeding QuantificationLabelsFromCSV...")
	if err := seed.SeedQuantificationLabelsFromCSV(db); err != nil {
		log.Printf("failed to seed QuantificationLabelsFromCSV: %v", err)
	}
	log.Println("✓ QuantificationLabels seeded successfully")

	// ヒューリスティクスデータのシード
	log.Println("Seeding heuristics data...")
	if err := seed.SeedHeuristics(db); err != nil {
		fmt.Errorf("failed to seed heuristics: %v", err)
	}
	log.Println("✓ Heuristics data seeded successfully")

	// 現象学的フレームワークデータのシード
	log.Println("Seeding phenomenological framework data...")
	if err := seed.SeedPhenomenologicalData(db); err != nil {
		fmt.Errorf("failed to seed phenomenological data: %v", err)
	}
	log.Println("✓ Phenomenological framework data seeded successfully")

	// 状態評価システムのシード
	log.Println("Seeding state evaluation data...")
	if err := seed.SeedStateEvaluations(db); err != nil {
		fmt.Errorf("failed to seed state evaluations: %v", err)
	}
	log.Println("✓ State evaluation data seeded successfully")

	// ツールマッチング結果のシード
	log.Println("Seeding tool matching results...")
	if err := seed.SeedToolMatchingResults(db); err != nil {
		fmt.Errorf("failed to seed tool matching results: %v", err)
	}
	log.Println("✓ Tool matching results seeded successfully")

	// プロセス監視のシード
	log.Println("Seeding process monitoring data...")
	if err := seed.SeedProcessMonitoring(db); err != nil {
		fmt.Errorf("failed to seed process monitoring: %v", err)
	}
	log.Println("✓ Process monitoring data seeded successfully")

	// 学習パターンのシード
	log.Println("Seeding learning patterns...")
	if err := seed.SeedLearningPatterns(db); err != nil {
		fmt.Errorf("failed to seed learning patterns: %v", err)
	}
	log.Println("✓ Learning patterns seeded successfully")

	// 学習パターンのシード
	log.Println("Seeding language optimization...")
	if err := seed.SeedLanguageOptimizationFromCSV(db); err != nil {
		fmt.Errorf("failed to seed learning patterns: %v", err)
	}
	log.Println("✓ language optimization seeded successfully")


	// 定量化ラベルのシード
	// log.Println("Seeding quantification labels...")
	// if err := seed.SeedQuantificationLabels(db); err != nil {
	// 	fmt.Errorf("failed to seed quantification labels: %v", err)
	// }
	// log.Println("✓ Quantification labels seeded successfully")

	// 知識エンティティのシード
	// log.Println("Seeding knowledge entities...")
	// if err := seed.SeedKnowledgeEntities(db); err != nil {
	// 	fmt.Errorf("failed to seed knowledge entities: %v", err)
	// }
	// log.Println("✓ Knowledge entities seeded successfully")

	log.Println("Database seeding completed successfully!")
}

// CleanAndSeed - データベースをクリーンアップしてからシード
// func CleanAndSeed() {
// 	db := model.DB

// 	log.Println("Cleaning database tables...")

// 	// テーブルのクリーンアップ（外部キー制約を考慮した逆順で削除）
// 	tables := []interface{}{
// 		&model.LearningPattern{},
// 		&model.ProcessMonitoring{},
// 		&model.ToolMatchingResult{},
// 		&model.StateEvaluation{},
// 		&model.QuantificationLabel{},
// 		&model.PhenomenologicalFramework{},
// 		&model.OptimizationModel{},
// 		// &model.RobotSpecification{},
// 		&model.HeuristicsModel{},
// 		&model.HeuristicsPattern{},
// 		&model.HeuristicsInsight{},
// 		&model.HeuristicsTracking{},
// 		&model.HeuristicsAnalysis{},
// 		&model.MemoryContext{},
// 		&model.Task{},
// 		&model.User{},
// 	}

// 	// for _, table := range tables {
// 	// 	if err := db.Exec("TRUNCATE TABLE ? RESTART IDENTITY CASCADE", table).Error; err != nil {
// 	// 		// TRUNCATEが失敗した場合は、DELETEを使用
// 	// 		if err := db.Delete(table, "1 = 1").Error; err != nil {
// 	// 			log.Printf("Warning: Failed to clean table: %v", err)
// 	// 		}
// 	// 	}
// 	// }

// 	for _, tableName := range tables {
//     if err := db.Exec("TRUNCATE TABLE " + tableName + " RESTART IDENTITY CASCADE").Error; err != nil {
//         log.Printf("TRUNCATE failed for table %s: %v", tableName, err)
//         if err := db.Delete(tableName, "1 = 1").Error; err != nil {
//             log.Printf("DELETE failed for table %s: %v", tableName, err)
//         }
//     }
// }

// 	log.Println("✓ Database cleaned")

// 	// シードデータの実行
// 	// return RunAllSeeds()
// }

// CleanDatabase - 外部キー制約を考慮し、全テーブルをクリーン
func CleanDatabase() {
	db := model.DB
	if db == nil {
		log.Fatal("DB is not initialized")
	}

	log.Println("Cleaning database tables...")

	// 外部キー依存を考慮して削除順序を逆に設定
	tables := []string{
		"learning_patterns",
		"process_monitorings",
		"tool_matching_results",
		"state_evaluations",
		"quantification_labels",
		"phenomenological_frameworks",
		"optimization_models",
		"robot_specifications",
		"heuristics_models",
		"heuristics_patterns",
		"heuristics_insights",
		"heuristics_trackings",
		"heuristics_analyses",
		"memory_contexts",
		"tasks",
		"memories",
		"assessments",
		"book",
	}

	for _, tableName := range tables {
		// TRUNCATE (PostgreSQLの場合)
		truncateSQL := fmt.Sprintf("TRUNCATE TABLE %s RESTART IDENTITY CASCADE", tableName)
		if err := db.Exec(truncateSQL).Error; err != nil {
			log.Printf("TRUNCATE failed for table %s: %v", tableName, err)
			
			// TRUNCATE が失敗した場合は DELETE でフォールバック
			deleteSQL := fmt.Sprintf("DELETE FROM %s", tableName)
			if err := db.Exec(deleteSQL).Error; err != nil {
				log.Printf("DELETE failed for table %s: %v", tableName, err)
			} else {
				log.Printf("DELETE succeeded for table %s", tableName)
			}
		} else {
			log.Printf("TRUNCATE succeeded for table %s", tableName)
		}
	}

	log.Println("✓ Database cleaned successfully")

	log.Println("✓ Database cleaned successfully")
}