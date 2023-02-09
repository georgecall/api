package main

import (
	"github.com/georgecall/api/database"
	"github.com/georgecall/api/server"
)

func main() {
	database.StartDB()
	bookServer := server.NewServer()

	bookServer.Run()
}
