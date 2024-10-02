package rest

import (
	"github.com/NhutHuyDev/rss-agg/internal/domain"
	"github.com/NhutHuyDev/rss-agg/internal/infra/db"
)

type APIConfig struct {
	DB          *db.Queries
	UserService domain.IUserService
}
