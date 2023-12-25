package main

import (
	"github.com/sylphritz/go-api-server/pkg/config"
	"github.com/sylphritz/go-api-server/pkg/db"
	"github.com/sylphritz/go-api-server/pkg/server"
	"github.com/sylphritz/go-api-server/pkg/session"
)

func main() {
	session.Init()
	config.LoadEnv()
	config.SetConfig()

	db.Connect()
	server.StartServer(config.GetConfig().Port)
}
