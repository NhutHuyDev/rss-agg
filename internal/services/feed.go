package services

import (
	"context"
	"errors"
	"time"

	"github.com/NhutHuyDev/rss-agg/internal/domain"
	"github.com/NhutHuyDev/rss-agg/internal/infra/db"
	"github.com/google/uuid"
)

type FeedServiceImpl struct {
	Queries *db.Queries
	Ctx     context.Context
}

func CastToFeed(dbFeed db.Feed) domain.Feed {
	return domain.Feed{
		ID:        dbFeed.ID,
		CreatedAt: dbFeed.CreatedAt,
		UpdatedAt: dbFeed.UpdatedAt,
		Name:      dbFeed.Name,
		Url:       dbFeed.Url,
	}
}

func CastToFeeds(dbFeeds []db.Feed) []domain.Feed {
	feeds := []domain.Feed{}

	for _, dbFeed := range dbFeeds {
		feeds = append(feeds, CastToFeed(dbFeed))
	}

	return feeds
}

// Implementation

func (feedService *FeedServiceImpl) SetContext(ctx context.Context) {
	feedService.Ctx = ctx
}

func (feedService *FeedServiceImpl) GetFeeds(limit int, page int) ([]domain.Feed, error) {
	if page <= 0 {
		return []domain.Feed{}, errors.New("'page' params is higher than 0")
	}

	if limit <= 0 {
		return []domain.Feed{}, errors.New("'limit' params is higher than 0")
	}

	offset := (page - 1) * limit

	result, err := feedService.Queries.GetFeeds(feedService.Ctx, db.GetFeedsParams{
		Limit:  int32(limit),
		Offset: int32(offset),
	})
	if err != nil {
		return []domain.Feed{}, err
	}

	return CastToFeeds(result), nil
}

func (feedService *FeedServiceImpl) CountFeeds() (int, error) {
	result, err := feedService.Queries.CountFeeds(feedService.Ctx)
	if err != nil {
		return 0, err
	}

	return int(result), nil
}

func (feedService *FeedServiceImpl) CreateFeed(feed domain.Feed, user_id uuid.UUID) (domain.Feed, error) {
	result, err := feedService.Queries.CreateFeed(feedService.Ctx, db.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      feed.Name,
		Url:       feed.Url,
		UserID:    user_id,
	})

	if err != nil {
		return domain.Feed{}, err
	}

	return CastToFeed(result), nil
}

func (feedService *FeedServiceImpl) GetNextFeedsToFetch(limit int) ([]domain.Feed, error) {
	result, err := feedService.Queries.GetNextFeedsToFetch(feedService.Ctx, int32(limit))
	if err != nil {
		return []domain.Feed{}, err
	}

	return CastToFeeds(result), nil
}

func (feedService *FeedServiceImpl) MarkFeedAsFetched(id uuid.UUID) (domain.Feed, error) {
	result, err := feedService.Queries.MarkFeedAsFetched(feedService.Ctx, id)
	if err != nil {
		return domain.Feed{}, err
	}

	return CastToFeed(result), nil
}
