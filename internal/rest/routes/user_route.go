package routes

import (
	"github.com/NhutHuyDev/rss-agg/api"
	"github.com/NhutHuyDev/rss-agg/internal/rest"
	"github.com/go-chi/chi/v5"
)

func NewUserRoute(apiCfg rest.APIConfig) *chi.Mux {
	router := chi.NewRouter()
	router.With(apiCfg.AuthMiddleware).Get("/", apiCfg.HandlerGetUser)
	router.With(apiCfg.ValidationMiddleware(&api.CreateUserDTO{})).Post("/", apiCfg.HandlerCreateUser)

	return router
}
