package main

import (
	"backend-go/config"
	"backend-go/router"
)

func main() {
	config.InitDB()

	r := router.SetupRouter(config.DbConn)

	r.Run(":8080") // default port
}
