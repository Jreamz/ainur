package internal

import (
	"fmt"
	"html/template"
	"net/http"
)

type provisionRequest struct {
	firstName string
	lastName  string
	email     string
	services  []string
}

func RootHandler(w http.ResponseWriter, _ *http.Request) {
	tmpl := template.Must(template.ParseGlob("static/*.html"))

	err := tmpl.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		return
	}
}

func ProvisionHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		//fmt.Errorf("error parsing form: %w", err)
		return
	}

	provisionRequest := &provisionRequest{
		firstName: r.FormValue("first_name"),
		lastName:  r.FormValue("last_name"),
		email:     r.FormValue("email"),
		services:  r.Form["services"],
	}

	fmt.Println(provisionRequest)
	w.WriteHeader(http.StatusOK)
	write, err := w.Write([]byte("User provisioning..."))
	if err != nil {
		return
	}
	fmt.Println(write)
}
