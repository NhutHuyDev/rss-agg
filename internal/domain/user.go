package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string
	ApiKey    string
}

type IUserService interface {
	SetContext(ctx context.Context)
	GetUser(user User) (User, error)
	CreateUser(user User) (User, error)
}
