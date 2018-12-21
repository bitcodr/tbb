package config

import (
	"log"
	"github.com/joho/godotenv"
)

//Environment load .env
func Environment() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error in loading .env file")
	}
}
