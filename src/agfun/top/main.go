package main

import (
	"agfun/top/dbcentral/mysqldb"
	"agfun/top/router"
	"agfun/top/service"
)

func main() {

	defaultSvc := service.GetDefaultSvc()
	defer defaultSvc.Dynamic.Close()
	defer defaultSvc.AuthDB.Close()
	defer defaultSvc.SysDB.Close()

	mysqldb.CreateTable()
	router := router.Init()
	router.Run(":9050")
}
