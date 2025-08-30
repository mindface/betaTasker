package seed

import (
	"fmt"
	"log"
	"github.com/godotask/model"
)

// RunAllSeeds - 全てのシードデータを実行
func RunAllSeeds() error {
	db := model.DB

	log.Println("Starting database seeding...")

	// Memory Contextsのシード
	log.Println("Seeding memory contexts...")
	if err := SeedMemoryContexts(); err != nil {
		return fmt.Errorf("failed to seed memory contexts: %v", err)
	}
	log.Println("✓ Memory contexts seeded successfully")

	// Books and Tasksのシード
	log.Println("Seeding books and tasks...")
	if err := SeedBooksAndTasks(); err != nil {
		return fmt.Errorf("failed to seed books and tasks: %v", err)
	}
	log.Println("✓ Books and tasks seeded successfully")

	// ヒューリスティクスデータのシード
	log.Println("Seeding heuristics data...")
	if err := SeedHeuristics(db); err != nil {
		return fmt.Errorf("failed to seed heuristics: %v", err)
	}
	log.Println("✓ Heuristics data seeded successfully")

	log.Println("Database seeding completed successfully!")
	return nil
}

// CleanAndSeed - データベースをクリーンアップしてからシード
func CleanAndSeed() error {
	db := model.DB

	log.Println("Cleaning database tables...")

	// Heuristicsテーブルのクリーンアップ（逆順で削除）
	tables := []interface{}{
		&model.HeuristicsModel{},
		&model.HeuristicsPattern{},
		&model.HeuristicsInsight{},
		&model.HeuristicsTracking{},
		&model.HeuristicsAnalysis{},
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
	return RunAllSeeds()
}