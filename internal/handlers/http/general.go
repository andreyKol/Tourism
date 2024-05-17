package http

import (
	"log/slog"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func (h HttpHandler) DocsPage(w http.ResponseWriter, r *http.Request) {
	filePath := filepath.Join("web", "pages", "api.html")

	_, err := os.Stat(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		ErrorResponse(w, r, err)
		return
	}

	http.ServeFile(w, r, filePath)
}

func (h HttpHandler) DocsFile(w http.ResponseWriter, r *http.Request) {
	filePath := filepath.Join("api", "app", "http.yml")

	_, err := os.Stat(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		ErrorResponse(w, r, err)
		return
	}

	http.ServeFile(w, r, filePath)
}

func (h HttpHandler) Static(w http.ResponseWriter, r *http.Request) {
	if !strings.HasPrefix(r.URL.Path, "/static/") {
		w.WriteHeader(http.StatusNotFound)
		slog.Error("url does not have /static/ prefix, but static handler is accessed")
		return
	}
	filename := strings.TrimPrefix(r.URL.Path, "/static/")

	http.ServeFile(w, r, path.Join("web", "static", filename))
}
