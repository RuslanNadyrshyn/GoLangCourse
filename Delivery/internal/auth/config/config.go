package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

const (
	accessLifetimeMinutes  = 5
	refreshLifetimeMinutes = 60
)

type Config struct {
	Port                   string
	AccessSecret           string
	RefreshSecret          string
	AccessLifetimeMinutes  int
	RefreshLifetimeMinutes int
}

func NewConfig(forTesting bool) *Config {
	envName := "cfg/.env"
	if forTesting {
		envName = "./../../.env.testing"
	}

	err := godotenv.Load(envName)
	if err != nil {
		log.Println(err)
		log.Println("Error loading .env file")
	}

	return &Config{
		Port:                   os.Getenv("PORT"),
		AccessSecret:           os.Getenv("ACCESS_SECRET"),
		RefreshSecret:          os.Getenv("REFRESH_SECRET"),
		AccessLifetimeMinutes:  accessLifetimeMinutes,
		RefreshLifetimeMinutes: refreshLifetimeMinutes,
	}
}
