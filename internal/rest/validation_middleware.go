package rest

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/NhutHuyDev/rss-agg/internal/utils"
)

type DecodeBodyCxtKeyType string

const DecodeBodyCxtKey = DecodeBodyCxtKeyType("decodeBody")

func (apiCfg *APIConfig) ValidationMiddleware(params interface{}) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			decoder := json.NewDecoder(r.Body)

			err := decoder.Decode(params)
			if err != nil {
				utils.RespondWithError(w, 400, fmt.Sprintf("error parsing JSON: %v", err))
				return
			}

			err = apiCfg.Validate.Struct(params)
			if err != nil {
				utils.RespondWithError(w, 400, fmt.Sprintf("DTO error: %v", err))
				return
			}

			ctx := context.WithValue(r.Context(), DecodeBodyCxtKey, params)

			r = r.WithContext(ctx)

			next.ServeHTTP(w, r)
		})
	}
}
