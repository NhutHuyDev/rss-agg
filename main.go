package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/NhutHuyDev/rss-agg/internal/db"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *db.Queries
}

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

	apiCfg := apiConfig{
		DB: db.New(conn),
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
	v1Router.Post("/users", apiCfg.HandlerCreateUser)

	router.Mount("/v1", v1Router)

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
