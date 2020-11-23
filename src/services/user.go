package services

import (
	"context"
	"github.com/cspinetta/go-tracing/src/base"
	"github.com/cspinetta/go-tracing/src/models"
	"github.com/cspinetta/go-tracing/src/repository"
	"go.opentelemetry.io/otel/label"
	oteltrace "go.opentelemetry.io/otel/trace"
	"strconv"
)

type IUserService interface {
	SaveUserInfo(ctx context.Context, user models.User) (models.User, error)
	GetUserInfo(ctx context.Context, id int64) (models.User, error)
	ListUser(ctx context.Context, offset int, limit int) ([]models.User, error)
}

type UserService struct {
	IUserService
	userRepository repository.IUserRepository
}

func NewUserService(userRepository repository.IUserRepository) IUserService {
	return &UserService{
		userRepository: userRepository,
	}
}

func (u *UserService) SaveUserInfo(ctx context.Context, user models.User) (models.User, error) {
	_, span := base.GlobalAppTracer.Start(ctx, "saveUser", oteltrace.WithAttributes(label.String("name", user.Name)))
	defer span.End()

	id, err := u.userRepository.Save(ctx, user)
	if err != nil {
		return models.User{}, err
	}

	return u.GetUserInfo(ctx, id)
}

func (u *UserService) GetUserInfo(ctx context.Context, id int64) (models.User, error) {
	_, span := base.GlobalAppTracer.Start(ctx, "getUserInfo", oteltrace.WithAttributes(label.String("id", strconv.FormatInt(id, 10))))
	defer span.End()

	return u.userRepository.FindById(ctx, id)
}

func (u *UserService) ListUser(ctx context.Context, offset int, limit int) ([]models.User, error) {
	_, span := base.GlobalAppTracer.Start(ctx, "getUserList", oteltrace.WithAttributes(
		label.String("offset", strconv.FormatInt(int64(offset), 10)),
		label.String("limit", strconv.FormatInt(int64(limit), 10))))
	defer span.End()

	return u.userRepository.List(ctx, offset, limit)
}
