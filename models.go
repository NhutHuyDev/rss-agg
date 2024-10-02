package main

// import (
// 	"time"

// 	"github.com/NhutHuyDev/rss-agg/internal/db"
// 	"github.com/google/uuid"
// )

// type Post struct {
// 	ID          uuid.UUID `json:"id"`
// 	CreatedAt   time.Time `json:"created_at"`
// 	UpdatedAt   time.Time `json:"updated_at"`
// 	Title       string    `json:"title"`
// 	Description *string   `json:"description"`
// 	PublishedAt time.Time `json:"published_at"`
// 	Url         string    `json:"url"`
// 	FeedID      uuid.UUID `json:"feed_id"`
// }

// func DatabasePostToPost(post db.Post) Post {
// 	var description *string
// 	if post.Description.Valid {
// 		description = &post.Description.String
// 	}

// 	return Post{
// 		ID:          post.ID,
// 		CreatedAt:   post.CreatedAt,
// 		UpdatedAt:   post.UpdatedAt,
// 		Title:       post.Title,
// 		Description: description,
// 		PublishedAt: post.PublishedAt,
// 		Url:         post.Url,
// 		FeedID:      post.FeedID,
// 	}
// }

// func DatabasePostsToPosts(dbPosts []db.Post) []Post {
// 	posts := []Post{}

// 	for _, dbPost := range dbPosts {
// 		posts = append(posts, DatabasePostToPost(dbPost))
// 	}

// 	return posts
// }
