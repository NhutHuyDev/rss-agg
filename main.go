package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")

	port := os.Getenv("PORT")

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
	v1Router.HandleFunc("/healthz", HandlerReadiness)
	v1Router.HandleFunc("/notfound", HandlerErr)

	router.Mount("/v1", v1Router)

	server := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}

	log.Printf("Server is starting on port %v", port)
	// listenAndServer() is a blocking function
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
