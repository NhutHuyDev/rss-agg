package api

import (
	"time"

	"github.com/google/uuid"
)

type CreateUserReq struct {
	Name string `json:"name"`
}

type CreateUserRes struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	ApiKey    string    `json:"api_key"`
}
