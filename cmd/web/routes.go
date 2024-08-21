package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/meetaayush/gotask/pkgs/handlers"
)

func routes() http.Handler {
	mux := chi.NewRouter()

	mux.Handle("/static/*", handlers.Static())

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/signup", handlers.Repo.Signup)
	return mux
}
