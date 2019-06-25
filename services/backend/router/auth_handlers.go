package router

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gorilla/context"
	"github.com/houseofbosons/houseofbosons/services/backend/db"
	"github.com/houseofbosons/houseofbosons/services/backend/middleware"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var (
	oauthStateStr string         // pseudo-random string
	oauthConfig   *oauth2.Config // contains configuration details for google oauth login
)

func init() {
	oauthConfig = &oauth2.Config{
		RedirectURL:  os.Getenv("GOOGLE_AUTH_REDIRECT_URL"),
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
	}

	oauthStateStr = "pseudo-random"
}

// GoogleLogin redirects login requests to google's oauth api
func GoogleLogin(w http.ResponseWriter, r *http.Request) {
	url := oauthConfig.AuthCodeURL(oauthStateStr)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect) // Redirecting to Google
}

// GoogleCallback handles googles user info response while oauth login
func GoogleCallback(w http.ResponseWriter, r *http.Request) {
	// reading the user's data from google apis
	data, err := userInfo(r.FormValue("state"), r.FormValue("code"))
	if err != nil {
		writeErr(w, 500, err)
		return
	}

	var admin db.Admin
	err = json.Unmarshal(data, &admin)
	if err != nil {
		writeErr(w, 500, err)
		return
	}

	// checking if that email is allowed to login or not, basically if admin's email or not
	if admin.Email != os.Getenv("AUTHORIZED_EMAIL") {
		writeErr(w, 500, fmt.Errorf("email %v is not authorized for login", admin.Email))
		return
	}

	// querying the admin data from database
	err = admin.Query()

	switch err {
	case sql.ErrNoRows:
		// no record (row) found with the admin's email
		fmt.Printf("Admin %v not found\n", admin.Email)
		_, err := admin.Insert()
		if err != nil {
			fmt.Printf("couldn't insert admin, %v\n", err)
			writeErr(w, 500, err)
		}
	case nil:
		// everything's file, user exists and allowed to login
	default:
		// some internal error
		writeErr(w, 500, err)
		return
	}

	// setting up user's cookie for authentication
	session, err := middleware.CookieStore.Get(r, "session")
	session.Values["admin_email"] = admin.Email
	session.Save(r, w)

	// writing the admin json to the client
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(admin)
}

// userInfo sends a request to google apis and gets the user's data for us
func userInfo(state, code string) ([]byte, error) {
	if state != oauthStateStr {
		return nil, fmt.Errorf("invalid oauth state")
	}

	token, err := oauthConfig.Exchange(oauth2.NoContext, code)
	if err != nil {
		return nil, err
	}

	resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

// CurrentUser writes currenntly logged in user's info to teh client
func CurrentUser(w http.ResponseWriter, r *http.Request) {
	// reading session info passed via middleware.CheckAuth middleware
	// that this function (handler) is going to be passed on in the router
	email := context.Get(r, "admin_e")
	fmt.Printf("%v\n", email)

	if email == nil {
		writeErr(w, http.StatusBadRequest, fmt.Errorf("unable to read auth cookie"))
		return
	}

	admin := db.Admin{Email: email.(string)}

	err := admin.Query()
	if err != nil {
		writeErr(w, http.StatusBadRequest, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(admin)
}
