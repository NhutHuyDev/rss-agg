package domain

import (
	"time"

	"github.com/google/uuid"
)

type Feed struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string
	Url       string
}

type IFeedService interface {
	GetFeeds() ([]Feed, error)
	CreateFeed(feed Feed, user_id uuid.UUID) (Feed, error)
}
