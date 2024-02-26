package session

import (
	"encoding/gob"

	"github.com/gorilla/sessions"
	"github.com/sylphritz/go-api-server/pkg/config"
	"github.com/sylphritz/go-api-server/pkg/db/schema"
)

type UserSessionInfo struct {
	ID            uint   `json:"id"`
	Email         string `json:"email"`
	Authenticated bool   `json:"authenticated"`
}

const UserKey = "user"

func Init() {
	gob.Register(&UserSessionInfo{})
}

func UserInfo(s *sessions.Session) (*UserSessionInfo, bool) {
	u, ok := s.Values[UserKey].(*UserSessionInfo)

	return u, ok
}

func SetUserSession(s *sessions.Session, u *schema.User, infinite bool) {
	if infinite {
		s.Options.MaxAge = 31536000 // One year
	} else {
		s.Options.MaxAge = config.GetConfig().SessionDuration
	}

	s.Values[UserKey] = UserSessionInfo{
		ID:            u.ID,
		Email:         u.Email,
		Authenticated: true,
	}
}

func IsUserSessionValid(s *sessions.Session) bool {
	u, ok := UserInfo(s)

	return ok && u.Authenticated
}
