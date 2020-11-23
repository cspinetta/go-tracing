package repository

import (
	"context"
	"errors"
	"github.com/cspinetta/go-tracing/src/models"
	"github.com/cspinetta/go-tracing/src/utils"
	"github.com/jmoiron/sqlx"
)

type IUserRepository interface {
	Save(ctx context.Context, user models.User) (int64, error)
	List(ctx context.Context, offset int, limit int) ([]models.User, error)
	FindById(ctx context.Context, id int64) (models.User, error)
}

type UserRepository struct {
	IUserRepository
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) IUserRepository {
	return &UserRepository{db: db}
}

func (u *UserRepository) Save(ctx context.Context, user models.User) (int64, error) {
	now := utils.TimeNow()
	user.CreatedAt = &now
	result, err := u.db.NamedExecContext(ctx, `
		INSERT INTO 
			user (
				name,
				birthday,
				created_at,
				updated_at
			)
		VALUES (
			:name,
			:birthday,
			:created_at,
			:updated_at
		)
	`, user)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (u *UserRepository) FindById(ctx context.Context, id int64) (models.User, error) {
	var users []models.User

	err := u.db.SelectContext(ctx, &users, `
select
	id,
	name,
	birthday,
	created_at,
	updated_at
from
	user
where
	id = ?
`, id)

	if err != nil {
		return models.User{}, err
	}

	if len(users) < 1 {
		return models.User{}, errors.New("user not found")
	}

	return users[0], nil
}

func (u *UserRepository) List(ctx context.Context, offset int, limit int) ([]models.User, error) {
	var users []models.User

	err := u.db.SelectContext(ctx, &users, `
SELECT
	id,
	name,
	birthday,
	created_at,
	updated_at
FROM
	user
LIMIT ?, ?
`, offset, limit)

	if err != nil {
		return []models.User{}, err
	}

	return users, nil
}
