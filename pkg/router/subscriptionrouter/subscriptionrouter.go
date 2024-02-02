package subscriptionrouter

import (
	"github.com/gin-gonic/gin"
	"github.com/sylphritz/go-api-server/pkg/controller/subscriptionctrl"
	"github.com/sylphritz/go-api-server/pkg/middleware"
	"github.com/sylphritz/go-api-server/pkg/util"
)

var BasePath string = "/subscription"

func SetupRoutes(r *gin.Engine) {
	r.GET(util.GetApiPath(BasePath, "/checkout"), middleware.AuthMiddleware, subscriptionctrl.GetCheckoutURL)
}
