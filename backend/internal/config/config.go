package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	RedisURL  string
	JWTSecret string
	APIURL    string
	Port      string
}

func Load() *Config {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Println("No se encontró .env, usando variables del sistema")
	}

	return &Config{
		RedisURL:  os.Getenv("REDIS_URL"),
		JWTSecret: os.Getenv("JWT_SECRET"),
		APIURL:    os.Getenv("API_URL"),
		Port:      os.Getenv("PORT"),
	}
}
