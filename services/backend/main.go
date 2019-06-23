package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/api", http.StatusSeeOther)
	})

	http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to houseofbosons!")
	})

	http.HandleFunc("/api/auth", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to auth!")
	})

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.ListenAndServe(":8080", nil)
}
