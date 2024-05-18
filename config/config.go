package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func InitConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func GetProxyDomain() string {
	return os.Getenv("DOMAIN")
}
