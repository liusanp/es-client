package main

import (
	"es-client/commons"
	"es-client/router"
	"log"
	"strconv"
)

func main() {
	commons.LoadConfig()
	port := int(commons.GetConfig("app.port").(int))
	log.Println("app.port:", port)
	r := router.Router()
	r.Run(":" + strconv.FormatInt(int64(port), 10))
}
