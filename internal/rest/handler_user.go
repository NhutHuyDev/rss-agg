package rest

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/NhutHuyDev/rss-agg/api"
	"github.com/NhutHuyDev/rss-agg/internal/domain"
	"github.com/NhutHuyDev/rss-agg/internal/utils"
	"github.com/google/uuid"
)

func (apiCfg *APIConfig) HandlerCreateUser(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	params := api.CreateUserReq{}
	err := decoder.Decode(&params)
	if err != nil {
		utils.RespondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	apiCfg.UserService.SetContext(r.Context())
	user, err := apiCfg.UserService.CreateUser(domain.User{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})
	if err != nil {
		utils.RespondWithError(w, 400, fmt.Sprintf("Could not create user: %v", err))
		return
	}

	utils.RespondWithJSON(w, 201, user)
}

func (apiCfg *APIConfig) HandlerGetUser(w http.ResponseWriter, r *http.Request, user domain.User) {
	utils.RespondWithJSON(w, 200, user)
}
