package router

import (
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
}

func GoogleLoginHandeler() {

}
