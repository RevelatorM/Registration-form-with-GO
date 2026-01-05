package main

import (
	"html/template"
	"log"
	"net/http"
)

// ========================
type SiteData struct {
	email         string
	password      string
	Success       bool
	StorageAccess string
}

// ========================
var (
	formTmpl    = template.Must(template.ParseFiles("../index.html")) //relation between GO and HTML
	successTmpl = template.Must(template.ParseFiles("../success.html"))
)

// ========================
func handler(w http.ResponseWriter, req *http.Request) { //separeted handler function
	if req.Method == http.MethodGet {
		formTmpl.Execute(w, nil)
		return
	}

	if req.Method == http.MethodPost {
		email := req.FormValue("email")
		password := req.FormValue("password")

		log.Println("Email:", email)
		log.Println("Password:", password)

		http.Redirect(w, req, "/success", http.StatusSeeOther)
		return
	}
}

// ========================
func successHandler(w http.ResponseWriter, req *http.Request) {
	successTmpl.Execute(w, nil)
}

// ========================
func main() {
	http.HandleFunc("/", handler) // "/" - root
	http.HandleFunc("/success", successHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
