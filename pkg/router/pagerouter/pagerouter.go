package pagerouter

import (
	"github.com/gin-gonic/gin"
	"github.com/sylphritz/go-api-server/pkg/controller/pagectrl"
	"github.com/sylphritz/go-api-server/pkg/middleware"
	"github.com/sylphritz/go-api-server/pkg/util"
)

var BasePath string = "/page"

func SetupRoutes(r *gin.Engine) {
	ctrl := pagectrl.PageController()

	r.GET(util.GetApiPath(BasePath, "/"), middleware.AuthMiddleware, ctrl.GetAll)
	r.GET(util.GetApiPath(BasePath, "/:id"), middleware.AuthMiddleware, ctrl.GetById)
	r.POST(util.GetApiPath(BasePath, "/"), middleware.AuthMiddleware, ctrl.Create)
	r.PUT(util.GetApiPath(BasePath, "/"), middleware.AuthMiddleware, ctrl.Update)
	r.DELETE(util.GetApiPath(BasePath, "/:id"), middleware.AuthMiddleware, ctrl.Delete)
}
