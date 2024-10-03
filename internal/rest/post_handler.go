package rest

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/NhutHuyDev/rss-agg/api"
	"github.com/NhutHuyDev/rss-agg/internal/domain"
	utils "github.com/NhutHuyDev/rss-agg/pkg"
)

func (apiCfg *APIConfig) HandlerGetPostsByUser(w http.ResponseWriter, r *http.Request) {
	user, _ := r.Context().Value(UserCxtKey).(domain.User)

	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil {
		utils.RespondWithError(w, 400, fmt.Sprintf("'limit' param must be a int number: %v", err))
		return
	}

	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		utils.RespondWithError(w, 400, fmt.Sprintf("'page' param must be a int number: %v", err))
		return
	}

	apiCfg.PostService.SetContext(r.Context())
	posts, err := apiCfg.PostService.GetPostsByUser(user.ID, page, limit)
	if err != nil {
		utils.RespondWithError(w, 400, fmt.Sprintf("Could not get posts: %v", err))
		return
	}

	if len(posts) == 0 {
		utils.RespondWithError(w, 404, "not found posts")
		return
	}

	totalPosts, err := apiCfg.PostService.CountPosts(user.ID)
	if err != nil {
		utils.RespondWithError(w, 400, fmt.Sprintf("Could not get posts: %v", err))
		return
	}

	totalPage, nextPage, _ := utils.GetPagination(totalPosts, limit, page)

	type GetPostsRes struct {
		Posts []api.PostRes `json:"posts"`
		utils.PaginationRes
	}

	getPostsRes := GetPostsRes{
		Posts: api.CastToPosts(posts),
		PaginationRes: utils.PaginationRes{
			Total:       totalPosts,
			TotalPage:   totalPage,
			NextPage:    nextPage,
			CurrentPage: page,
		},
	}

	utils.RespondWithJSON(w, 201, getPostsRes)
}
