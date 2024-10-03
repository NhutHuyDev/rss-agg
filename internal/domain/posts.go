package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type Post struct {
	ID          uuid.UUID
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Title       string
	Description *string
	PublishedAt time.Time
	Url         string
	FeedID      uuid.UUID
}

type IPostService interface {
	SetContext(ctx context.Context)
	GetPostsByUser(user_id uuid.UUID, page int, limit int) ([]Post, error)
	CountPosts(user_id uuid.UUID) (int, error)
	CreatePost(post Post) (Post, error)
}
