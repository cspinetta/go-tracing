package models

type HealthCheckResponse struct {
	Status         string `json:"status"`
	DbConnectionOk bool   `json:"db_conn_ok"`
}
