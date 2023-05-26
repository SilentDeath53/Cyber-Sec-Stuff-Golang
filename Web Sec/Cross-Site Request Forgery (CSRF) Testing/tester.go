package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type CSRFToken struct {
	Token string
}

func generateCSRFToken() string {
	// Replace this with your actual CSRF token generation logic
	return "random_csrf_token"
}

func verifyCSRFToken(token string) bool {
	// Replace this with your actual CSRF token verification logic
	return token == "random_csrf_token"
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	csrfToken := generateCSRFToken()

	// Save CSRF token in session or as a hidden field in the form
	// Here, we save it as a hidden field in the form for simplicity
	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := CSRFToken{Token: csrfToken}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func processFormHandler(w http.ResponseWriter, r *http.Request) {
	csrfToken := r.FormValue("csrfToken")

	if !verifyCSRFToken(csrfToken) {
		http.Error(w, "Invalid CSRF token", http.StatusBadRequest)
		return
	}

	fmt.Fprintln(w, "Form submitted successfully!")
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/process", processFormHandler)

	fmt.Println("Server started on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
