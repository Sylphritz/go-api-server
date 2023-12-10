package server

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/sylphritz/go-api-server/pkg/router"
)

func getPortString(p int) string {
	return fmt.Sprintf(":%d", p)
}

func StartServer(port int) {
	r := gin.Default()

	// Register routes
	router.SetupRoutes(r)

	log.Println("Starting the web server on port:", port)
	r.Run(getPortString(port))
}
