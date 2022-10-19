package main

import (
	"assignment-02/database"
	"assignment-02/routers"
)

func main() {
	database.InitializeDB()

	var PORT = ":8080"

	routers.StartServer().Run(PORT)
}
