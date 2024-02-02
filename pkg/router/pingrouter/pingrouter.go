package pingrouter

import (
	"github.com/gin-gonic/gin"
	"github.com/sylphritz/go-api-server/pkg/controller/pingctrl"
	"github.com/sylphritz/go-api-server/pkg/util"
)

var BasePath string = "/ping"

func SetupRoutes(r *gin.Engine) {
	r.GET(util.GetApiPath(BasePath, "/"), pingctrl.GetPing)
}
