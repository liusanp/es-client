package main

import (
	"es-client/commons"
	"es-client/router"
)

func main() {
	commons.InitConfig()
	r := router.Router()
	r.Run(":8081")
}
