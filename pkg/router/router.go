package router

import (
	"github.com/gin-gonic/gin"
	"github.com/sylphritz/go-api-server/pkg/router/authrouter"
	"github.com/sylphritz/go-api-server/pkg/router/blogrouter"
	"github.com/sylphritz/go-api-server/pkg/router/categoryrouter"
	"github.com/sylphritz/go-api-server/pkg/router/emailrouter"
	"github.com/sylphritz/go-api-server/pkg/router/pingrouter"
	"github.com/sylphritz/go-api-server/pkg/router/postrouter"
	"github.com/sylphritz/go-api-server/pkg/router/subscriptionrouter"
	"github.com/sylphritz/go-api-server/pkg/router/userrouter"
)

func SetupRoutes(r *gin.Engine) {
	pingrouter.SetupRoutes(r)
	userrouter.SetupRoutes(r)
	authrouter.SetupRoutes(r)
	subscriptionrouter.SetupRoutes(r)
	emailrouter.SetupRoutes(r)
	blogrouter.SetupRoutes(r)
	categoryrouter.SetupRoutes(r)
	postrouter.SetupRoutes(r)
}
