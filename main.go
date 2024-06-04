package main

import (
	"es-client/commons"
	"es-client/models"
	"es-client/router"
	"fmt"
	"log"
)

var (
	config models.Config
)

func run() {
	config = commons.InitESClient()
	r := router.Router()

	address := fmt.Sprintf(":%d", config.App.Port)
	log.Fatal(r.Run(address))
}

func main() {
	run()
}
