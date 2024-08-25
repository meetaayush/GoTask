package handlers

import (
	"io/fs"
	"log"
	"net/http"

	"github.com/a-h/templ"
)

func Render(w http.ResponseWriter, r *http.Request, template templ.Component) error {
	return template.Render(r.Context(), w)
}

func Static(staticFiles fs.FS) http.Handler {
	staticFs, err := fs.Sub(staticFiles, "static")
	if err != nil {
		log.Fatal("error creating filesystem: ", err)
	}
	return http.StripPrefix("/static/", http.FileServer(http.FS(staticFs)))
}
