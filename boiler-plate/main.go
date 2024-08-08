package main

import (
	golog "log"

	database "example/boiler-plate/database"

	"example/boiler-plate/server"
)

func main() {
	db := database.ConnectDataBase()
	golog.Println("Connected to database")

	server, err := server.Init(db)
	if err != nil {
		golog.Fatal("Error initializing server", err)
	}
	golog.Println("Initialized server")

	err = server.Start()
	if err != nil {
		golog.Fatal("Error starting server", err)
	}
	golog.Println("Started server")
}
