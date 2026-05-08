package main

import (
	"backend-go/config"
	"backend-go/router"
)

func main() {
	db := config.InitDB()
	router.InitialRouter(db)
}
