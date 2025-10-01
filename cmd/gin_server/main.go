package main

import (
	"go-user-service/data_base"

	"github.com/gin-gonic/gin"
)

// func get()  {

// }
func main() {
	router := gin.Default()
	data_base.ConnectDb()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.Run() // listens on 0.0.0.0:8080 by default
}
