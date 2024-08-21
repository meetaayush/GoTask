package handlers

import (
	"fmt"
	"net/http"

	"github.com/meetaayush/gotask/pkgs/config"
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
	str := fmt.Sprintf("Hello from home page. Env: %v", repo.App.Environment)
	w.Write([]byte(str))
}

func (repo *Repository) Signup(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Sign up page"))
}
