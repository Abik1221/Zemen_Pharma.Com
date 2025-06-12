package config

import (
	"log"

	"github.com/joho/godotenv"
)

func InitCinfig() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("error while loading the .env file")
	}
}
