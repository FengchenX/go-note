package main

import (
	"agfun/tv-mgr/dbcentral/mysqldb"
	"agfun/tv-mgr/router"
	"agfun/tv-mgr/service"
)

func main() {

	defaultSvc := service.GetDefaultSvc()
	defer defaultSvc.Dynamic.Close()
	defer defaultSvc.AuthDB.Close()
	defer defaultSvc.SysDB.Close()

	mysqldb.CreateTable()
	router := router.Init()
	router.Run(":8080")
}
