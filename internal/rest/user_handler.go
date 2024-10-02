package rest

import (
	"fmt"
	"net/http"
	"time"

	"github.com/NhutHuyDev/rss-agg/api"
	"github.com/NhutHuyDev/rss-agg/internal/domain"
	utils "github.com/NhutHuyDev/rss-agg/pkg"
	"github.com/google/uuid"
)

func (apiCfg *APIConfig) HandlerCreateUser(w http.ResponseWriter, r *http.Request) {
	params, _ := r.Context().Value(DecodeBodyCxtKey).(*api.CreateUserDTO)

	apiCfg.UserService.SetContext(r.Context())
	user, err := apiCfg.UserService.CreateUser(domain.User{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})

	if err != nil {
		utils.RespondWithError(w, 400, fmt.Sprintf("could not create user: %v", err))
		return
	}

	utils.RespondWithJSON(w, 201, api.CastToUser(user))
}

func (apiCfg *APIConfig) HandlerGetUser(w http.ResponseWriter, r *http.Request) {
	user, ok := r.Context().Value(UserCxtKey).(domain.User)
	if !ok {
		utils.RespondWithError(w, 400, "could not get user")
		return
	}

	utils.RespondWithJSON(w, 200, api.CastToUser(user))
}
