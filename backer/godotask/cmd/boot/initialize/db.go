package initialize

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"

  "github.com/godotask/infrastructure/db/model"
)

func InitDB() {
	// dsn := os.Getenv("DATABASE_DSN")
	dsn := "host=db user=dbgodotask password=dbgodotask dbname=dbgodotask port=5432 sslmode=disable"
	var err error
  model.DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}
	err = model.DB.AutoMigrate(model.Models()...)
	if err != nil {
		log.Fatalf("failed to migrate models: %v", err)
	}
}
