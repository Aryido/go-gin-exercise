package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("gin-rest", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "hello",
		})
	})

	router.Run(":8080")
}
