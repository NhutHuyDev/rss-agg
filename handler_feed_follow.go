package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/NhutHuyDev/rss-agg/internal/db"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) HandlerCreateFeedFollow(w http.ResponseWriter, r *http.Request, user db.User) {
	type parameters struct {
		FeedID uuid.UUID `json:"feed_id"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	feedFollow, err := apiCfg.DB.CreateFeedFollow(r.Context(), db.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    params.FeedID,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Could not follow feed: %v", err))
		return
	}

	respondWithJSON(w, 201, DatabaseFeedFollowToFeedFollow(feedFollow))
}

func (apiCfg *apiConfig) HandlerGetFeedFollows(w http.ResponseWriter, r *http.Request, user db.User) {
	feedFollows, err := apiCfg.DB.GetFeedFollows(r.Context(), user.ID)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Could not get feed follows: %v", err))
	}

	respondWithJSON(w, 201, DatabaseFeedFollowsToFeedFollows(feedFollows))
}

func (apiCfg *apiConfig) HandlerDeleteFeedFollows(w http.ResponseWriter, r *http.Request, user db.User) {
	feedFollowIDStr := chi.URLParam(r, "feed_folow_id")
	feedFollowID, err := uuid.Parse(feedFollowIDStr)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Could not parse feed_follow_id param: %v", err))
	}

	err = apiCfg.DB.DeleteFeedFollows(r.Context(), db.DeleteFeedFollowsParams{
		ID:     feedFollowID,
		UserID: user.ID,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Could not delete feed follow: %v", err))
	}

	respondWithJSON(w, 200, struct{}{})
}
