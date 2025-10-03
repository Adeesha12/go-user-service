package config

import (
	"go-user-service/utils/logger"

	"os"
	"path/filepath"
	"sync"

	"github.com/joho/godotenv"
)

type Config struct {
	DBUser     string
	DBPassword string
	DBName     string
	DBHost     string
	DBPort     string
	AppEnv     string
}

var (
	cfg  *Config
	once sync.Once
)

func LoadConfig() *Config {
	once.Do(func() {

		tryLoadEnv() // Load .env file only once

		cfg = &Config{
			DBUser:     os.Getenv("DB_USER"),
			DBPassword: os.Getenv("DB_PASSWORD"),
			DBName:     os.Getenv("DB_NAME"),
			DBHost:     os.Getenv("DB_HOST"),
			DBPort:     os.Getenv("DB_PORT"),
			AppEnv:     os.Getenv("APP_ENV"), // dev, test, prod
		}
	})

	return cfg
}

func tryLoadEnv() {
	// Try current directory
	if err := godotenv.Load(".env"); err == nil {
		logger.Info.Println("🟢 .env loaded from current directory")
		return
	}

	// Try project root (one level up from tests)
	if err := godotenv.Load("../.env"); err == nil {
		logger.Info.Println("🟢 .env loaded from project root")
		return
	}

	// Try two levels up (if running from subdirectories)
	if err := godotenv.Load("../../.env"); err == nil {
		logger.Info.Println("🟢 .env loaded from project root (two levels up)")
		return
	}

	// Get absolute path
	absPath, _ := filepath.Abs(".")
	logger.Warn.Printf("🟡 Current working directory: %s", absPath)
	logger.Error.Println("🔴  No .env file found, relying on system environment variables")
}
