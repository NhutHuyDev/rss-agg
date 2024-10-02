package routes

// import (
// 	"encoding/json"
// 	"fmt"
// 	"net/http"
// 	"time"

// 	"github.com/NhutHuyDev/rss-agg/internal/db"
// 	"github.com/google/uuid"
// )

// func (apiCfg *apiConfig) HandlerCreateFeed(w http.ResponseWriter, r *http.Request, user db.User) {
// 	type parameters struct {
// 		Name string `json:"name"`
// 		Url  string `json:"url"`
// 	}
// 	decoder := json.NewDecoder(r.Body)
// 	params := parameters{}
// 	err := decoder.Decode(&params)
// 	if err != nil {
// 		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
// 		return
// 	}

// 	feed, err := apiCfg.DB.CreateFeed(r.Context(), db.CreateFeedParams{
// 		ID:        uuid.New(),
// 		CreatedAt: time.Now().UTC(),
// 		UpdatedAt: time.Now().UTC(),
// 		Name:      params.Name,
// 		Url:       params.Url,
// 		UserID:    user.ID,
// 	})
// 	if err != nil {
// 		respondWithError(w, 400, fmt.Sprintf("Could not create feed: %v", err))
// 		return
// 	}

// 	respondWithJSON(w, 201, DatabaseFeedToFeed(feed))
// }

// func (apiCfg *apiConfig) HandlerGetFeeds(w http.ResponseWriter, r *http.Request) {

// 	feeds, err := apiCfg.DB.GetFeeds(r.Context())
// 	if err != nil {
// 		respondWithError(w, 400, fmt.Sprintf("Could not get feeds: %v", err))
// 		return
// 	}

// 	respondWithJSON(w, 201, DatabaseFeedsToFeeds(feeds))
// }
