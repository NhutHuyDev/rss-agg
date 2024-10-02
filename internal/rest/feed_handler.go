package rest

import (
	"fmt"
	"net/http"
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
	apiCfg.FeedService.SetContext(r.Context())
	feeds, err := apiCfg.FeedService.GetFeeds()
	if err != nil {
		utils.RespondWithError(w, 400, fmt.Sprintf("Could not get feeds: %v", err))
		return
	}

	utils.RespondWithJSON(w, 201, api.CastToFeeds(feeds))
}
