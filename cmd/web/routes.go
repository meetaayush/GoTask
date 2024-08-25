package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/meetaayush/gotask/pkgs/handlers"
)

func routes() http.Handler {
	mux := chi.NewRouter()

	mux.Handle("/static/*", handlers.Static())

	mux.Get("/favicon.ico", handlers.Favicon)
	mux.Get("/", handlers.Repo.Home)
	mux.Get("/signup", handlers.Repo.Signup)
	mux.Get("/signin", handlers.Repo.Signin)
	mux.Get("/forgot-password", handlers.Repo.ForgotPassword)
	return mux
}
