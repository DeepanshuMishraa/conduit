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
	INSTANCE     string
}

func Load() (*Config, error) {
	_ = godotenv.Load()

	DATABASE_URL := os.Getenv("DATABASE_URL")
	PORT := os.Getenv("PORT")
	INSTANCE := os.Getenv("INSTANCE")

	if DATABASE_URL == "" || PORT == "" || INSTANCE == "" {
		return nil, errors.New("Environment variables not set")
	}

	log.Println("All env vars loaded")

	return &Config{
		DATABASE_URL: DATABASE_URL,
		PORT:         PORT,
		INSTANCE:     INSTANCE,
	}, nil
}
