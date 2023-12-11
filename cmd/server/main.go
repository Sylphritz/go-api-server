package main

import (
	"github.com/sylphritz/go-api-server/pkg/config"
	"github.com/sylphritz/go-api-server/pkg/db"
	"github.com/sylphritz/go-api-server/pkg/server"
)

func main() {
	db.Connect()
	server.StartServer(config.GetConfig().Port)
}
