package controllers

import (
	"github.com/cspinetta/go-tracing/src/models"
	"github.com/cspinetta/go-tracing/src/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type IUserController interface {
	SaveUserInfo(ctx *gin.Context)
	GetUserInfo(ctx *gin.Context)
	ListUser(ctx *gin.Context)
}

type UserController struct {
	IUserController
	UserService services.IUserService
}

func NewUserController(userService services.IUserService) IUserController {
	return &UserController{
		UserService: userService,
	}
}

func (u *UserController) SaveUserInfo(ctx *gin.Context) {
	var req models.User

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.ValidationError{Message: "Request validation error. Reason: " + err.Error()})
		return
	}
	user, err := u.UserService.SaveUserInfo(ctx.Request.Context(), req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.ValidationError{Message: "Unexpected error. Reason: " + err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func (u *UserController) GetUserInfo(ctx *gin.Context) {
	userIdText := ctx.Param("user-id")

	if userIdText == "" {
		ctx.JSON(http.StatusBadRequest, models.ValidationError{Message: "Request validation error. Reason: user-id is not a valid value"})
		return
	}

	userId, err := strconv.Atoi(userIdText)
	if err != err {
		ctx.JSON(http.StatusBadRequest, models.ValidationError{Message: "Request validation error. Reason: user-id is not a valid value"})
		return
	}

	user, err := u.UserService.GetUserInfo(ctx.Request.Context(), int64(userId))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.ValidationError{Message: "Unexpected error. Reason: " + err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func (u *UserController) ListUser(ctx *gin.Context) {
	var req models.ListUserRequest

	err := ctx.ShouldBindQuery(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.ValidationError{Message: "Request validation error. Reason: " + err.Error()})
		return
	}
	users, err := u.UserService.ListUser(ctx.Request.Context(), req.Offset, req.Limit)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.ValidationError{Message: "Unexpected error. Reason: " + err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, users)
}
