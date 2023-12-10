package router

import (
	"github.com/gin-gonic/gin"
	"github.com/sylphritz/go-api-server/pkg/router/ping"
)

func SetupRoutes(r *gin.Engine) {
	ping.SetupRoutes(r)
}
