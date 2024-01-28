package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sylphritz/go-api-server/pkg/auth"
	"github.com/sylphritz/go-api-server/pkg/db/schema"
	"github.com/sylphritz/go-api-server/pkg/response"
	"github.com/sylphritz/go-api-server/pkg/service"
	"github.com/sylphritz/go-api-server/pkg/session"
	"golang.org/x/oauth2"
)

type GoogleOAuthURLResponse struct {
	State string `json:"state"`
	URL   string `json:"url"`
}

type RequestWithCallbackURL struct {
	AuthCode    string `json:"auth_code"`
	Verifier    string `json:"verifier"`
	CallbackURL string `json:"callback_url"`
	Infinite    bool   `json:"infinite"`
}

func GetGoogleOAuthURL(c *gin.Context) {
	var request RequestWithCallbackURL

	ok := getRequiredBodyFields[RequestWithCallbackURL](c, &request)
	if !ok {
		return
	}

	url, state := auth.GetGoogleOAuthURL(request.CallbackURL)
	c.JSON(http.StatusOK, GoogleOAuthURLResponse{state, url})
}

func ExchangeToken(c *gin.Context) {
	var request RequestWithCallbackURL

	ok := getRequiredBodyFields[RequestWithCallbackURL](c, &request)
	if !ok {
		return
	}

	config := auth.GetGoogleOAuthConfig()
	config.RedirectURL = request.CallbackURL
	token, err := config.Exchange(c.Request.Context(), request.AuthCode, oauth2.VerifierOption(request.Verifier))
	if err != nil {
		response.RespondInternalErrorWithMessage(c, err)
		return
	}

	user, err := auth.GetGoogleUserInfo(token)
	if err != nil {
		response.RespondInternalErrorWithMessage(c, err)
		return
	}

	err = service.UpdateOrCreateUser(user)
	if err != nil {
		response.RespondInternalErrorWithMessage(c, err)
		return
	}

	userSession, err := session.GetSession(c.Request)
	if err != nil {
		response.RespondInternalErrorWithMessage(c, err)
		return
	}

	session.SetUserSession(userSession, user, request.Infinite)
	err = userSession.Save(c.Request, c.Writer)
	if err != nil {
		response.RespondInternalErrorWithMessage(c, err)
		return
	}

	c.JSON(http.StatusOK, struct{}{})
}

func CheckValidSession(c *gin.Context) {
	userSession, err := session.GetSession(c.Request)
	if err != nil {
		response.RespondInternalErrorWithMessage(c, err)
		return
	}

	if session.IsUserSessionValid(userSession) {
		user := schema.User{}
		fmt.Println(userSession.Values[session.UserKey].(*session.UserSessionInfo))
		service.GetUserById(userSession.Values[session.UserKey].(*session.UserSessionInfo).ID, &user)

		c.JSON(http.StatusOK, user)
		return
	}

	c.JSON(http.StatusOK, false)
}
