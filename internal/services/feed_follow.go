package services

import (
	"context"
	"time"

	"github.com/NhutHuyDev/rss-agg/internal/domain"
	"github.com/NhutHuyDev/rss-agg/internal/infra/db"
	"github.com/google/uuid"
)

type FeedFollowServiceImpl struct {
	Queries *db.Queries
	Ctx     context.Context
}

func CastToFeedFollow(feedFollow db.FeedFollow) domain.FeedFollow {
	return domain.FeedFollow{
		ID:        feedFollow.ID,
		CreatedAt: feedFollow.CreatedAt,
		UpdatedAt: feedFollow.UpdatedAt,
		UserID:    feedFollow.UserID,
		FeedID:    feedFollow.FeedID,
	}
}

func CastToFeedFollows(dbFeedFollows []db.FeedFollow) []domain.FeedFollow {
	feedFollows := []domain.FeedFollow{}

	for _, dbFeedFollow := range dbFeedFollows {
		feedFollows = append(feedFollows, CastToFeedFollow(dbFeedFollow))
	}

	return feedFollows
}

// Implementation

func (feedfollowService *FeedFollowServiceImpl) SetContext(ctx context.Context) {
	feedfollowService.Ctx = ctx
}

func (feedfollowService *FeedFollowServiceImpl) GetFeedFollows(user_id uuid.UUID) ([]domain.FeedFollow, error) {
	result, err := feedfollowService.Queries.GetFeedFollows(feedfollowService.Ctx, user_id)
	if err != nil {
		return []domain.FeedFollow{}, err
	}

	return CastToFeedFollows(result), nil
}

func (feedfollowService *FeedFollowServiceImpl) CreateFeedFollow(feed_id uuid.UUID, user_id uuid.UUID) (domain.FeedFollow, error) {
	result, err := feedfollowService.Queries.CreateFeedFollow(feedfollowService.Ctx, db.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		FeedID:    feed_id,
		UserID:    user_id,
	})

	if err != nil {
		return domain.FeedFollow{}, err
	}

	return CastToFeedFollow(result), nil
}

func (feedfollowService *FeedFollowServiceImpl) DeleteFeedFollow(id uuid.UUID, user_id uuid.UUID) error {
	err := feedfollowService.Queries.DeleteFeedFollow(feedfollowService.Ctx, db.DeleteFeedFollowParams{
		ID:     id,
		UserID: user_id,
	})

	if err != nil {
		return err
	}

	return nil
}
