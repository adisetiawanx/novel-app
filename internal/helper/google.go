package helper

import (
	"github.com/adisetiawanx/novel-app/internal/app"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"strings"
)

func GetGoogleOAuthConfig() *oauth2.Config {
	return &oauth2.Config{
		ClientID:     app.Config.GoogleOauth.ClientID,
		ClientSecret: app.Config.GoogleOauth.ClientSecret,
		RedirectURL:  app.Config.GoogleOauth.RedirectURL,
		Scopes: []string{
			"email",
			"profile",
		},
		Endpoint: google.Endpoint,
	}
}

func IsMobileDevice(userAgent string) bool {
	return strings.Contains(userAgent, "Mobile") || strings.Contains(userAgent, "Android") || strings.Contains(userAgent, "iPhone")
}
