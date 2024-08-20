package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

type Config struct {
	HTTP_PORT   string
	AUTH_PORT   string
	DB_HOST     string
	DB_PORT     string
	DB_USER     string
	DB_NAME     string
	DB_PASSWORD string
	EMAIL       string
	APP_KEY     string
}

func Load() *Config {
	if err := godotenv.Load(".env"); err != nil {
		log.Printf("error while loading .env file: %v", err)
	}

	return &Config{
		HTTP_PORT:   cast.ToString(coalesce("HTTP_PORT", "auth-service:8081")),
		AUTH_PORT:   cast.ToString(coalesce("AUTH_PORT", "auth-service:50051")),
		DB_HOST:     cast.ToString(coalesce("DB_HOST", "postgres1")),
		DB_PORT:     cast.ToString(coalesce("DB_PORT", "5432")),
		DB_USER:     cast.ToString(coalesce("DB_USER", "postgres")),
		DB_NAME:     cast.ToString(coalesce("DB_NAME", "auth_service")),
		DB_PASSWORD: cast.ToString(coalesce("DB_PASSWORD", "password")),
		EMAIL:       cast.ToString(coalesce("EMAIL", "email@mail.com")),
		APP_KEY:     cast.ToString(coalesce("APP_KEY", "key")),
	}
}

func coalesce(key string, value interface{}) interface{} {
	val, exist := os.LookupEnv(key)
	if exist {
		return val
	}
	return value
}
