package tests

import (
	"fmt"
	"testing"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"go-user-service/config"
	"go-user-service/utils/logger"
)

func TestDb(t *testing.T) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		config.DBHost, config.DBUser, config.DBPassword, config.DBName, config.DBPort)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Error.Printf("failed to connect database: %v", err)
		// log.Fatalf("failed to connect database: %v", err)
	}
	err2 := db.Exec("SELECT 1").Error
	if err2 != nil {
		logger.Error.Printf("database connection verification failed: %v", err2)
		// log.Fatalf("database connection verification failed: %v", err2)
	}
	logger.Info.Println("Success ")
	// ... rest of your application logic
}
