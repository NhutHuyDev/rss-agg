package main

import (
	"fmt"
	"net/http"

	"github.com/NhutHuyDev/rss-agg/auth"
	"github.com/NhutHuyDev/rss-agg/internal/db"
)

type authedHandler func(http.ResponseWriter, *http.Request, db.User)

func (apiCfg apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetApiKey(r.Header)
		if err != nil {
			respondWithError(w, 403, fmt.Sprintf("Auth error: %v", err))
			return
		}

		user, err := apiCfg.DB.GetUserByAPIKey(r.Context(), apiKey)
		if err != nil {
			respondWithError(w, 403, fmt.Sprintf("Could not get user: %v", err))
			return
		}

		handler(w, r, user)
	}
}
