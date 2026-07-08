package config

import (
	"errors"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DATABASE_URL string
	PORT         string
}

func Load() (*Config, error) {
	err := godotenv.Load()

	if err != nil {
		return nil, err
	}

	DATABASE_URL := os.Getenv("DATABASE_URL")
	PORT := os.Getenv("PORT")

	if DATABASE_URL == "" || PORT == "" {
		return nil, errors.New("Environment variables not set")
	}

	log.Println("All env vars loaded")

	return &Config{
		DATABASE_URL: DATABASE_URL,
		PORT:         PORT,
	}, nil
}
