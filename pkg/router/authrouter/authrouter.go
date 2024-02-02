package authrouter

import (
	"github.com/gin-gonic/gin"
	"github.com/sylphritz/go-api-server/pkg/controller/authctrl"
	"github.com/sylphritz/go-api-server/pkg/util"
)

var BasePath string = "/auth"

func SetupRoutes(r *gin.Engine) {
	r.POST(util.GetApiPath(BasePath, "/google"), authctrl.GetGoogleOAuthURL)
	r.POST(util.GetApiPath(BasePath, "/google/exchange"), authctrl.ExchangeToken)
	r.POST(util.GetApiPath(BasePath, "/session"), authctrl.CheckValidSession)
}
