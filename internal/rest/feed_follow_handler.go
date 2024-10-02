package rest

import (
	"fmt"
	"net/http"

	"github.com/NhutHuyDev/rss-agg/api"
	"github.com/NhutHuyDev/rss-agg/internal/domain"
	utils "github.com/NhutHuyDev/rss-agg/pkg"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func (apiCfg *APIConfig) HandlerCreateFeedFollow(w http.ResponseWriter, r *http.Request) {
	params, _ := r.Context().Value(DecodeBodyCxtKey).(*api.CreateFeedFollowDTO)
	user, _ := r.Context().Value(UserCxtKey).(domain.User)

	apiCfg.FeedFollowService.SetContext(r.Context())
	feedFollow, err := apiCfg.FeedFollowService.CreateFeedFollow(params.FeedID, user.ID)
	if err != nil {
		utils.RespondWithError(w, 400, fmt.Sprintf("could not follow feed: %v", err))
		return
	}

	utils.RespondWithJSON(w, 201, api.CastToFeedFollow(feedFollow))
}

func (apiCfg *APIConfig) HandlerGetFeedFollows(w http.ResponseWriter, r *http.Request) {
	user, _ := r.Context().Value(UserCxtKey).(domain.User)

	apiCfg.FeedFollowService.SetContext(r.Context())
	feedFollows, err := apiCfg.FeedFollowService.GetFeedFollows(user.ID)
	if err != nil {
		utils.RespondWithError(w, 400, fmt.Sprintf("could not get feed follows: %v", err))
	}

	utils.RespondWithJSON(w, 201, api.CastToFeedFollows(feedFollows))
}

func (apiCfg *APIConfig) HandlerDeleteFeedFollows(w http.ResponseWriter, r *http.Request) {
	feedFollowIDStr := chi.URLParam(r, "feed_folow_id")
	feedFollowID, err := uuid.Parse(feedFollowIDStr)
	if err != nil {
		utils.RespondWithError(w, 400, fmt.Sprintf("could not parse feed_follow_id param: %v", err))
	}

	user, _ := r.Context().Value(UserCxtKey).(domain.User)

	apiCfg.FeedFollowService.SetContext(r.Context())
	err = apiCfg.FeedFollowService.DeleteFeedFollow(feedFollowID, user.ID)
	if err != nil {
		utils.RespondWithError(w, 400, fmt.Sprintf("could not delete feed follow: %v", err))
	}

	utils.RespondWithJSON(w, 200, struct{}{})
}
