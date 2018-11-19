package main

import (
	"agfun/auth/router"
	"agfun/auth/service"
	"log"
)

func main() {

	defaultSvc := service.GetDefaultSvc()
	defer defaultSvc.Dynamic.Close()
	defer defaultSvc.SysDB.Close()
	defer defaultSvc.AuthDB.Close()

	r := router.Init()

	log.Fatal(r.Run(":8080"))
}
