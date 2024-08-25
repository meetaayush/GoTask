package handlers

import (
	"bytes"
	"io/fs"
	"log"
	"net/http"
	"time"

	"github.com/a-h/templ"
	"github.com/meetaayush/gotask/pkgs/assets"
)

func Render(w http.ResponseWriter, r *http.Request, template templ.Component) error {
	return template.Render(r.Context(), w)
}

func Static() http.Handler {
	staticFs, err := fs.Sub(assets.StaticFiles, "static")
	if err != nil {
		log.Fatal("error creating filesystem: ", err)
	}
	return http.StripPrefix("/static/", http.FileServer(http.FS(staticFs)))
}

// Favicon returns the favicon.ico file
func Favicon(w http.ResponseWriter, r *http.Request) {
	file, err := assets.StaticFiles.ReadFile("static/images/favicon.ico")
	if err != nil {
		http.Error(w, "error finding favicon file", http.StatusNotFound)
		return
	}

	http.ServeContent(w, r, "favicon.ico", time.Now(), bytes.NewReader(file))
}
