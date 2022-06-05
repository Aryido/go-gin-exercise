package service

import (
	"go-gin-exercise/entity"
	"net/http"

	"github.com/gin-gonic/gin"
)

var userList = []entity.User{}

func FindAllUsers(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, userList)
}

func addUser(ctx *gin.Context) {
	user := entity.User{}
	err := ctx.BindJSON(&user)

	if err != nil {
		ctx.JSON(http.StatusNotAcceptable, "Error : "+err.Error())
		return
	}

	userList = append(userList, user)
	ctx.JSON(http.StatusOK, userList)
}
