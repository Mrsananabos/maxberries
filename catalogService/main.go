package main

import (
	"catalogService/routes"
	"catalogService/storage/migration"
)

func main() {
	err := migration.DoMigration()
	if err != nil {
		panic(err)
	}

	routes.LoadRoutes()
}
