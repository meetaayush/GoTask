package handlers

import (
	"net/http"

	"github.com/meetaayush/gotask/models"
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
	data := make(map[string]string)
	data["title"] = "Gotask | Homepage"
	tm := &models.TemplateModel{
		StringMap: data,
	}
	Render(w, r, pages.Home(tm))
}

func (repo *Repository) Signup(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]string)
	data["title"] = "Gotask | Register a new account"
	tm := &models.TemplateModel{
		StringMap: data,
	}
	Render(w, r, auth.Signup(tm))
}

func (repo *Repository) Signin(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]string)
	data["title"] = "Gotask | Login to your account"
	tm := &models.TemplateModel{
		StringMap: data,
	}
	Render(w, r, auth.Signin(tm))
}
