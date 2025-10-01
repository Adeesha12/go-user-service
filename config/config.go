// config.go
package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// var (
// 	DBHost     string
// 	DBPort     string
// 	DBUser     string
// 	DBPassword string
// 	DBName     string
// 	ServerPort string
// )

var (
	DBHost     = "localhost"
	DBPort     = "5432"
	DBUser     = ""
	DBPassword = ""
	DBName     = ""
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	DBHost = getEnv("DB_HOST", DBHost)
	DBPort = getEnv("DB_PORT", DBPort)
	DBUser = getEnv("DB_USER", DBUser)
	DBPassword = getEnv("DB_PASSWORD", DBPassword)
	DBName = getEnv("DB_NAME", DBName)

}

func getEnv(key string, defaultValue string) string {
	if value, exits := os.LookupEnv(key); exits {
		return value
	}
	return defaultValue
}
