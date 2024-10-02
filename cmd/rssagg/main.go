package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/NhutHuyDev/rss-agg/internal/infra/db"
	"github.com/NhutHuyDev/rss-agg/internal/rest"
	"github.com/NhutHuyDev/rss-agg/internal/rest/routes"
	"github.com/NhutHuyDev/rss-agg/internal/services"
	utils "github.com/NhutHuyDev/rss-agg/pkg"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	godotenv.Load(".env")

	portStr := os.Getenv("PORT")
	if portStr == "" {
		log.Fatal("PORT is not found in the environment")
	}

	dbConnectioStr := os.Getenv("DB_CONNECTION_STR")
	if dbConnectioStr == "" {
		log.Fatal("DB_CONNECTION_STR is not found in the environment")
	}

	conn, err := sql.Open("postgres", dbConnectioStr)
	if err != nil {
		log.Fatal("Cannot connection to database")
	}

	queries := db.New(conn)

	// go StartScraping(db, 1, time.Minute)

	apiCfg := rest.APIConfig{
		DB:       queries,
		Validate: validator.New(),
		UserService: &services.UserServiceImpl{
			Queries: queries,
		},
		FeedService: &services.FeedServiceImpl{
			Queries: queries,
		},
		FeedFollowService: &services.FeedFollowServiceImpl{
			Queries: queries,
		},
	}

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	// Routes
	router.Get("/v1/healthz", func(w http.ResponseWriter, r *http.Request) {
		utils.RespondWithJSON(w, 200, struct {
			Status string `json:"status"`
		}{
			Status: "OK",
		})
	})

	router.Mount("/v1/users", routes.NewUserRoute(apiCfg))
	router.Mount("/v1/feeds", routes.NewFeedRoute(apiCfg))
	router.Mount("/v1/feed_follows", routes.NewFeedFollowRoute(apiCfg))

	// v1Router.Get("/posts", apiCfg.middlewareAuth(apiCfg.HandlerGetPostsForUser))

	router.NotFound(func(w http.ResponseWriter, r *http.Request) {
		utils.RespondWithError(w, 404, "not found")
	})

	server := &http.Server{
		Addr:    ":" + portStr,
		Handler: router,
	}

	log.Printf("Server is starting on port %v", portStr)

	// Note: listenAndServer() is a blocking function
	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
