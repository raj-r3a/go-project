package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("PORT is not found in the env")
	}

	router := chi.NewRouter()

	corsOptions := cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodPost, http.MethodOptions},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}

	router.Use(cors.Handler(corsOptions))

	v1Router := chi.NewRouter()

	v1Router.Get("/health", handlerReadiness)
	v1Router.Get("/err", handlerError)

	router.Mount("/v1", v1Router)

	server := &http.Server{
		Handler: router,
		Addr:    ":" + port,
	}

	log.Printf("server starting at port %v", port)

	error := server.ListenAndServe()

	if error != nil {
		log.Fatal(error, "something went wrong in starting the server")
	}

	fmt.Println("PORT:", port)
}
