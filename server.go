package main

import (
	"go-blueprint/database"
	"go-blueprint/routes"
)

func main() {
	database.Init()

	e := routes.Init()

	e.Logger.Fatal(e.Start(":1234"))
}
