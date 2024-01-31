package email_router

import (
	"github.com/gin-gonic/gin"
	"github.com/sylphritz/go-api-server/pkg/controller"
	"github.com/sylphritz/go-api-server/pkg/util"
)

var BasePath string = "/email"

func SetupRoutes(r *gin.Engine) {
	r.GET(util.GetApiPath(BasePath, "/test"), controller.SendTestEmail)
}
