package oauth

import (
	"os"

	"github.com/coreos/go-oidc"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type GoogleOAuthDriver struct{}

func NewGoogleOAuthDriver() *GoogleOAuthDriver {
	return &GoogleOAuthDriver{}
}

func (oauth *GoogleOAuthDriver) GenerateUrl() string {
	// Google認証を行うためのリダイレクト先のURL
	redirectURL := os.Getenv("BACKEND_URL") + "/signup"

	clientID := os.Getenv("GOOGLE_CLIENT_ID")
	clientSecret := os.Getenv("GOOGLE_CLIENT_SECRET")
	//	認証情報を取得
	config := &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURL:  redirectURL,
		Endpoint:     google.Endpoint,
		Scopes:       []string{oidc.ScopeOpenID, "email"},
	}
	// URLの生成
	return config.AuthCodeURL("state", oauth2.AccessTypeOffline)
}
