package main

import "github.com/gin-gonic/gin"

func RunServer() *gin.Engine {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello world!!",
		})
	})
	return r
}

func main() {
	RunServer().Run(":5000")
}

