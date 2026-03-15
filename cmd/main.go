package main

import (
	"ainur/internal"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {

	// Chi router
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	// Html
	r.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Routes
	r.Get("/", internal.RootHandler)
	r.Post("/provision", internal.ProvisionHandler)

	// Start server
	err := http.ListenAndServe(":3000", r)
	if err != nil {
		return
	}

}
