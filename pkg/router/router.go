package router

import (
	"github.com/gin-gonic/gin"
	"github.com/sylphritz/go-api-server/pkg/router/ping"
	"github.com/sylphritz/go-api-server/pkg/router/user"
)

func SetupRoutes(r *gin.Engine) {
	ping.SetupRoutes(r)
	user.SetupRoutes(r)
}
