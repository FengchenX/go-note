package main

import (
	"agfun/agfun-service/dbcentral/mysqldb"
	"agfun/agfun-service/router"
	"agfun/test/controller"
	"log"
)

func main() {
	sysdb := mysqldb.InitSysDB()
	defer sysdb.Close()

	r := router.Init()
	group := r.Group("/test")
	group.GET("/", controller.Test)

	log.Fatal(r.Run(":8080"))
}
