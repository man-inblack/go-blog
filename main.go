package main

import "github.com/gin-gonic/gin"

func main() {

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})
	r.GET("/", index)
	r.GET("/blog/:id", getTextByID)
	r.GET("/blog", getTextList)

	r.Run(":8086")
}
