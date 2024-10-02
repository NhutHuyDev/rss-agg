package services

import (
	"context"
	"database/sql"
	"time"

	"github.com/NhutHuyDev/rss-agg/internal/domain"
	"github.com/NhutHuyDev/rss-agg/internal/infra/db"
	"github.com/google/uuid"
)

type PostServiceImpl struct {
	Queries *db.Queries
	Ctx     context.Context
}

func CastToPost(dbPost db.Post) domain.Post {
	var description *string
	if dbPost.Description.Valid {
		description = &dbPost.Description.String
	}

	return domain.Post{
		ID:          dbPost.ID,
		CreatedAt:   dbPost.CreatedAt,
		UpdatedAt:   dbPost.UpdatedAt,
		Title:       dbPost.Title,
		Description: description,
		PublishedAt: dbPost.PublishedAt,
		Url:         dbPost.Url,
		FeedID:      dbPost.FeedID,
	}
}

func CastToPosts(dbPosts []db.Post) []domain.Post {
	posts := []domain.Post{}

	for _, dbPost := range dbPosts {
		posts = append(posts, CastToPost(dbPost))
	}

	return posts
}

// Implementation

func (postService *PostServiceImpl) GetPostsByUser(user_id uuid.UUID, page int, limit int) ([]domain.Post, error) {
	offset := (page - 1) * limit

	posts, err := postService.Queries.GetPostsByUser(postService.Ctx, db.GetPostsByUserParams{
		Limit:  int32(limit),
		Offset: int32(offset),
		UserID: user_id,
	})
	if err != nil {
		return []domain.Post{}, err
	}

	return CastToPosts(posts), nil
}

func (postService *PostServiceImpl) CreatePost(post domain.Post) (domain.Post, error) {
	description := sql.NullString{}
	if *post.Description != "" {
		description.String = *post.Description
		description.Valid = true
	}

	result, err := postService.Queries.CreatePost(postService.Ctx,
		db.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now().UTC(),
			UpdatedAt:   time.Now().UTC(),
			Title:       post.Title,
			Description: description,
			Url:         post.Url,
			PublishedAt: post.PublishedAt,
			FeedID:      post.ID,
		})

	if err != nil {
		return domain.Post{}, err
	}

	return CastToPost(result), nil
}
