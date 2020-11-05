package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Health is the health entrypoint
var Health HealthControllerInterface

const polo = "polo"

func init() {
	Health = &health{}
}

// HealthControllerInterface is the health controller interface
type HealthControllerInterface interface {
	Polo(ctx *gin.Context)
}

type health struct{}

// Polo is a health check
func (h *health) Polo(ctx *gin.Context) {
	ctx.String(http.StatusOK, polo)
}
