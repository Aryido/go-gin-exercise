package main

import (
	"go-gin-exercise/service"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	v1 := router.Group("v1")
	service.AddUserRouter(v1)
	router.Run(":8080")
}
