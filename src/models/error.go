package models

type ValidationError struct {
	Message string `json:"message" binding:"required"`
}
