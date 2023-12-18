package auth

import (
	"context"
	"encoding/json"
	"io"
	"log"

	"github.com/sylphritz/go-api-server/pkg/config"
	"github.com/sylphritz/go-api-server/pkg/db/schema"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var GoogleOAuthScopes = []string{
	"https://www.googleapis.com/auth/userinfo.email",
}

func GetGoogleOAuthConfig() *oauth2.Config {
	appConfig := config.GetConfig()
	return &oauth2.Config{
		ClientID:     appConfig.OAuthGoogleClientID,
		ClientSecret: appConfig.OAuthGoogleClientSecret,
		Scopes:       GoogleOAuthScopes,
		Endpoint:     google.Endpoint,
	}
}

func GetGoogleOAuthURL(redirectUri string) (string, string) {
	config := GetGoogleOAuthConfig()
	config.RedirectURL = redirectUri

	// Prevent CSRF attacks
	verifier := oauth2.GenerateVerifier()
	url := config.AuthCodeURL(verifier, oauth2.AccessTypeOffline, oauth2.S256ChallengeOption(verifier))

	return url, verifier
}

func GetGoogleUserInfo(token *oauth2.Token) (*schema.User, error) {
	endpoint := "https://www.googleapis.com/oauth2/v3/userinfo"
	authClient := oauth2.NewClient(context.Background(), oauth2.StaticTokenSource(token))

	// Make a GET request to retrieve user information
	res, err := authClient.Get(endpoint)
	if err != nil {
		log.Println("Error trying to get user info.")
		return nil, err
	}
	defer res.Body.Close()

	// Read the response
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println("Error trying to read the response body.")
		return nil, err
	}

	user := schema.User{}
	if err := json.Unmarshal(body, &user); err != nil {
		log.Println("Error trying to parse user info.")
		return nil, err
	}

	user.GoogleOAuthToken = token.AccessToken
	user.GoogleOAuthRefreshToken = token.RefreshToken

	return &user, nil
}
