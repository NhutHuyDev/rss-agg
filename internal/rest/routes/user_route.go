package routes

import (
	"github.com/NhutHuyDev/rss-agg/api"
	"github.com/NhutHuyDev/rss-agg/internal/rest"
	"github.com/go-chi/chi/v5"
)

const pattern = "/users"

func NewUserRoute(apiCfg rest.APIConfig) *chi.Mux {
	router := chi.NewRouter()
	router.With(apiCfg.AuthMiddleware).Get(pattern, apiCfg.HandlerGetUser)
	router.With(apiCfg.ValidationMiddleware(&api.CreateUserDTO{})).Post(pattern, apiCfg.HandlerCreateUser)

	return router
}
