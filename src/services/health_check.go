package services

import (
	"context"
	"github.com/cspinetta/go-tracing/src/models"
	"github.com/jmoiron/sqlx"
)

type IHealthCheckService interface {
	VerifyStatus(ctx context.Context) models.HealthCheckResponse
}

type HealthCheckService struct {
	IHealthCheckService
	db *sqlx.DB
}

func NewHealthCheck(db *sqlx.DB) IHealthCheckService {
	return &HealthCheckService{
		db: db,
	}
}

func (h *HealthCheckService) VerifyStatus(ctx context.Context) models.HealthCheckResponse {
	err := h.db.PingContext(ctx)
	if err != nil {
		return models.HealthCheckResponse{
			Status:         "unhealth",
			DbConnectionOk: false,
		}
	}
	return models.HealthCheckResponse{
		Status:         "health",
		DbConnectionOk: true,
	}
}
