package main

import (
	"agfun/tv/dbcentral/mysqldb"
	"agfun/tv/router"
	"agfun/tv/service"
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
