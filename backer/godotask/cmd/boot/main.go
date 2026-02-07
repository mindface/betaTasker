package main

import (
	"github.com/godotask/infrastructure/router"
	"github.com/godotask/cmd/boot/initialize"
	"github.com/joho/godotenv"

	"github.com/rs/zerolog/log"
)

func main() {
	initialize.InitLog()

	log.Info().Msg("application starting")

	err := godotenv.Load()
	if err != nil {
		log.Error().Err(err).Msg("Error loading .env file: proceeding with environment variables")
	}
	initialize.InitDB()
	router.Init()

	router.GetRouter().Run(":8080")
}
