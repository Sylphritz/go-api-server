package ping_router

import (
	"github.com/gin-gonic/gin"
	"github.com/sylphritz/go-api-server/pkg/controller"
	"github.com/sylphritz/go-api-server/pkg/util"
)

var BasePath string = "/ping"

func SetupRoutes(r *gin.Engine) {
	r.GET(util.GetApiPath(BasePath, "/"), controller.GetPing)
}
