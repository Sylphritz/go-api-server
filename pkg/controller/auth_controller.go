package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sylphritz/go-api-server/pkg/auth"
	"github.com/sylphritz/go-api-server/pkg/db/service"
	"github.com/sylphritz/go-api-server/pkg/response"
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
		response.RespondWithError(c, response.NewErrorResponse(
			http.StatusInternalServerError,
			err.Error(),
		))
		return
	}

	user, err := auth.GetGoogleUserInfo(token)
	if err != nil {
		log.Println(err)
		response.RespondWithError(c, response.NewErrorResponse(
			http.StatusInternalServerError,
			err.Error(),
		))
		return
	}

	// TODO: Should be find or create instead
	service.CreateUser(user)
}
