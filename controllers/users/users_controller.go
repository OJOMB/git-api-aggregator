package controllers

import (
	"net/http"

	"github.com/OJOMB/git-api-aggregator/domain/gitusers"
	"github.com/OJOMB/git-api-aggregator/services"
	"github.com/OJOMB/git-api-aggregator/utils"
	"github.com/gin-gonic/gin"
)

// Users is the exported entry point for the Repos controller
var Users UsersControllerInterface

func init() {
	Users = &users{}
}

// UsersControllerInterface is the interface for the Repos controller
type UsersControllerInterface interface {
	GetUser(ctx *gin.Context)
}

type users struct{}

func (users) GetUser(ctx *gin.Context) {
	var getUserRequest gitusers.GetUserRequest
	if err := ctx.ShouldBindJSON(&getUserRequest); err != nil {
		ctx.JSON(400, utils.NewBadRequestError("Received invalid Get User request"))
		return
	}

	result, err := services.Users.GetUser(&getUserRequest)
	if err != nil {
		ctx.JSON(err.GetStatus(), err)
		return
	}

	ctx.JSON(http.StatusCreated, result)
}
