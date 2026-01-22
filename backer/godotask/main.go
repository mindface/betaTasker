package main

import (
	"log"

	"github.com/godotask/server"
	"github.com/godotask/model"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	model.InitDB()
	server.Init()

	server.GetRouter().Run(":8080")
}
