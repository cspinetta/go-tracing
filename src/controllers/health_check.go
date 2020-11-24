package controllers

import (
	"github.com/cspinetta/go-tracing/src/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type HealthCheckController struct {
	healthCheckService services.IHealthCheckService
}

func NewHealthCheckHandler(healthCheckService services.IHealthCheckService) *HealthCheckController {
	return &HealthCheckController{
		healthCheckService: healthCheckService,
	}
}

func (ctrl *HealthCheckController) HandleHealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, ctrl.healthCheckService.VerifyStatus(c.Request.Context()))
}
