package internal

import (
	"fmt"
	"html/template"
	"net/http"
)

func RootHandler(w http.ResponseWriter, _ *http.Request) {
	tmpl := template.Must(template.ParseGlob("static/*.html"))

	err := tmpl.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		return
	}
}

func ProvisionHandler(_ http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		return
	}

	services := r.Form["services"]
	fmt.Println(services)
	t := fmt.Sprintf("%T", services)
	fmt.Println(t)
}
