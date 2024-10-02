package api

import (
	"time"

	"github.com/NhutHuyDev/rss-agg/internal/domain"
	"github.com/google/uuid"
)

type UserRes struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	ApiKey    string    `json:"api_key"`
}

func CastToUser(user domain.User) UserRes {
	return UserRes{
		ID:        user.ID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Name:      user.Name,
		ApiKey:    user.ApiKey,
	}
}

//DTO

type CreateUserDTO struct {
	Name string `json:"name" validate:"required,min=3,max=32"`
}
