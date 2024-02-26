package middleware

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/sylphritz/go-api-server/pkg/response"
	"github.com/sylphritz/go-api-server/pkg/session"
)

// AuthMiddleware checks whether a user is logged in
// and abort the request when not.
func AuthMiddleware(c *gin.Context) {
	userSession, err := session.GetSession(c.Request)
	if err != nil {
		response.RespondInternalErrorWithMessage(c, err)
		return
	}

	valid := session.IsUserSessionValid(userSession)
	if !valid {
		response.RespondWithError(c, response.NewUnauthorizedError("Unauthorized"))
		return
	}

	u, ok := session.UserInfo(userSession)
	if !ok {
		response.RespondInternalErrorWithMessage(c, errors.New("error: could not get user info from cookie"))
		return
	}

	c.Set(session.UserKey, u)

	c.Next()
}
