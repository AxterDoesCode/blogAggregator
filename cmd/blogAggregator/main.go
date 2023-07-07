package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"

	"github.com/AxterDoesCode/blogAggregator/internal/database"
	"github.com/AxterDoesCode/blogAggregator/pkg/apiconfig"
	httphandler "github.com/AxterDoesCode/blogAggregator/pkg/httpHandler"
)

func main() {
	godotenv.Load()
	port := os.Getenv("PORT")
	r := chi.NewRouter()
	v1Router := chi.NewRouter()
	dbURL := os.Getenv("GOOSE_DBSTRING")
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal(err)
		return
	}

	dbQueries := database.New(db)

	apiCfg := apiconfig.ApiConfig{
		DB: dbQueries,
	}

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
	v1Router.Post("/users", apiCfg.HandleCreateUser)

	server := &http.Server{
		Addr:    ":" + port,
		Handler: r,
	}
	log.Printf("Serving on port : %s\n", port)
	log.Fatal(server.ListenAndServe())
}
