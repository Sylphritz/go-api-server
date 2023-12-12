package user

import (
	"path"

	"github.com/gin-gonic/gin"
	"github.com/sylphritz/go-api-server/pkg/controller"
)

var BasePath string = "/users"

func SetupRoutes(r *gin.Engine) {
	r.POST(path.Join(BasePath, "/"), controller.RegisterUser)
}
