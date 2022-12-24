package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		// If you set TrustedPlatform, ClientIP() will resolve the
		// corresponding header and return IP directly
		c.JSON(200, gin.H{"msg": "ClientIP: " + c.ClientIP()})
	})

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	fmt.Printf("ClientIP: %s\n", "http://localhost:8080")

	r.Run("localhost:8080") // listen and serve on 0.0.0.0:8080
}
