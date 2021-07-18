package main

import "github.com/gin-gonic/gin"

func RunServer() *gin.Engine {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello world!!",
		})
	})
	r.GET("/newapi", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "New testing api added to the application!!",
		})
	})
	return r
}

func main() {
	RunServer().Run(":5000")
}

