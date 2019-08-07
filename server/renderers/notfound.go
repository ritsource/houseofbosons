package renderers

import (
	"fmt"
	"net/http"
)

// NotFoundHandler .
func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	renderErr(w, 404, fmt.Sprintf("Page Not Found"))
}
