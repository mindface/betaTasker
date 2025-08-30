package main

import (
	"flag"
	"log"
	"os"

	"github.com/godotask/model"
	"github.com/godotask/seed"
	"github.com/joho/godotenv"
)

func main() {
	// フラグの定義
	var (
		clean = flag.Bool("clean", false, "Clean database before seeding")
		only  = flag.String("only", "", "Seed only specific data (e.g., 'heuristics')")
	)
	flag.Parse()

	// .envファイルの読み込み
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: .env file not found")
	}

	// データベース接続
	log.Println("Connecting to database...")
	model.InitDB()

	// シード実行
	var err error
	if *clean {
		log.Println("Running with --clean flag: cleaning and seeding database...")
		err = seed.CleanAndSeed()
	} else if *only != "" {
		log.Printf("Seeding only: %s\n", *only)
		err = runSpecificSeed(*only)
	} else {
		log.Println("Running seed...")
		err = seed.RunAllSeeds()
	}

	if err != nil {
		log.Fatalf("Seed failed: %v", err)
	}

	log.Println("✅ Seed completed successfully!")
}

func runSpecificSeed(seedType string) error {
	db := model.DB

	switch seedType {
	case "heuristics":
		return seed.SeedHeuristics(db)
	// case "memory":
	// 	return seed.SeedMemoryContexts()
	// case "books":
	// 	return seed.SeedBooksAndTasks()
	// 他のシードタイプも追加可能
	// case "users":
	//     return seed.SeedUsers(db)
	default:
		log.Fatalf("Unknown seed type: %s", seedType)
		os.Exit(1)
	}
	return nil
}