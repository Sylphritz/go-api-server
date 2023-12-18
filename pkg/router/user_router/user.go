package user_router

import (
	"github.com/gin-gonic/gin"
	"github.com/sylphritz/go-api-server/pkg/controller"
	"github.com/sylphritz/go-api-server/pkg/util"
)

var BasePath string = "/users"

func SetupRoutes(r *gin.Engine) {
	r.POST(util.GetApiPath(BasePath, "/"), controller.RegisterUser)
}
