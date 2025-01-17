package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	EthereumNodeURL string
	ServerPort      string
}

func LoadConfig() Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return Config{
		EthereumNodeURL: os.Getenv("ETHEREUM_NODE_URL"),
		ServerPort:      os.Getenv("SERVER_PORT"),
	}
}