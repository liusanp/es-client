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
	var err error
	config, err = commons.InitESClient()
	if err != nil {
		log.Fatal(err)
	}
	r := router.Router()

	address := fmt.Sprintf(":%d", config.App.Port)
	log.Fatal(r.Run(address))
}

func main() {
	run()
}
