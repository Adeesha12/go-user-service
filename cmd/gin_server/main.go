package main

import (
	"github.com/gin-gonic/gin"
)

// func get()  {

// }
func main() {
	router := gin.Default()
	db()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.Run() // listens on 0.0.0.0:8080 by default
}
