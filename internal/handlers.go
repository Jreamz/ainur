package internal

import (
	"html/template"
	"log"
	"net/http"
)

// CreateUserRequest is the request object containing the form data
type CreateUserRequest struct {
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

// SearchResultsHandler closure calls the client SearchUsersList, renders search-results htmx fragment and returns http.HandlerFunc
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

// CreateUserHandler renders the provision-form fragment or provision-validate fragment and returns http.HandlerFunc
func CreateUserHandler(tmpl *template.Template, client *AuthentikClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			log.Printf("error parsing form: %v", err)
			return
		}

		// Create a new CreateUserRequest object from form data
		request := &CreateUserRequest{
			FirstName: r.FormValue("first_name"),
			LastName:  r.FormValue("last_name"),
			Email:     r.FormValue("email"),
			Services:  r.Form["services"],
		}

		// Form validation on any empty fields
		if request.FirstName == "" || request.LastName == "" || request.Email == "" || len(request.Services) == 0 {
			err := tmpl.ExecuteTemplate(w, "form-validate", request)
			if err != nil {
				log.Printf("error parsing form: %v", err)
				return
			}
			return
		}

		// Create the user request
		_, err = client.CreateUserRequest(request)
		if err != nil {
			// API call fails, render failure fragment
			err = tmpl.ExecuteTemplate(w, "create-form-failure", request)
			if err != nil {
				log.Printf("error parsing form: %v", err)
				return
			}
			return
		}

		// API call success, render success fragment
		err = tmpl.ExecuteTemplate(w, "create-form-success", request)
		if err != nil {
			log.Printf("error parsing form: %v", err)
			return
		}
	}
}
