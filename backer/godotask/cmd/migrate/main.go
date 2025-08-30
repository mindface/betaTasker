package main

import (
	"log"
	"os"

	"github.com/godotask/model"
	"github.com/joho/godotenv"
)

func main() {
	// .envファイルの読み込み
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: .env file not found")
	}

	// データベース接続と初期化
	log.Println("Connecting to database and running migrations...")
	model.InitDB()

	// InitDBが既にAutoMigrateを実行しているので、追加の処理は不要
	log.Println("✅ Database migration completed successfully!")
	
	// テーブル確認
	var tableNames []string
	model.DB.Raw("SELECT table_name FROM information_schema.tables WHERE table_schema = 'public' ORDER BY table_name").Scan(&tableNames)
	
	log.Println("Created tables:")
	for _, table := range tableNames {
		log.Printf("  - %s", table)
	}

	os.Exit(0)
}