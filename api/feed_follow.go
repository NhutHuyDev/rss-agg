package api

import (
	"time"

	"github.com/NhutHuyDev/rss-agg/internal/domain"
	"github.com/google/uuid"
)

type FeedFollowRes struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	FeedID    uuid.UUID `json:"feed_id"`
	UserID    uuid.UUID `json:"user_id"`
}

func CastToFeedFollow(feedFollow domain.FeedFollow) FeedFollowRes {
	return FeedFollowRes{
		ID:        feedFollow.ID,
		CreatedAt: feedFollow.CreatedAt,
		UpdatedAt: feedFollow.UpdatedAt,
		UserID:    feedFollow.UserID,
		FeedID:    feedFollow.FeedID,
	}
}

func CastToFeedFollows(feedFollows []domain.FeedFollow) []FeedFollowRes {
	arr := []FeedFollowRes{}

	for _, dbFeedFollow := range feedFollows {
		arr = append(arr, CastToFeedFollow(dbFeedFollow))
	}

	return arr
}

// DTO

type CreateFeedFollowDTO struct {
	FeedID uuid.UUID `json:"feed_id"`
}
