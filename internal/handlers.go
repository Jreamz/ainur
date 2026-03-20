package internal

import (
	"html/template"
	"log"
	"net/http"
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

// RootHandler closure renders the home.html template and returns http.HandlerFunc
func RootHandler(tmpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := tmpl.ExecuteTemplate(w, "home.html", nil)
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
			log.Printf("template error: %v", err)
			return
		}
	}
}

func SearchResultsHandler(tmpl *template.Template, client *AuthentikClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query().Get("search")
		log.Printf("searching for: %s", query)

		users, _ := client.SearchUsersList(query)

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

		// API call success, render success fragment
		err = tmpl.ExecuteTemplate(w, "provision-success", provisionRequest)
		if err != nil {
			return
		}
	}
}
