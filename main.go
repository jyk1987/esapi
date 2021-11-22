package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jyk1987/es/log"
)

func main() {
	log.Log.Debug("test")
	log.Log.Debug(gin.Default)
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
