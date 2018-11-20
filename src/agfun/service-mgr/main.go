package main

import (
	"agfun/service-mgr/dbcentral/mysqldb"
	"agfun/service-mgr/router"
	"agfun/service-mgr/service"
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
