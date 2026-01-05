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
	tmpl = template.Must(template.ParseFiles("index.html")) //relation between GO and HTML
)

// ========================
func handler(w http.ResponseWriter, req *http.Request) { //separeted handler function
	data := SiteData{
		email:    req.FormValue("email"),
		password: req.FormValue("password"),
	}
	data.Success = true
	data.StorageAccess = "Registrtion Completed!"
	tmpl.Execute(w, data)
}

// ========================
func main() {
	http.HandleFunc("/", handler) // "/" - root
	log.Fatal(http.ListenAndServe(":80", nil))
}
