package data_base

import (
	"fmt"
	"go-user-service/config"
	"go-user-service/utils/logger"

	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	ID        uint      `gorm:"primaryKey"`
	Name      string    `gorm:"size:100;not null"`
	Email     string    `gorm:"size:100;uniqueIndex;not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

type Database struct {
	DB *gorm.DB
}

func ConnectDb(cfg *config.Config) *Database {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Error.Printf("🔴 Failed to connect to database:%v", err)
	}
	err2 := db.Exec("SELECT 1").Error
	if err2 != nil {
		logger.Error.Printf("🔴 database connection verification failed: %v", err2)

	}
	logger.Info.Println("🟢 Database connected successfully")
	return &Database{DB: db}
}

func (d *Database) Close() {
	if sqlDB, err := d.DB.DB(); err == nil {
		sqlDB.Close()
	}
}
