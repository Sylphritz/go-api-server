package ping

import (
	"github.com/gin-gonic/gin"
	"github.com/sylphritz/go-api-server/pkg/controllers"
)

func SetupRoutes(r *gin.Engine) {
	r.GET("/ping", controllers.GetPing)
}
