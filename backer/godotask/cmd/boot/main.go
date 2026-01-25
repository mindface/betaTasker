package main

import (
	"log"

	"github.com/godotask/infrastructure/router"
	"github.com/godotask/infrastructure/db/model"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	model.InitDB()
	router.Init()

	router.GetRouter().Run(":8080")
}
