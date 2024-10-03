package domain

import (
	"context"
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
	SetContext(ctx context.Context)
	GetFeeds() ([]Feed, error)
	CountFeeds() (int, error)
	CreateFeed(feed Feed, user_id uuid.UUID) (Feed, error)
	GetNextFeedsToFetch(limit int) ([]Feed, error)
	MarkFeedAsFetched(id uuid.UUID) (Feed, error)
}
