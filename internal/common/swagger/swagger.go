package swagger

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

func ServeSwagger(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./docs/swagger.json")
}

func AddSwaggerRoutes(r chi.Router) {
	r.Get("/swagger/*", func(w http.ResponseWriter, r *http.Request) {
		http.StripPrefix("/swagger", http.FileServer(http.Dir("./docs"))).ServeHTTP(w, r)
	})
	r.Get("/swagger.json", ServeSwagger)
}
