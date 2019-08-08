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
	http.HandleFunc("/thread/", renderers.ThreadHandler)

	http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to houseofbosons api!")
	})

	http.HandleFunc("/api/auth/google", handlers.GoogleLogin)
	http.HandleFunc("/api/auth/google/callback", handlers.GoogleCallback)
	http.HandleFunc("/api/auth/current_user", middleware.CheckAuth(handlers.CurrentUser))

	http.HandleFunc("/api/post/all", handlers.CreateBlog)
	http.HandleFunc("/api/post/new", handlers.ReadBlogs)
	http.HandleFunc("/api/post/single", handlers.ReadBlog)
	http.HandleFunc("/api/post/edit", handlers.EditBlog)
	http.HandleFunc("/api/post/delete/temp", handlers.DeleteBlog)
	http.HandleFunc("/api/post/delete/perm", handlers.DeleteBlogPrem)

	http.HandleFunc("/api/post/idstr/available", handlers.CheckIDStr)

	http.HandleFunc("/api/topic/new", handlers.CreateTopic)
	http.HandleFunc("/api/topic/all", handlers.ReadTopics)
	http.HandleFunc("/api/topic/edit", handlers.EditTopic)
	http.HandleFunc("/api/topic/delete", handlers.DeleteTopic)

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.ListenAndServe(":8080", nil)
}
