package main

import (
	"agfun/movie/service"
	"github.com/kataras/iris"
)

func main() {

	defaultSvc := service.NewMovieSvc()
	defer defaultSvc.SysDB.Close()
	defer defaultSvc.AuthDB.Close()
	defer defaultSvc.Dynamic.Close()
	defaultSvc.SysDB.CreateTable()

	router := service.NewRouter()
	router.Svr = defaultSvc
	router.Handle()

	defaultSvc.Run(iris.Addr(":8080"))
}
