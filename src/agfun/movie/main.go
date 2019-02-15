package main

import (
	"agfun/movie/dbcentral/mysqldb"
	"agfun/movie/router"
	"agfun/movie/service"
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
