package ping

import (
	"path"

	"github.com/gin-gonic/gin"
	"github.com/sylphritz/go-api-server/pkg/controller"
)

var BasePath string = "/ping"

func SetupRoutes(r *gin.Engine) {
	r.GET(path.Join(BasePath, "/"), controller.GetPing)
}
