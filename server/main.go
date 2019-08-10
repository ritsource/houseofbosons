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

	mux.HandleFunc("/api/post/new", middleware.CheckAuth(handlers.CreateBlog))
	mux.HandleFunc("/api/post/all", middleware.CheckAuth(handlers.ReadBlogs))
	mux.HandleFunc("/api/post/single", middleware.CheckAuth(handlers.ReadBlog))
	mux.HandleFunc("/api/post/edit", middleware.CheckAuth(handlers.EditBlog))
	mux.HandleFunc("/api/post/delete/temp", middleware.CheckAuth(handlers.DeleteBlog))
	mux.HandleFunc("/api/post/delete/perm", middleware.CheckAuth(handlers.DeleteBlogPrem))

	mux.HandleFunc("/api/post/idstr/available", middleware.CheckAuth(handlers.IDStrAvailable))

	mux.HandleFunc("/api/topic/new", middleware.CheckAuth(handlers.CreateTopic))
	mux.HandleFunc("/api/topic/all", middleware.CheckAuth(handlers.ReadTopics))
	mux.HandleFunc("/api/topic/edit", middleware.CheckAuth(handlers.EditTopic))
	mux.HandleFunc("/api/topic/delete", middleware.CheckAuth(handlers.DeleteTopic))

	fs := http.FileServer(http.Dir("static/"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	handler := cors.New(cors.Options{
		AllowedOrigins:   []string{os.Getenv("ADMIN_ORIGIN")},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowCredentials: true,
	}).Handler(mux)

	http.ListenAndServe(":8080", handler)
}
