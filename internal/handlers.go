package internal

import (
	"html/template"
	"log"
	"net/http"
)

// ProvisionRequest is the request object containing the form data
type ProvisionRequest struct {
	FirstName string
	LastName  string
	Email     string
	Clearance string
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

// ProvisionHandler renders the provision-form fragment or provsion-validate fragment and returns http.HandlerFunc
func ProvisionHandler(tmpl *template.Template) http.HandlerFunc {
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

		// TODO API calls to endpoints here

		// API calls success, render success page
		err = tmpl.ExecuteTemplate(w, "provision-failure", provisionRequest)
		if err != nil {
			return
		}

		// API calls failed, render error page
	}
}
