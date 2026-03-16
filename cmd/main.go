package main

import (
	"ainur/internal"
	"html/template"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {

	// Chi router
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	// Serve all Html
	r.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Render all templates
	tmpl := template.Must(template.ParseGlob("static/*.html"))

	// Routes
	r.Get("/", internal.RootHandler(tmpl))
	r.Post("/provision", internal.ProvisionHandler(tmpl))

	// Start server
	err := http.ListenAndServe(":3000", r)
	if err != nil {
		return
	}

}
