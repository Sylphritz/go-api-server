package session

import (
	"encoding/gob"
	"fmt"

	"github.com/gorilla/sessions"
	"github.com/sylphritz/go-api-server/pkg/config"
	"github.com/sylphritz/go-api-server/pkg/db/schema"
)

type UserSessionInfo struct {
	Email         string `json:"email"`
	Authenticated bool   `json:"authenticated"`
}

const userKey = "user"

func Init() {
	gob.Register(&UserSessionInfo{})
}

func SetUserSession(s *sessions.Session, u *schema.User, infinite bool) {
	if infinite {
		s.Options.MaxAge = 31536000 // One year
	} else {
		s.Options.MaxAge = config.GetConfig().SessionDuration
	}

	s.Values[userKey] = UserSessionInfo{
		Email:         u.Email,
		Authenticated: true,
	}

	fmt.Println(s.Values[userKey])
	fmt.Println(u)
}

func IsUserSessionValid(s *sessions.Session) bool {
	u, ok := s.Values[userKey].(*UserSessionInfo)

	return ok && u.Authenticated
}
