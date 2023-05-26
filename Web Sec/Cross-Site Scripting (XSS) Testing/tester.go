package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type TemplateData struct {
	Message string
}

func renderTemplate(w http.ResponseWriter, tmpl string, data TemplateData) {
	t, err := template.ParseFiles(tmpl)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = t.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func vulnerableHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		message := r.FormValue("message")

		data := TemplateData{
			Message: message,
		}

		renderTemplate(w, "vulnerable.html", data)
	} else {
		renderTemplate(w, "form.html", TemplateData{})
	}
}

func main() {
	http.HandleFunc("/vulnerable", vulnerableHandler)

	fmt.Println("Server started on port 8080")
	http.ListenAndServe(":8080", nil)
}
