package blogrouter

import (
	"github.com/gin-gonic/gin"
	"github.com/sylphritz/go-api-server/pkg/controller/blogctrl"
	"github.com/sylphritz/go-api-server/pkg/util"
)

var BasePath string = "/blog"

func SetupRoutes(r *gin.Engine) {
	blogctrlusr := blogctrl.BlogControllerByUser()

	r.GET(util.GetApiPath(BasePath, "/"), blogctrlusr.GetAll)
}