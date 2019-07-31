package renderers

import (
	"fmt"
	"math"
	"net/http"
	"strconv"
	"text/template"

	"github.com/houseofbosons/houseofbosons/server/db"
	"gopkg.in/mgo.v2/bson"
)

/*
PageIdxes holds values related to page-navigation in `/posts` handler.
Last, is the lastest possible page, and Current represents the currently
requested page for by the client
*/
type PageIdxes struct {
	Last    int
	Current int
}

/*
BlogsHandler renders lists out all the blogs as UI for the `/posts`
page, renders only the public `blogs` and includes page navigation
*/
func BlogsHandler(w http.ResponseWriter, r *http.Request) {
	nbpp := 3          // max number of blog-posts to be shown in 1 page
	num := PageIdxes{} // num holds page index related daa

	// reading the page-number from query staring
	var err error
	num.Current, _ = strconv.Atoi(r.URL.Query().Get("pagenum"))

	// bs represents the slice of blogs
	var bs db.Blogs

	// reading the number of documents from database
	nbs, err := bs.Count(bson.M{"is_deleted": false, "is_public": true})
	if err != nil {
		renderErr(w, 422, "Unable to Read from Database")
	}

	// the last possible page-index
	num.Last = int(math.Ceil(float64(nbs / nbpp)))

	// if requested page-index doesn't exist the set it back to 0
	if num.Current > num.Last {
		num.Current = num.Last
	}

	// reading all the requested blogs posts data (only the ones taht exists in this page)
	err = bs.ReadFew(bson.M{"is_deleted": false, "is_public": true}, bson.M{}, num.Current*nbpp, nbpp)
	if err != nil {
		renderErr(w, 422, "Unable to Read from Database")
	}

	// parsing templates
	t, err := template.ParseFiles(
		"static/pages/posts.html",
		"static/partials/header.html",
		"static/partials/footer.html",
	)
	if err != nil {
		writeErr(w, 500, err)
	}

	fmt.Printf("%+v\n", len(bs))

	err = t.Execute(w, struct {
		PageData PageIdxes
		Posts    db.Blogs
	}{
		PageData: num,
		Posts:    bs,
	})
	if err != nil {
		writeErr(w, 500, err)
	}
}
