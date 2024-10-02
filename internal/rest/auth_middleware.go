package rest

import (
	"context"
	"fmt"
	"net/http"

	utils "github.com/NhutHuyDev/rss-agg/pkg"
)

type UserCxtKeyType string

const UserCxtKey = UserCxtKeyType("user")

func (apiCfg *APIConfig) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := utils.GetApiKey(r.Header)
		if err != nil {
			utils.RespondWithError(w, 403, fmt.Sprintf("Auth error: %v", err))
			return
		}

		apiCfg.UserService.SetContext(r.Context())
		user, err := apiCfg.UserService.GetUserByAPIKey(apiKey)
		if err != nil {
			utils.RespondWithError(w, 403, fmt.Sprintf("Could not get user: %v", err))
			return
		}

		ctx := context.WithValue(r.Context(), UserCxtKey, user)

		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
