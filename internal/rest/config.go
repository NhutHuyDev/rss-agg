package rest

import (
	"github.com/NhutHuyDev/rss-agg/internal/domain"
	"github.com/NhutHuyDev/rss-agg/internal/infra/db"
	"github.com/go-playground/validator/v10"
)

type APIConfig struct {
	DB                *db.Queries
	Validate          *validator.Validate
	UserService       domain.IUserService
	FeedService       domain.IFeedService
	FeedFollowService domain.IFeedFollowService
}
