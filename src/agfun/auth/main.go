package main

import (
	"agfun/auth/router"
	"agfun/auth/service"
	"log"
)

func main() {

	service.Init()

	r := router.Init()

	log.Fatal(r.Run(":8080"))
}
