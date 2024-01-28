package subscription_router

import (
	"github.com/gin-gonic/gin"
	"github.com/sylphritz/go-api-server/pkg/controller"
	"github.com/sylphritz/go-api-server/pkg/middleware"
	"github.com/sylphritz/go-api-server/pkg/util"
)

var BasePath string = "/subscription"

func SetupRoutes(r *gin.Engine) {
	r.GET(util.GetApiPath(BasePath, "/checkout"), middleware.AuthMiddleware, controller.GetCheckoutURL)
}
