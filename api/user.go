package api

import (
	"time"

	"github.com/google/uuid"
)

type CreateUserDTO struct {
	Name string `json:"name" validate:"required,min=3,max=32"`
}

type UserRes struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	ApiKey    string    `json:"api_key"`
}
