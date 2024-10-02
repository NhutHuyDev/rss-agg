package api

import (
	"time"

	"github.com/NhutHuyDev/rss-agg/internal/domain"
	"github.com/google/uuid"
)

type FeedRes struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	Url       string    `json:"url"`
}

func CastToFeed(feed domain.Feed) FeedRes {
	return FeedRes{
		ID:        feed.ID,
		CreatedAt: feed.CreatedAt,
		UpdatedAt: feed.UpdatedAt,
		Name:      feed.Name,
		Url:       feed.Url,
	}
}

func CastToFeeds(feeds []domain.Feed) []FeedRes {
	arr := []FeedRes{}

	for _, feed := range feeds {
		arr = append(arr, CastToFeed(feed))
	}

	return arr
}

// DTO

type CreateFeedDTO struct {
	Name string `json:"name" validate:"required,min=3,max=32"`
	Url  string `json:"url" validate:"required,min=3"`
}
