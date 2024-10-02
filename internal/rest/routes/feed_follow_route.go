package routes

import (
	"github.com/NhutHuyDev/rss-agg/api"
	"github.com/NhutHuyDev/rss-agg/internal/rest"
	"github.com/go-chi/chi/v5"
)

func NewFeedFollowRoute(apiCfg rest.APIConfig) *chi.Mux {
	router := chi.NewRouter()

	router.With(apiCfg.AuthMiddleware).Get("/", apiCfg.HandlerGetFeedFollows)
	router.With(apiCfg.AuthMiddleware, apiCfg.ValidationMiddleware(&api.CreateFeedFollowDTO{})).Post("/", apiCfg.HandlerCreateFeedFollow)
	router.With(apiCfg.AuthMiddleware).Delete("/{feed_folow_id}", apiCfg.HandlerDeleteFeedFollows)

	return router
}
