package renderers

import (
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"strconv"
	"strings"
	"text/template"

	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday"
	"gopkg.in/mgo.v2"

	"github.com/ritsource/houseofbosons/server/db"
	"gopkg.in/mgo.v2/bson"
)

/*
GetDocument does
*/
func GetDocument(src string) ([]byte, error) {
	resp, err := http.Get(src)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

/*
BlogHandler renders
*/
func BlogHandler(w http.ResponseWriter, r *http.Request) {
	// reading the Blog-ID from URL path
	pts := strings.Split(r.URL.Path, "/")

	/*
		Also, because the handler pattern is has a `/` at the
		end of the string (like `/posts/`), we don't need to
		be worried about length of `pts` to be less than 3
	*/

	idstr := pts[2] // the Blog-ID

	// if the Blog-ID is empty then redirect to `/posts` route
	if idstr == "" {
		http.Redirect(w, r, "/posts", http.StatusSeeOther)
		return
	}

	// reading from blog-document database
	var b db.Blog
	err := b.Read(bson.M{"id_str": idstr, "is_deleted": false, "is_public": true}, bson.M{})
	switch err {
	case mgo.ErrNotFound:
		// if there's no piblic-blog found with requested IDStr, then rendering a 404-Not-Found error
		renderErr(w, 404, fmt.Sprintf("Post \"%v\" Not Found", idstr))
		return
	case nil:
		// if `err == nil` everything's fine
	default:
		// some internal errorerror
		renderErr(w, 500, fmt.Sprintf("Sorry, Unable to Read Data"))
		return
	}

	// if series then redirecting to appropriate `/thread/:id` route
	if b.IsSeries {
		http.Redirect(w, r, fmt.Sprintf("/thread/%v", idstr), http.StatusSeeOther)
		return
	}

	// source document type
	var src string
	switch b.DocType {
	case db.DocTypeMD:
		src = b.MDSrc
	case db.DocTypeHTML:
		src = b.HTMLSrc
	}

	// reading the source document
	doc, err := GetDocument(src)
	if err != nil {
		renderErr(w, 422, fmt.Sprintf("Sorry, Unable to Find the Document"))
		return
	}

	// parsing document data (html or markdown) into HTML (unsafe)
	var unsafe []byte
	if b.DocType == db.DocTypeMD {
		unsafe = blackfriday.MarkdownCommon(doc) // generating HTML from Markdown
	} else {
		unsafe = doc
	}

	// document HTML
	html := bluemonday.UGCPolicy().SanitizeBytes(unsafe)

	// parsing templates
	t, err := template.ParseFiles(
		"static/pages/each-post.html",
		"static/partials/header.html",
		"static/partials/footer.html",
		"static/partials/head-links.html",
	)
	if err != nil {
		renderErr(w, 500, "Internal Server Error")
		return
	}

	// executing template
	err = t.Execute(w, struct {
		Post db.Blog
		HTML string
	}{
		Post: b,
		HTML: string(html),
	})

	if err != nil {
		writeErr(w, 500, err)
	}
}

/*
ThreadHandler renders
*/
func ThreadHandler(w http.ResponseWriter, r *http.Request) {
	// reading the Blog-ID from URL path
	pts := strings.Split(r.URL.Path, "/")
	idstr := pts[2]

	// if the Blog-ID is empty then redirect to `/posts` route
	if idstr == "" {
		http.Redirect(w, r, "/posts", http.StatusSeeOther)
		return
	}

	// reading requested index from URL
	index, err := strconv.Atoi(r.URL.Query().Get("index"))
	if err != nil || index < 0 {
		renderErr(w, 400, "Invalid Index")
		return
	}

	// reading from blog-document database
	var b db.Blog
	err = b.Read(bson.M{"id_str": idstr, "is_deleted": false, "is_public": true, "is_series": true}, bson.M{})
	switch err {
	case mgo.ErrNotFound:
		// if there's no piblic-blog found with requested IDStr, then rendering a 404-Not-Found error
		renderErr(w, 404, fmt.Sprintf("Thread \"%v\" Not Found", idstr))
		return
	case nil:
		// if `err == nil` everything's fine
	default:
		// some internal errorerror
		renderErr(w, 500, fmt.Sprintf("Sorry, Unable to Read Data"))
		return
	}

	// If thread doesn't include any subblog
	if len(b.SubBlogs) == 0 {
		renderErr(w, http.StatusNoContent, "Sorry, Empty Thread")
		return
	}

	// If index overflow, then redirect to index 0
	if index+1 > len(b.SubBlogs) {
		http.Redirect(w, r, fmt.Sprintf("/thread/%v?index=0", idstr), http.StatusSeeOther)
		return
	}

	// source document type
	var src string
	switch b.SubBlogs[index].DocType {
	case db.DocTypeMD:
		src = b.SubBlogs[index].MDSrc
	case db.DocTypeHTML:
		src = b.SubBlogs[index].HTMLSrc
	}

	// reading the source document
	doc, err := GetDocument(src)
	if err != nil {
		renderErr(w, 422, fmt.Sprintf("Sorry, Unable to Find the Document"))
		return
	}

	// parsing document data (html or markdown) into HTML (unsafe)
	var unsafe []byte
	if b.DocType == db.DocTypeMD {
		unsafe = blackfriday.MarkdownCommon(doc) // generating HTML from Markdown
	} else {
		unsafe = doc
	}

	// document HTML
	html := bluemonday.UGCPolicy().SanitizeBytes(unsafe)

	// parsing templates
	t, err := template.ParseFiles(
		"static/pages/each-thread.html",
		"static/partials/header.html",
		"static/partials/footer.html",
		"static/partials/head-links.html",
	)
	if err != nil {
		renderErr(w, 500, "Internal Server Error")
		return
	}

	// num contains data for thread navigation
	num := NavData{
		Last:    len(b.SubBlogs) - 1,
		Current: index,
	}

	// executing template
	err = t.Execute(w, struct {
		Post        db.Blog
		NavArray    []NavArray
		NavCurrent  int
		OneBasedIdx int
		SubPost     db.SubBlog
		HTML        string
	}{
		Post:        b,
		NavArray:    getNavArray(num),
		NavCurrent:  index,
		OneBasedIdx: index + 1,
		SubPost:     b.SubBlogs[index],
		HTML:        string(html),
	})

	if err != nil {
		writeErr(w, 500, err)
	}
}

/*
BlogsHandler renders lists out all the blogs as UI for the `/posts`
page, renders only the public `blogs` and includes page navigation
*/
func BlogsHandler(w http.ResponseWriter, r *http.Request) {
	nbpp := 8        // max number of blog-posts to be shown in 1 page
	num := NavData{} // num holds page index related daa

	// reading the page-number from query staring
	var err error
	num.Current, _ = strconv.Atoi(r.URL.Query().Get("pagenum"))
	topic := r.URL.Query().Get("topic")

	// bs represents the slice of blogs
	var bs db.Blogs

	// mongodb filter
	var mgoSelector bson.M
	if topic != "" {
		mgoSelector = bson.M{"topics": bson.M{"$all": []string{topic}}, "is_deleted": false, "is_public": true}
	} else {
		mgoSelector = bson.M{"is_deleted": false, "is_public": true}
	}

	// reading the number of documents from database
	nbs, err := bs.Count(mgoSelector)
	if err != nil {
		renderErr(w, 422, "Unable to Read from Database")
		return
	}

	// the last possible page-index
	num.Last = int(math.Ceil(float64(nbs / nbpp)))

	// if requested page-index doesn't exist the set it back to 0
	if num.Current > num.Last {
		num.Current = num.Last
	}

	// reading all the requested blogs posts data (only the ones taht exists in this page)
	err = bs.ReadFew(mgoSelector, bson.M{}, num.Current*nbpp, nbpp)
	if err != nil {
		renderErr(w, 422, "Unable to Read Blogs from Database")
		return
	}

	// reading `topic` documents from the database
	var ts db.Topics
	err = ts.ReadAll()
	if err != nil {
		renderErr(w, 422, "Unable to Read Topics from Database")
		return
	}

	// parsing templates
	t, err := template.ParseFiles(
		"static/pages/posts.html",
		"static/partials/posts-item.html",
		"static/partials/header.html",
		"static/partials/footer.html",
		"static/partials/head-links.html",
	)
	if err != nil {
		writeErr(w, 500, err)
		return
	}

	err = t.Execute(w, struct {
		// NavIdxes [][]int
		NavArray   []NavArray
		NavCurrent int
		Posts      db.Blogs
		Topics     db.Topics
		TopicSel   string
	}{
		NavArray:   getNavArray(num),
		NavCurrent: num.Current,
		Posts:      bs,
		Topics:     ts,
		TopicSel:   topic,
	})
	if err != nil {
		writeErr(w, 500, err)
	}
}

/*
NavData holds values related to page-navigation in `/posts` handler.
Last, is the lastest possible page, and Current represents the currently
requested page for by the client. Can also be used in thread-navigation
*/
type NavData struct {
	Last    int
	Current int
}

/*
NavArray struct holds each element
of nav-tool to be shown in the view
*/
type NavArray struct {
	B bool
	V int
}

/*
getNavArray returns the desired navigation page number array
*/
func getNavArray(idxs NavData) []NavArray {
	arr := []NavArray{}

	var lf int8 // int that keep tracks how many of last iretations are consecutively false
	var cf bool // bool that if more false to be added in the array or not

	for i := 0; i < idxs.Last+1; i++ {
		if i-1 == idxs.Current || i == idxs.Current || i+1 == idxs.Current {
			arr = append(arr, NavArray{B: true, V: i})
			lf = 0
			cf = true
		} else if i+1 == idxs.Last || i == idxs.Last {
			arr = append(arr, NavArray{B: true, V: i})
			lf = 0
			cf = true
		} else if i == 0 {
			arr = append(arr, NavArray{B: true, V: i})
			lf = 0
			cf = true
		} else {
			if cf {
				arr = append(arr, NavArray{B: false, V: i})
			}

			lf++
		}

		if lf == 1 {
			cf = false
		}
	}

	return arr
}
