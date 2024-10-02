package services

import (
	"context"
	"time"

	"github.com/NhutHuyDev/rss-agg/internal/domain"
	"github.com/NhutHuyDev/rss-agg/internal/infra/db"
	"github.com/google/uuid"
)

type UserServiceImpl struct {
	Queries *db.Queries
	Ctx     context.Context
}

func CastToUser(dbUser db.User) domain.User {
	return domain.User{
		ID:        dbUser.ID,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
		Name:      dbUser.Name,
		ApiKey:    dbUser.ApiKey,
	}
}

// Implementation
func (userService *UserServiceImpl) SetContext(ctx context.Context) {
	userService.Ctx = ctx
}

func (userService *UserServiceImpl) CreateUser(user domain.User) (domain.User, error) {
	result, err := userService.Queries.CreateUser(userService.Ctx, db.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      user.Name,
	})

	if err != nil {
		return domain.User{}, err
	}

	return CastToUser(result), nil
}
