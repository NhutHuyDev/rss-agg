package domain

import (
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
	GetPostsByUser(user_id uuid.UUID, page int, limit int) ([]Post, error)
	CreatePost(post Post) (Post, error)
}
