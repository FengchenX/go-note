package main

import (
	"agfun/agfun-service/dbcentral/mysqldb"
	"agfun/agfun-service/router"
	"agfun/auth/controller"
	"log"
)

func main() {
	authdb := mysqldb.InitAuthDB()
	defer authdb.Close()

	r := router.Init()

	group := r.Group("/auth")
	group.GET("/", controller.Test)

	log.Fatal(r.Run(":8080"))
}
