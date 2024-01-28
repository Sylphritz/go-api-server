package router

import (
	"github.com/gin-gonic/gin"
	"github.com/sylphritz/go-api-server/pkg/router/auth_router"
	"github.com/sylphritz/go-api-server/pkg/router/ping_router"
	"github.com/sylphritz/go-api-server/pkg/router/subscription_router"
	"github.com/sylphritz/go-api-server/pkg/router/user_router"
)

func SetupRoutes(r *gin.Engine) {
	ping_router.SetupRoutes(r)
	user_router.SetupRoutes(r)
	auth_router.SetupRoutes(r)
	subscription_router.SetupRoutes(r)
}
