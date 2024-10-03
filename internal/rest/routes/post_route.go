package routes

import (
	"github.com/NhutHuyDev/rss-agg/internal/rest"
	"github.com/go-chi/chi/v5"
)

func NewPostRoute(apiCfg rest.APIConfig) *chi.Mux {
	router := chi.NewRouter()
	router.With(apiCfg.AuthMiddleware).Get("/", apiCfg.HandlerGetPostsByUser)
	return router
}
