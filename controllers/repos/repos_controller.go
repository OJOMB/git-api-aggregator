package controllers

import (
	"net/http"

	"github.com/OJOMB/git-api-aggregator/domain/gitrepositories"
	"github.com/OJOMB/git-api-aggregator/services"
	"github.com/OJOMB/git-api-aggregator/utils"
	"github.com/gin-gonic/gin"
)

// Repos is the exported entry point for the Repos controller
var Repos ReposControllerInterface

func init() {
	Repos = &repos{}
}

// ReposControllerInterface is the interface for the Repos controller
type ReposControllerInterface interface {
	CreateRepo(ctx *gin.Context)
}

type repos struct{}

// CreateRepo controller
func (r *repos) CreateRepo(ctx *gin.Context) {
	var request gitrepositories.CreateRepoRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(400, utils.NewBadRequestError("Received invalid request"))
		return
	}

	result, err := services.Repos.CreateRepo(&request)
	if err != nil {
		ctx.JSON(err.GetStatus(), err)
		return
	}

	ctx.JSON(http.StatusCreated, result)
}
