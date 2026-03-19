package internal

import (
	"html/template"
	"log"
	"net/http"
	"time"
)

// User represents a user object
type User struct {
	ID    int
	Name  string
	Email string
}

// ProvisionRequest is the request object containing the form data
type ProvisionRequest struct {
	FirstName string
	LastName  string
	Email     string
	Services  []string
}

// RootHandler closure renders the index.html template and returns http.HandlerFunc
func RootHandler(tmpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := tmpl.ExecuteTemplate(w, "index.html", nil)
		if err != nil {
			return
		}
	}
}

// SearchHandler closure renders the search.html template and returns http.HandlerFunc
func SearchHandler(tmpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := tmpl.ExecuteTemplate(w, "search.html", nil)
		if err != nil {
			return
		}
	}
}

func SearchResultsHandler(tmpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query().Get("search")
		log.Printf("searching for: %s", query)

		// Just a fake delay to simulate api and trigger the progress bar
		time.Sleep(3 * time.Second)

		users := []User{
			{ID: 1, Name: "Gandalf the Grey", Email: "gandalf@treemail.com"},
			{ID: 2, Name: "Aragorn Elessar", Email: "aragorn@gondor.com"},
			{ID: 3, Name: "Legolas Greenleaf", Email: "legolas@mirkwood.com"},
			{ID: 4, Name: "Gimli son of Gloin", Email: "gimli@erebor.com"},
			{ID: 5, Name: "Frodo Baggins", Email: "frodo@shire.com"},
			{ID: 6, Name: "Samwise Gamgee", Email: "sam@shire.com"},
			{ID: 7, Name: "Boromir of Gondor", Email: "boromir@gondor.com"},
			{ID: 8, Name: "Peregrin Took", Email: "pippin@shire.com"},
			{ID: 9, Name: "Meriadoc Brandybuck", Email: "merry@shire.com"},
			{ID: 10, Name: "Elrond Halfelven", Email: "elrond@rivendell.com"},
			{ID: 11, Name: "Gandalf the Grey", Email: "gandalf@treemail.com"},
			{ID: 12, Name: "Aragorn Elessar", Email: "aragorn@gondor.com"},
			{ID: 13, Name: "Legolas Greenleaf", Email: "legolas@mirkwood.com"},
			{ID: 14, Name: "Gimli son of Gloin", Email: "gimli@erebor.com"},
			{ID: 15, Name: "Frodo Baggins", Email: "frodo@shire.com"},
			{ID: 16, Name: "Samwise Gamgee", Email: "sam@shire.com"},
			{ID: 17, Name: "Boromir of Gondor", Email: "boromir@gondor.com"},
			{ID: 18, Name: "Peregrin Took", Email: "pippin@shire.com"},
			{ID: 19, Name: "Meriadoc Brandybuck", Email: "merry@shire.com"},
			{ID: 20, Name: "Elrond Halfelven", Email: "elrond@rivendell.com"},
		}

		err := tmpl.ExecuteTemplate(w, "search-results", users)
		if err != nil {
			log.Printf("template error: %v", err)
			return
		}
	}
}

// ProvisionHandler renders the provision-form fragment or provision-validate fragment and returns http.HandlerFunc
func ProvisionHandler(tmpl *template.Template, client *AuthentikClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			log.Printf("error parsing form: %v", err)
			return
		}

		// Create a new ProvisionRequest object from form data
		provisionRequest := &ProvisionRequest{
			FirstName: r.FormValue("first_name"),
			LastName:  r.FormValue("last_name"),
			Email:     r.FormValue("email"),
			Services:  r.Form["services"],
		}

		// Form validation on any empty fields
		if provisionRequest.FirstName == "" || provisionRequest.LastName == "" || provisionRequest.Email == "" || len(provisionRequest.Services) == 0 {
			err := tmpl.ExecuteTemplate(w, "provision-validate", provisionRequest)
			if err != nil {
				return
			}
			return
		}

		// Create the user request
		_, err = client.CreateUserRequest(provisionRequest)
		if err != nil {
			// API call fails, render failure fragment
			err = tmpl.ExecuteTemplate(w, "provision-failure", provisionRequest)
			if err != nil {
				return
			}
			return
		}

		// API calls success, render success fragment
		err = tmpl.ExecuteTemplate(w, "provision-success", provisionRequest)
		if err != nil {
			return
		}
	}
}
