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
	var err error
  DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}
}
