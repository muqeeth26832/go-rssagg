package main

import (
	"log"

	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
)

type Env struct {
	PORT string `env:"PORT,required"`
	DB_URL string `env:"DB_URL,required"`
}

func EnvConfig() *Env {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %e", err)
	}

	envConfig := &Env{}

	if err := env.Parse(envConfig); err != nil {
		log.Fatalf("Error parsing environment variables: %e", err)
	}

	return envConfig
}
