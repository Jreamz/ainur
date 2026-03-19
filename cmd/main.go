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
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.URLFormat)

	// Serve all CSS, HTMX
	r.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Render all templates and fragments
	tmpl := template.New("")
	template.Must(tmpl.ParseGlob("static/templates/*.html"))
	template.Must(tmpl.ParseGlob("static/fragments/*.html"))

	// Create a request client
	client := internal.NewAuthentikClient("https://api.test.com", "token")

	// Routes
	r.Get("/", internal.RootHandler(tmpl))
	r.Post("/provision", internal.ProvisionHandler(tmpl, client))
	r.Get("/search", internal.SearchHandler(tmpl))
	r.Get("/search-results", internal.SearchResultsHandler(tmpl))

	// Start server - this block forever, nothing runs after this
	srvErr := http.ListenAndServe(":3000", r)
	if srvErr != nil {
		return
	}

}
