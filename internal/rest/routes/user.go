package routes

import (
	"github.com/NhutHuyDev/rss-agg/internal/rest"
	"github.com/go-chi/chi/v5"
)

func NewUserRoute(apiCfg rest.APIConfig) *chi.Mux {
	router := chi.NewRouter()
	router.Get("/users", apiCfg.MiddlewareAuth(apiCfg.HandlerGetUser))
	router.Post("/users", apiCfg.HandlerCreateUser)

	return router
}
