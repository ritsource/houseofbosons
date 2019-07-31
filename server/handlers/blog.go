/*
`blog.go` contains handler function for houseofboson-blogs-api
*/

package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/houseofbosons/houseofbosons/server/db"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// func init() {
// 	b := db.Blog{
// 		URLIDs:        []string{"blog-3", "third-blog"},
// 		Title:         "Title-3",
// 		Description:   "A monorepo containing all the libraries and services for House-of-Bosons, a Science-blog on Physics, Math, Astronomy & More.",
// 		Author:        "Ritwik Saha",
// 		FormattedDate: "21  July, 2019",
// 		DocType:       db.DocTypeMD,
// 		MDSrc:         "https://gitlab.com/ritwik310/blog-documents/raw/master/Write-a-Torrent-Client-in-Go-0/Write-a-Torrent-Client-in-Go-0.md",
// 		Thumbnail:     "https://gitlab.com/ritwik310/blog-documents/raw/master/Write-a-Torrent-Client-in-Go-0/Torrent-Client-P2P-Messaging-3.png",
// 		IsFeatured:    false,
// 		IsPublic:      true,
// 		IsDeleted:     false,
// 		IsSeries:      false,
// 	}

// 	b.CreatedAt = int32(time.Now().Unix())

// 	err := b.Create()
// 	if err != nil {
// 		logrus.Errorf("%v\n", err)
// 		return
// 	}
// }

/*
CreateBlog creates a new blog from provided data read on `http.Request` body,
the body needs to contain all required fields in json, otherwise it will be
saved as `null` in the database
*/
func CreateBlog(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeErr(w, 404, fmt.Errorf("%v request to %v not found", r.Method, r.URL.Path))
		return
	}

	// reading the json body, provided in the request
	var b db.Blog
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&b)
	if err != nil {
		writeErr(w, 500, err)
		return
	}

	// manually setting `CreatedAt`, that holds creation time
	b.CreatedAt = int32(time.Now().Unix())

	// inserting a new document in the database
	err = b.Create()
	if err != nil {
		writeErr(w, 422, err)
		return
	}

	// redirecting to `/api/private/blogs` route handler
	http.Redirect(w, r, "/api/private/blogs", http.StatusTemporaryRedirect) // 302 - POST to GET
}

/*
ReadBlog reads a a single blog from the database by
it's `ID` and writes the back the data in json
*/
func ReadBlog(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeErr(w, 404, fmt.Errorf("%v request to %v not found", r.Method, r.URL.Path))
		return
	}

	// retrieving id from `query-string`
	id := r.URL.Query().Get("id")
	if len(id) == 0 {
		writeErr(w, http.StatusBadRequest, fmt.Errorf("no blog-id provided"))
		return
	}

	// reading the document from database
	var b db.Blog
	err := b.Read(bson.M{"_id": bson.ObjectIdHex(id)}, bson.M{})
	switch err {
	case mgo.ErrNotFound:
		writeErr(w, 404, err)
	case nil:
		// everything's fine
	default:
		writeErr(w, 500, err)
		return
	}

	// writing json data to the client
	writeJSON(w, b)
}

/*
ReadBlogs reads all the `blogs` documents from database without any filter,
and that's why only an authenticated admin client is allowed to access that
*/
func ReadBlogs(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeErr(w, 404, fmt.Errorf("%v request to %v not found", r.Method, r.URL.Path))
		return
	}

	// reading all the `blog` documents from database
	var bs db.Blogs
	err := bs.ReadAll(bson.M{}, bson.M{})
	if err != nil {
		writeErr(w, 500, err)
		return
	}

	writeJSON(w, bs)
}

/*
EditBlog edits a blog document
*/
func EditBlog(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		writeErr(w, 404, fmt.Errorf("%v request to %v not found", r.Method, r.URL.Path))
		return
	}

	// retrieving id from query string
	id := r.URL.Query().Get("id")
	if len(id) == 0 {
		writeErr(w, http.StatusBadRequest, fmt.Errorf("no blog-id provided"))
		return
	}

	// reading request body
	var body map[string]interface{} // because cannot use b (type db.Blog) as type bson.M in argument to bl.Update
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&body)
	if err != nil {
		writeErr(w, 500, err)
		return
	}

	// editing document
	var b db.Blog
	err = b.Update(bson.M{"_id": bson.ObjectIdHex(id)}, body) // Update Document in Database
	if err != nil {
		writeErr(w, 422, err)
		return
	}

	writeJSON(w, b) // Write Data
}

/*
DeleteBlog deleted a blog (not permanently) by updating
the document's `is_deleted` property to `true`
*/
func DeleteBlog(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		writeErr(w, 404, fmt.Errorf("%v request to %v not found", r.Method, r.URL.Path))
		return
	}

	// retrieving id from query string
	id := r.URL.Query().Get("id")
	if len(id) == 0 {
		writeErr(w, http.StatusBadRequest, fmt.Errorf("no blog-id provided"))
		return
	}

	// editing document to `is_deleted: true`
	var b db.Blog
	err := b.Delete(bson.ObjectIdHex(id))
	if err != nil {
		writeErr(w, 422, err)
		return
	}

	// writing the updated data
	writeJSON(w, b)
}

/*
DeleteBlogPrem permanently deletes a blog document from the database
*/
func DeleteBlogPrem(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		writeErr(w, 404, fmt.Errorf("%v request to %v not found", r.Method, r.URL.Path))
		return
	}

	// retrieving id from query string
	id := r.URL.Query().Get("id")
	if len(id) == 0 {
		writeErr(w, http.StatusBadRequest, fmt.Errorf("no blog-id provided"))
		return
	}

	// deleting document (permanently)
	var b db.Blog
	b.ID = bson.ObjectIdHex(id)
	err := b.DeletePermanent()
	if err != nil {
		writeErr(w, 422, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("{\"message\": \"Successfully deleted\"}"))
}
