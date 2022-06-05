package service

import (
	"github.com/gin-gonic/gin"
)

func AddUserRouter(rg *gin.RouterGroup) {
	user := rg.Group("/users")
	user.GET("/", FindAllUsers)
	user.POST("/", addUser)
}
