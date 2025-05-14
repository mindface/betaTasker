package model

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)


var DB *gorm.DB

func InitDB() {
	dsn := "root:dbgodotask@tcp(dbgodotask:3306)/dbgodotask?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}
}
