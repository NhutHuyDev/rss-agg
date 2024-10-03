package rest

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/NhutHuyDev/rss-agg/api"
	"github.com/NhutHuyDev/rss-agg/internal/domain"
	utils "github.com/NhutHuyDev/rss-agg/pkg"
	"github.com/google/uuid"
)

func (apiCfg *APIConfig) HandlerCreateFeed(w http.ResponseWriter, r *http.Request) {
	params, _ := r.Context().Value(DecodeBodyCxtKey).(*api.CreateFeedDTO)
	user, _ := r.Context().Value(UserCxtKey).(domain.User)

	apiCfg.FeedService.SetContext(r.Context())
	feed, err := apiCfg.FeedService.CreateFeed(domain.Feed{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
		Url:       params.Url,
	}, user.ID)

	if err != nil {
		utils.RespondWithError(w, 400, fmt.Sprintf("Could not create feed: %v", err))
		return
	}

	utils.RespondWithJSON(w, 201, api.CastToFeed(feed))
}

func (apiCfg *APIConfig) HandlerGetFeeds(w http.ResponseWriter, r *http.Request) {
	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil {
		utils.RespondWithError(w, 400, fmt.Sprintf("'limit' param must be a int number: %v", err))
		return
	}

	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		utils.RespondWithError(w, 400, fmt.Sprintf("'page' param must be a int number: %v", err))
		return
	}

	apiCfg.FeedService.SetContext(r.Context())
	feeds, err := apiCfg.FeedService.GetFeeds(limit, page)
	if err != nil {
		utils.RespondWithError(w, 400, fmt.Sprintf("Could not get feeds: %v", err))
		return
	}

	totalFeeds, err := apiCfg.FeedService.CountFeeds()
	if err != nil {
		utils.RespondWithError(w, 400, fmt.Sprintf("Could not get feeds: %v", err))
		return
	}

	totalPage, nextPage, _ := utils.GetPagination(totalFeeds, limit, page)

	type GetFeedsRes struct {
		Feeds []api.FeedRes `json:"feeds"`
		utils.PaginationRes
	}

	getFeedsRes := GetFeedsRes{
		Feeds: api.CastToFeeds(feeds),
		PaginationRes: utils.PaginationRes{
			Total:       totalFeeds,
			TotalPage:   totalPage,
			NextPage:    nextPage,
			CurrentPage: page,
		},
	}

	utils.RespondWithJSON(w, 201, getFeedsRes)
}
