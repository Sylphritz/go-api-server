package session

import (
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/sylphritz/go-api-server/pkg/config"
)

func GetCookieStore() *sessions.CookieStore {
	config := config.GetConfig()
	store := sessions.NewCookieStore([]byte(config.SessionSecret))

	return store
}

func GetSession(r *http.Request) (*sessions.Session, error) {
	config := config.GetConfig()
	store := GetCookieStore()

	s, err := store.Get(r, config.SessionKey)
	if err != nil {
		return nil, err
	}

	return s, nil
}
