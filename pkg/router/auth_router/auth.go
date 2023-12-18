package auth_router

import (
	"github.com/gin-gonic/gin"
	"github.com/sylphritz/go-api-server/pkg/controller"
	"github.com/sylphritz/go-api-server/pkg/util"
)

var BasePath string = "/auth"

func SetupRoutes(r *gin.Engine) {
	r.POST(util.GetApiPath(BasePath, "/google"), controller.GetGoogleOAuthURL)
	r.POST(util.GetApiPath(BasePath, "/google/exchange"), controller.ExchangeToken)
}
