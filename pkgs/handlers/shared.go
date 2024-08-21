package handlers

import (
	"net/http"

	"github.com/a-h/templ"
)

func Render(w http.ResponseWriter, r *http.Request, template templ.Component) error {
	return template.Render(r.Context(), w)
}

func Static() http.Handler {
	return http.StripPrefix("/static/", http.FileServer(http.Dir("./static")))
}
