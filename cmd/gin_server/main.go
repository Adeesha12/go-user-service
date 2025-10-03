package main

import (
	"go-user-service/data_base"

	"go-user-service/config"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.LoadConfig()

	router := gin.Default()
	db := data_base.ConnectDb(cfg)
	defer db.Close()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.Run() // listens on 0.0.0.0:8080 by default
}
