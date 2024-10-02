package domain

import (
	"time"

	"github.com/google/uuid"
)

type FeedFollow struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	FeedID    uuid.UUID
	UserID    uuid.UUID
}

type IFeedFollowService interface {
	GetFeedFollows(user_id uuid.UUID) ([]FeedFollow, error)
	CreateFeedFollow(feed_id uuid.UUID, user_id uuid.UUID) (FeedFollow, error)
	DeleteFeedFollow(id uuid.UUID, user_id uuid.UUID) error
}
