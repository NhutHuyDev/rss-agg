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
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
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
		DB: queries,
		UserService: &services.UserServiceImpl{
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
	v1Router := chi.NewRouter()
	v1Router.Get("/healthz", HandlerReadiness)

	router.Mount("/v1", routes.NewUserRoute(apiCfg))

	// v1Router.Post("/feeds", apiCfg.middlewareAuth(apiCfg.HandlerCreateFeed))
	// v1Router.Get("/feeds", apiCfg.HandlerGetFeeds)

	// v1Router.Get("/feed_follows", apiCfg.middlewareAuth(apiCfg.HandlerGetFeedFollows))
	// v1Router.Post("/feed_follows", apiCfg.middlewareAuth(apiCfg.HandlerCreateFeedFollow))
	// v1Router.Delete("/feed_follows/{feed_folow_id}", apiCfg.middlewareAuth(apiCfg.HandlerDeleteFeedFollows))

	// v1Router.Get("/posts", apiCfg.middlewareAuth(apiCfg.HandlerGetPostsForUser))

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
