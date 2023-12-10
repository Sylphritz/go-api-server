package main

import (
	"github.com/sylphritz/go-api-server/pkg/config"
	"github.com/sylphritz/go-api-server/pkg/server"
)

func main() {
	server.StartServer(config.GetConfig().Port)
}
