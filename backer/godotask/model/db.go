package model

import (
	"os"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func InitDB() {
	dsn := os.Getenv("DATABASE_DSN")
	// dsn := "host=db user=dbgodotask password=dbgodotask dbname=test_db port=5432 sslmode=disable"
	var err error
  DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}
	err = DB.AutoMigrate(Models()...)
	if err != nil {
		log.Fatalf("failed to migrate models: %v", err)
	}
}
