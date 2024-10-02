package rest

import (
	"fmt"
	"net/http"

	"github.com/NhutHuyDev/rss-agg/internal/domain"
	"github.com/NhutHuyDev/rss-agg/internal/services"
	"github.com/NhutHuyDev/rss-agg/internal/utils"
)

type authedHandler func(http.ResponseWriter, *http.Request, domain.User)

func (apiCfg *APIConfig) MiddlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := utils.GetApiKey(r.Header)
		if err != nil {
			utils.RespondWithError(w, 403, fmt.Sprintf("Auth error: %v", err))
			return
		}

		user, err := apiCfg.DB.GetUserByAPIKey(r.Context(), apiKey)
		if err != nil {
			utils.RespondWithError(w, 403, fmt.Sprintf("Could not get user: %v", err))
			return
		}

		handler(w, r, services.CastToUser(user))
	}
}
