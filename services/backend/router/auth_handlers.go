package router

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var (
	oauthStateStr string
	oauthConfig   *oauth2.Config
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

// GoogleLoginHandeler .
func GoogleLoginHandeler(w http.ResponseWriter, r *http.Request) {
	url := oauthConfig.AuthCodeURL(oauthStateStr)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect) // Redirecting to Google
}

// GoogleCallbackHandler .
func GoogleCallbackHandler(w http.ResponseWriter, r *http.Request) {
	_, err := userInfo(r.FormValue("state"), r.FormValue("code"))
	if err != nil {
		writeErr(w, 500, err)
		return
	}

	fmt.Fprintf(w, "login has been successful")
}

// userInfo .
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
