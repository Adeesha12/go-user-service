package data_base

import (
	"fmt"
	"log"
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

func db() {
	dsn := "user:password@tcp(localhost:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// Auto migrate
	err = db.AutoMigrate(&User{})
	if err != nil {
		log.Fatal(err)
	}

	// CRUD operations with GORM
	user := User{Name: "John Doe", Email: "john@example.com"}

	// Create
	result := db.Create(&user)
	if result.Error != nil {
		log.Fatal(result.Error)
	}

	// Read
	var fetchedUser User
	db.First(&fetchedUser, user.ID)
	fmt.Printf("User: %+v\n", fetchedUser)

	// Update
	db.Model(&fetchedUser).Update("Name", "Jane Doe")

	// Delete
	db.Delete(&fetchedUser)
}
