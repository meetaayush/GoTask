package handlers

import (
	"io/fs"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

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

// Favicon returns the favicon.ico file
func Favicon(w http.ResponseWriter, r *http.Request) {
	// Specify the path to favicon.ico file
	faviconPath := filepath.Join("cmd", "web", "static", "images", "favicon.ico")

	// Open the favicon file
	file, err := os.Open(faviconPath)
	if err != nil {
		http.Error(w, "Favicon not found", http.StatusNotFound)
		return
	}
	defer file.Close()

	// Set the content type
	w.Header().Set("Content-Type", "image/x-icon")

	// Serve the file
	http.ServeContent(w, r, "favicon.ico", time.Now(), file)
}
