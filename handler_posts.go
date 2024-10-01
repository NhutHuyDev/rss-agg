package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/NhutHuyDev/rss-agg/internal/db"
)

func (apiCfg *apiConfig) HandlerGetPostsForUser(w http.ResponseWriter, r *http.Request, user db.User) {
	pageNum, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		pageNum = 1
	}

	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil {
		limit = 10
	}

	offset := (pageNum - 1) * limit

	posts, err := apiCfg.DB.GetPostsForUser(r.Context(), db.GetPostsForUserParams{
		Limit:  int32(limit),
		Offset: int32(offset),
		UserID: user.ID,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("could not get posts: %v", err))
		return
	}

	respondWithJSON(w, 200, DatabasePostsToPosts(posts))
}
