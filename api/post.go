package api

import (
	"time"

	"github.com/NhutHuyDev/rss-agg/internal/domain"
	"github.com/google/uuid"
)

type PostRes struct {
	ID          uuid.UUID `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Title       string    `json:"title"`
	Description *string   `json:"description"`
	PublishedAt time.Time `json:"published_at"`
	Url         string    `json:"url"`
	FeedID      uuid.UUID `json:"feed_id"`
}

func CastToPost(post domain.Post) PostRes {
	return PostRes{
		ID:          post.ID,
		CreatedAt:   post.CreatedAt,
		UpdatedAt:   post.UpdatedAt,
		Title:       post.Title,
		Description: post.Description,
		PublishedAt: post.PublishedAt,
		Url:         post.Url,
		FeedID:      post.FeedID,
	}
}

func CastToPosts(posts []domain.Post) []PostRes {
	arr := []PostRes{}

	for _, post := range posts {
		arr = append(arr, CastToPost(post))
	}

	return arr
}

// DTO
