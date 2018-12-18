package main

import (
	"agfun/movie-mgr/dbcentral/mysqldb"
	"agfun/movie-mgr/router"
	"agfun/movie-mgr/service"
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
