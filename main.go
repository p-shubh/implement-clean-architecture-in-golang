package main

import (
	"main.go/config"
	"main.go/connections/database"
	"main.go/server"
)

func main() {
	c := config.GetConfig()
	db := database.PostgresDatabase(c)
	server.EchoServer(c, db).Start()
}
