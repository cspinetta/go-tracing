package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type HealthCheckController struct {
}

func NewHealthCheckHandler() *HealthCheckController {
	return &HealthCheckController{}
}

func (ctrl *HealthCheckController) HandlePing(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}
