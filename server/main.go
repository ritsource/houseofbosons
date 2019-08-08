package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/houseofbosons/houseofbosons/server/middleware"
	"github.com/houseofbosons/houseofbosons/server/renderers"
	"github.com/rs/cors"

	"github.com/houseofbosons/houseofbosons/server/handlers"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", renderers.IndexHandler)
	mux.HandleFunc("/posts", renderers.BlogsHandler)
	mux.HandleFunc("/post/", renderers.BlogHandler)
	mux.HandleFunc("/thread/", renderers.ThreadHandler)

	mux.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to houseofbosons api!")
	})

	mux.HandleFunc("/api/auth/google", handlers.GoogleLogin)
	mux.HandleFunc("/api/auth/google/callback", handlers.GoogleCallback)
	mux.HandleFunc("/api/auth/current_user", middleware.CheckAuth(handlers.CurrentUser))

	mux.HandleFunc("/api/post/all", handlers.CreateBlog)
	mux.HandleFunc("/api/post/new", handlers.ReadBlogs)
	mux.HandleFunc("/api/post/single", handlers.ReadBlog)
	mux.HandleFunc("/api/post/edit", handlers.EditBlog)
	mux.HandleFunc("/api/post/delete/temp", handlers.DeleteBlog)
	mux.HandleFunc("/api/post/delete/perm", handlers.DeleteBlogPrem)

	mux.HandleFunc("/api/post/idstr/available", handlers.CheckIDStr)

	mux.HandleFunc("/api/topic/new", handlers.CreateTopic)
	mux.HandleFunc("/api/topic/all", handlers.ReadTopics)
	mux.HandleFunc("/api/topic/edit", handlers.EditTopic)
	mux.HandleFunc("/api/topic/delete", handlers.DeleteTopic)

	fs := http.FileServer(http.Dir("static/"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	handler := cors.New(cors.Options{
		AllowedOrigins:   []string{os.Getenv("ADMIN_ORIGIN")},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowCredentials: true,
	}).Handler(mux)

	http.ListenAndServe(":8080", handler)
}
