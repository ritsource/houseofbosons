package main

import (
	"fmt"
	"net/http"

	"github.com/houseofbosons/houseofbosons/server/middleware"
	"github.com/houseofbosons/houseofbosons/server/renderers"

	"github.com/houseofbosons/houseofbosons/server/handlers"
)

func main() {
	http.HandleFunc("/", renderers.IndexHandler)
	http.HandleFunc("/posts", renderers.BlogsHandler)
	http.HandleFunc("/post/", renderers.BlogHandler)

	http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to houseofbosons api!")
	})

	http.HandleFunc("/api/auth/google", handlers.GoogleLogin)
	http.HandleFunc("/api/auth/google/callback", handlers.GoogleCallback)
	http.HandleFunc("/api/auth/current_user", middleware.CheckAuth(handlers.CurrentUser))

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.ListenAndServe(":8080", nil)
}
