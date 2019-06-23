package main

import (
	"fmt"
	"net/http"

	"github.com/houseofbosons/houseofbosons/services/backend/router"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/api", http.StatusSeeOther)
	})

	http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to houseofbosons!")
	})

	http.HandleFunc("/api/auth/google", router.GoogleLoginHandeler)
	http.HandleFunc("/api/auth/google/callback", router.GoogleCallbackHandler)

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.ListenAndServe(":8080", nil)
}
