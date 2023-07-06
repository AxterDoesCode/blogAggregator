package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"

	httphandler "github.com/AxterDoesCode/blogAggregator/pkg/httpHandler"
)

func main() {
	godotenv.Load()
	port := os.Getenv("PORT")
	r := chi.NewRouter()
	v1Router := chi.NewRouter()
	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))
	r.Mount("/v1", v1Router)
	v1Router.Get("/readiness", httphandler.Readiness)
	v1Router.Get("/err", httphandler.ErrHandler)
	server := &http.Server{
		Addr:    ":" + port,
		Handler: r,
	}
	log.Printf("Serving on port : %s\n", port)
	log.Fatal(server.ListenAndServe())
}
