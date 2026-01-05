package main

import (
	"log"
	"net/http"
)

// ========================
type SiteData struct {
	email    string
	password string
}

// ========================
func handler(w http.ResponseWriter, req *http.Request) { //separeted handler function
	data := SiteData {
		email: r.FormValue("email")
		password: r.FormValue("password")
	}
}

// ========================
func main() {
	http.HandlerFunc("/", handler) // "/" - root
	log.Fatal(http.ListenAndServe(":80", nil))
}
