package main

import (
	"github.com/hafifamudi/task-5-vix-btpns-hafifnurmuhammad/database"
	"github.com/hafifamudi/task-5-vix-btpns-hafifnurmuhammad/router"
)

func main() {
	database.InitDB()
	database.MigrateDB()
	r := router.RouteInit()
	r.Run(":8080")
}
