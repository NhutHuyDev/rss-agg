package routes

import (
	"github.com/NhutHuyDev/rss-agg/api"
	"github.com/NhutHuyDev/rss-agg/internal/rest"
	"github.com/go-chi/chi/v5"
)

func NewFeedRoute(apiCfg rest.APIConfig) *chi.Mux {
	router := chi.NewRouter()
	router.Get("/", apiCfg.HandlerGetFeeds)
	router.With(apiCfg.AuthMiddleware, apiCfg.ValidationMiddleware(&api.CreateFeedDTO{})).Post("/", apiCfg.HandlerCreateFeed)

	return router
}
