package handlers

import (
	"net/http"

	"github.com/meetaayush/gotask/pkgs/config"
	"github.com/meetaayush/gotask/views/pages"
	"github.com/meetaayush/gotask/views/pages/auth"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

func NewHandler(r *Repository) {
	Repo = r
}

func (repo *Repository) Home(w http.ResponseWriter, r *http.Request) {
	Render(w, r, pages.Home())
}

func (repo *Repository) Signup(w http.ResponseWriter, r *http.Request) {
	Render(w, r, auth.Signup())
}
