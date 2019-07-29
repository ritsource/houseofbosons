package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/houseofbosons/houseofbosons/server/db"
	"github.com/houseofbosons/houseofbosons/server/middleware"
	"github.com/sirupsen/logrus"

	"github.com/houseofbosons/houseofbosons/server/router"
)

func main() {
	blog := db.Blog{
		ID:            "blog-no-2",
		Title:         "title",
		Description:   "Something",
		Author:        "Something",
		FormattedDate: "Something",
		DocType:       "Something",
		MDSrc:         "Something",
		HTMLSrc:       "Something",
		CreatedAt:     time.Now(),
	}

	x, err := blog.Insert()

	// new := db.Blog{
	// 	ID:            "blog-no-1",
	// 	Title:         "title",
	// 	Description:   "Something",
	// 	Author:        "Some More",
	// 	FormattedDate: "Something",
	// 	DocType:       "Something",
	// 	MDSrc:         "Something",
	// 	HTMLSrc:       "Something",
	// 	CreatedAt:     time.Now(),
	// 	IsFeatured:    false,
	// 	IsPublic:      false,
	// 	IsDeleted:     false,
	// }

	// x, err := blog.DeletePerm()
	if err != nil {
		logrus.Errorf("%v", err)
	}

	fmt.Printf("%+v\n", x)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to houseofbosons!")
	})

	http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to houseofbosons api!")
	})

	http.HandleFunc("/api/auth/google", router.GoogleLogin)
	http.HandleFunc("/api/auth/google/callback", router.GoogleCallback)
	http.HandleFunc("/api/auth/current_user", middleware.CheckAuth(router.CurrentUser))

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.ListenAndServe(":8080", nil)
}
