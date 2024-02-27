package blogrouter

import (
	"github.com/gin-gonic/gin"
	"github.com/sylphritz/go-api-server/pkg/controller/blogctrl"
	"github.com/sylphritz/go-api-server/pkg/middleware"
	"github.com/sylphritz/go-api-server/pkg/util"
)

var BasePath string = "/blog"

func SetupRoutes(r *gin.Engine) {
	blogctrlusr := blogctrl.BlogControllerByUser()

	r.GET(util.GetApiPath(BasePath, "/"), middleware.AuthMiddleware, blogctrlusr.GetAll)
	r.POST(util.GetApiPath(BasePath, "/"), middleware.AuthMiddleware, blogctrlusr.Create)
	r.GET(util.GetApiPath(BasePath, "/:id"), middleware.AuthMiddleware, blogctrlusr.GetById)
	// r.PUT(util.GetApiPath(BasePath, "/:id"), middleware.AuthMiddleware, blogctrlusr.GetById)
	// r.DELETE(util.GetApiPath(BasePath, "/:id"), middleware.AuthMiddleware, blogctrlusr.GetById)
}
