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

	r := server.GetRouter()
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
