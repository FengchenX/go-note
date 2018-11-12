package main

import (
	"agfun/agfun-service/router"
	"agfun/test/controller"
	"log"
)

func main() {
	r := router.Init()
	group:=r.Group("/test" )
	group.GET("/", controller.Test)

	log.Fatal(r.Run(":8080"))
}